package db

import (
	"context"
	"fmt"
	"routine/common/dto"
	"routine/common/helper"
	"routine/ent"
	"routine/ent/program"
	"sync"
)

// QueryProgramBySlug queries program by slug and returns corresponding program
func QueryProgramBySlug(dbClient *ent.Client, c context.Context, slug string) (pg *ent.Program, err error) {
	pg, err = dbClient.Program.Query().
		Where(program.SlugEQ(slug)).
		WithTags().
		WithWeeklyRoutines(func(q *ent.WeeklyRoutineQuery) {
			q.WithDailyRoutines(func(q *ent.DailyRoutineQuery) {
				q.WithRoutineActs()
			})
		}).
		First(c)
	if err != nil {
		return nil, err
	}
	return pg, nil
}

// QueryProgramsByName queries programs by name and returns 5 programs skipping given number of programs
func QueryProgramsByName(dbClient *ent.Client, c context.Context, name string, skip int) (programs []*ent.Program, err error) {
	programs, err = dbClient.Program.Query().
		Where(program.TitleContains(name)).
		Offset(skip).
		Limit(5).
		WithTags().
		All(c)
	if ent.IsNotFound(err) {
		programs = []*ent.Program{}
	} else if err != nil {
		return nil, err
	}
	return programs, nil
}

func CreateDailyProgram(dbClient *ent.Client, c context.Context, p interface{}) (newProgram *ent.Program, err error) {
	return nil, nil
}

// CreateWeeklyProgram creates weekly program and returns created program's ID.
func CreateWeeklyProgram(dbClient *ent.Client, c context.Context, p *dto.CreateWeeklyProgramDto) (pid uint64, err error) {

	// to conduct atomic operation, use transaction
	tx, err := dbClient.Tx(c)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %w", err)
	}

	// first query tags and create tags if not exists
	tags, err := QueryAndCreateTags(tx.Client(), c, p.Tags)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to query and create tags: %w", err))
	}

	// not to query db again, get each index first
	weekIdxMap := make(map[int]int)   // e.g. [week 1] : week 1's index
	dayIdxMap := make(map[[2]int]int) // e.g. [week 2, day 3] : week 2's day 3's index
	// do this concurrently
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, wr := range p.WeeklyRoutines {
			weekIdxMap[wr.Week] = i
		}
	}()
	go func() {
		defer wg.Done()
		for i, dr := range p.DailyRoutines {
			dayIdxMap[[2]int{dr.Week, dr.Day}] = i
		}
	}()

	// first create program while getting indices.
	newProgram, err := tx.Program.Create().
		SetTitle(p.Title).
		SetSlug(helper.Slugify(p.Title) + "-" + helper.RandomHex(4)).
		SetType("weekly").
		SetAuthor(p.Author).
		SetNillableDescription(p.Description).
		SetNillableImage(p.Image).
		AddTagIDs(tags...).
		Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create program: %w", err))
	}

	wg.Wait()

	// then create weekly routines
	bulkW := make([]*ent.WeeklyRoutineCreate, len(p.WeeklyRoutines))
	for i, wr := range p.WeeklyRoutines {
		bulkW[i] = tx.WeeklyRoutine.Create().SetProgramID(newProgram.ID).SetWeek(wr.Week)
	}
	newWeeklyRoutines, err := tx.WeeklyRoutine.CreateBulk(bulkW...).Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create weekly routines: %w", err))
	}

	// then create daily routines
	bulkD := make([]*ent.DailyRoutineCreate, len(p.DailyRoutines))
	for i, dr := range p.DailyRoutines {
		weekId := newWeeklyRoutines[weekIdxMap[dr.Week]].ID
		bulkD[i] = tx.DailyRoutine.Create().SetProgramID(newProgram.ID).SetWeeklyRoutineID(weekId).SetDay(dr.Day)
	}
	newDailyRoutines, err := tx.DailyRoutine.CreateBulk(bulkD...).Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create daily routines: %w", err))
	}

	// then create routine acts
	bulkA := make([]*ent.RoutineActCreate, len(p.RoutineActs))
	for i, ra := range p.RoutineActs {
		dayId := newDailyRoutines[dayIdxMap[[2]int{ra.Week, ra.Day}]].ID
		bulkA[i] = tx.RoutineAct.Create().
			SetDailyRoutineID(dayId).SetActID(ra.ActID).
			SetNillableReps(ra.Reps).SetNillableLap(ra.Lap).
			SetOrder(ra.Order).
			SetNillableWRatio(ra.WRatio).SetWarmup(ra.Warmup)
	}
	_, err = tx.RoutineAct.CreateBulk(bulkA...).Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create routine acts: %w", err))
	}

	tx.Commit()

	return newProgram.ID, nil
}
