// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/dailyroutine"
	"routine/ent/dailyroutinerec"
	"routine/ent/predicate"
	"routine/ent/programrec"
	"routine/ent/routineactrec"
	"routine/ent/weeklyroutinerec"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DailyRoutineRecUpdate is the builder for updating DailyRoutineRec entities.
type DailyRoutineRecUpdate struct {
	config
	hooks    []Hook
	mutation *DailyRoutineRecMutation
}

// Where appends a list predicates to the DailyRoutineRecUpdate builder.
func (drru *DailyRoutineRecUpdate) Where(ps ...predicate.DailyRoutineRec) *DailyRoutineRecUpdate {
	drru.mutation.Where(ps...)
	return drru
}

// SetAuthor sets the "author" field.
func (drru *DailyRoutineRecUpdate) SetAuthor(u uint64) *DailyRoutineRecUpdate {
	drru.mutation.ResetAuthor()
	drru.mutation.SetAuthor(u)
	return drru
}

// AddAuthor adds u to the "author" field.
func (drru *DailyRoutineRecUpdate) AddAuthor(u int64) *DailyRoutineRecUpdate {
	drru.mutation.AddAuthor(u)
	return drru
}

// SetProgramRecID sets the "program_rec_id" field.
func (drru *DailyRoutineRecUpdate) SetProgramRecID(u uint64) *DailyRoutineRecUpdate {
	drru.mutation.SetProgramRecID(u)
	return drru
}

// SetNillableProgramRecID sets the "program_rec_id" field if the given value is not nil.
func (drru *DailyRoutineRecUpdate) SetNillableProgramRecID(u *uint64) *DailyRoutineRecUpdate {
	if u != nil {
		drru.SetProgramRecID(*u)
	}
	return drru
}

// ClearProgramRecID clears the value of the "program_rec_id" field.
func (drru *DailyRoutineRecUpdate) ClearProgramRecID() *DailyRoutineRecUpdate {
	drru.mutation.ClearProgramRecID()
	return drru
}

// SetWeeklyRoutineRecID sets the "weekly_routine_rec_id" field.
func (drru *DailyRoutineRecUpdate) SetWeeklyRoutineRecID(u uint64) *DailyRoutineRecUpdate {
	drru.mutation.SetWeeklyRoutineRecID(u)
	return drru
}

// SetNillableWeeklyRoutineRecID sets the "weekly_routine_rec_id" field if the given value is not nil.
func (drru *DailyRoutineRecUpdate) SetNillableWeeklyRoutineRecID(u *uint64) *DailyRoutineRecUpdate {
	if u != nil {
		drru.SetWeeklyRoutineRecID(*u)
	}
	return drru
}

// ClearWeeklyRoutineRecID clears the value of the "weekly_routine_rec_id" field.
func (drru *DailyRoutineRecUpdate) ClearWeeklyRoutineRecID() *DailyRoutineRecUpdate {
	drru.mutation.ClearWeeklyRoutineRecID()
	return drru
}

// SetDailyRoutineID sets the "daily_routine_id" field.
func (drru *DailyRoutineRecUpdate) SetDailyRoutineID(u uint64) *DailyRoutineRecUpdate {
	drru.mutation.SetDailyRoutineID(u)
	return drru
}

// SetNillableDailyRoutineID sets the "daily_routine_id" field if the given value is not nil.
func (drru *DailyRoutineRecUpdate) SetNillableDailyRoutineID(u *uint64) *DailyRoutineRecUpdate {
	if u != nil {
		drru.SetDailyRoutineID(*u)
	}
	return drru
}

// ClearDailyRoutineID clears the value of the "daily_routine_id" field.
func (drru *DailyRoutineRecUpdate) ClearDailyRoutineID() *DailyRoutineRecUpdate {
	drru.mutation.ClearDailyRoutineID()
	return drru
}

// SetDate sets the "date" field.
func (drru *DailyRoutineRecUpdate) SetDate(t time.Time) *DailyRoutineRecUpdate {
	drru.mutation.SetDate(t)
	return drru
}

// SetStatus sets the "status" field.
func (drru *DailyRoutineRecUpdate) SetStatus(d dailyroutinerec.Status) *DailyRoutineRecUpdate {
	drru.mutation.SetStatus(d)
	return drru
}

// SetComment sets the "comment" field.
func (drru *DailyRoutineRecUpdate) SetComment(s string) *DailyRoutineRecUpdate {
	drru.mutation.SetComment(s)
	return drru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (drru *DailyRoutineRecUpdate) SetNillableComment(s *string) *DailyRoutineRecUpdate {
	if s != nil {
		drru.SetComment(*s)
	}
	return drru
}

// ClearComment clears the value of the "comment" field.
func (drru *DailyRoutineRecUpdate) ClearComment() *DailyRoutineRecUpdate {
	drru.mutation.ClearComment()
	return drru
}

// SetUpdatedAt sets the "updated_at" field.
func (drru *DailyRoutineRecUpdate) SetUpdatedAt(t time.Time) *DailyRoutineRecUpdate {
	drru.mutation.SetUpdatedAt(t)
	return drru
}

// SetDailyRoutine sets the "daily_routine" edge to the DailyRoutine entity.
func (drru *DailyRoutineRecUpdate) SetDailyRoutine(d *DailyRoutine) *DailyRoutineRecUpdate {
	return drru.SetDailyRoutineID(d.ID)
}

// SetProgramRec sets the "program_rec" edge to the ProgramRec entity.
func (drru *DailyRoutineRecUpdate) SetProgramRec(p *ProgramRec) *DailyRoutineRecUpdate {
	return drru.SetProgramRecID(p.ID)
}

// SetWeeklyRoutineRec sets the "weekly_routine_rec" edge to the WeeklyRoutineRec entity.
func (drru *DailyRoutineRecUpdate) SetWeeklyRoutineRec(w *WeeklyRoutineRec) *DailyRoutineRecUpdate {
	return drru.SetWeeklyRoutineRecID(w.ID)
}

// AddRoutineActRecIDs adds the "routine_act_recs" edge to the RoutineActRec entity by IDs.
func (drru *DailyRoutineRecUpdate) AddRoutineActRecIDs(ids ...uint64) *DailyRoutineRecUpdate {
	drru.mutation.AddRoutineActRecIDs(ids...)
	return drru
}

// AddRoutineActRecs adds the "routine_act_recs" edges to the RoutineActRec entity.
func (drru *DailyRoutineRecUpdate) AddRoutineActRecs(r ...*RoutineActRec) *DailyRoutineRecUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return drru.AddRoutineActRecIDs(ids...)
}

// Mutation returns the DailyRoutineRecMutation object of the builder.
func (drru *DailyRoutineRecUpdate) Mutation() *DailyRoutineRecMutation {
	return drru.mutation
}

// ClearDailyRoutine clears the "daily_routine" edge to the DailyRoutine entity.
func (drru *DailyRoutineRecUpdate) ClearDailyRoutine() *DailyRoutineRecUpdate {
	drru.mutation.ClearDailyRoutine()
	return drru
}

// ClearProgramRec clears the "program_rec" edge to the ProgramRec entity.
func (drru *DailyRoutineRecUpdate) ClearProgramRec() *DailyRoutineRecUpdate {
	drru.mutation.ClearProgramRec()
	return drru
}

// ClearWeeklyRoutineRec clears the "weekly_routine_rec" edge to the WeeklyRoutineRec entity.
func (drru *DailyRoutineRecUpdate) ClearWeeklyRoutineRec() *DailyRoutineRecUpdate {
	drru.mutation.ClearWeeklyRoutineRec()
	return drru
}

// ClearRoutineActRecs clears all "routine_act_recs" edges to the RoutineActRec entity.
func (drru *DailyRoutineRecUpdate) ClearRoutineActRecs() *DailyRoutineRecUpdate {
	drru.mutation.ClearRoutineActRecs()
	return drru
}

// RemoveRoutineActRecIDs removes the "routine_act_recs" edge to RoutineActRec entities by IDs.
func (drru *DailyRoutineRecUpdate) RemoveRoutineActRecIDs(ids ...uint64) *DailyRoutineRecUpdate {
	drru.mutation.RemoveRoutineActRecIDs(ids...)
	return drru
}

// RemoveRoutineActRecs removes "routine_act_recs" edges to RoutineActRec entities.
func (drru *DailyRoutineRecUpdate) RemoveRoutineActRecs(r ...*RoutineActRec) *DailyRoutineRecUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return drru.RemoveRoutineActRecIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (drru *DailyRoutineRecUpdate) Save(ctx context.Context) (int, error) {
	drru.defaults()
	return withHooks(ctx, drru.sqlSave, drru.mutation, drru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (drru *DailyRoutineRecUpdate) SaveX(ctx context.Context) int {
	affected, err := drru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (drru *DailyRoutineRecUpdate) Exec(ctx context.Context) error {
	_, err := drru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drru *DailyRoutineRecUpdate) ExecX(ctx context.Context) {
	if err := drru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (drru *DailyRoutineRecUpdate) defaults() {
	if _, ok := drru.mutation.UpdatedAt(); !ok {
		v := dailyroutinerec.UpdateDefaultUpdatedAt()
		drru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drru *DailyRoutineRecUpdate) check() error {
	if v, ok := drru.mutation.Status(); ok {
		if err := dailyroutinerec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "DailyRoutineRec.status": %w`, err)}
		}
	}
	return nil
}

func (drru *DailyRoutineRecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := drru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dailyroutinerec.Table, dailyroutinerec.Columns, sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64))
	if ps := drru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := drru.mutation.Author(); ok {
		_spec.SetField(dailyroutinerec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := drru.mutation.AddedAuthor(); ok {
		_spec.AddField(dailyroutinerec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := drru.mutation.Date(); ok {
		_spec.SetField(dailyroutinerec.FieldDate, field.TypeTime, value)
	}
	if value, ok := drru.mutation.Status(); ok {
		_spec.SetField(dailyroutinerec.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := drru.mutation.Comment(); ok {
		_spec.SetField(dailyroutinerec.FieldComment, field.TypeString, value)
	}
	if drru.mutation.CommentCleared() {
		_spec.ClearField(dailyroutinerec.FieldComment, field.TypeString)
	}
	if value, ok := drru.mutation.UpdatedAt(); ok {
		_spec.SetField(dailyroutinerec.FieldUpdatedAt, field.TypeTime, value)
	}
	if drru.mutation.DailyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.DailyRoutineTable,
			Columns: []string{dailyroutinerec.DailyRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drru.mutation.DailyRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.DailyRoutineTable,
			Columns: []string{dailyroutinerec.DailyRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if drru.mutation.ProgramRecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.ProgramRecTable,
			Columns: []string{dailyroutinerec.ProgramRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drru.mutation.ProgramRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.ProgramRecTable,
			Columns: []string{dailyroutinerec.ProgramRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if drru.mutation.WeeklyRoutineRecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.WeeklyRoutineRecTable,
			Columns: []string{dailyroutinerec.WeeklyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutinerec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drru.mutation.WeeklyRoutineRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.WeeklyRoutineRecTable,
			Columns: []string{dailyroutinerec.WeeklyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutinerec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if drru.mutation.RoutineActRecsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutinerec.RoutineActRecsTable,
			Columns: []string{dailyroutinerec.RoutineActRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drru.mutation.RemovedRoutineActRecsIDs(); len(nodes) > 0 && !drru.mutation.RoutineActRecsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutinerec.RoutineActRecsTable,
			Columns: []string{dailyroutinerec.RoutineActRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drru.mutation.RoutineActRecsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutinerec.RoutineActRecsTable,
			Columns: []string{dailyroutinerec.RoutineActRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, drru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyroutinerec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	drru.mutation.done = true
	return n, nil
}

// DailyRoutineRecUpdateOne is the builder for updating a single DailyRoutineRec entity.
type DailyRoutineRecUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DailyRoutineRecMutation
}

// SetAuthor sets the "author" field.
func (drruo *DailyRoutineRecUpdateOne) SetAuthor(u uint64) *DailyRoutineRecUpdateOne {
	drruo.mutation.ResetAuthor()
	drruo.mutation.SetAuthor(u)
	return drruo
}

// AddAuthor adds u to the "author" field.
func (drruo *DailyRoutineRecUpdateOne) AddAuthor(u int64) *DailyRoutineRecUpdateOne {
	drruo.mutation.AddAuthor(u)
	return drruo
}

// SetProgramRecID sets the "program_rec_id" field.
func (drruo *DailyRoutineRecUpdateOne) SetProgramRecID(u uint64) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetProgramRecID(u)
	return drruo
}

// SetNillableProgramRecID sets the "program_rec_id" field if the given value is not nil.
func (drruo *DailyRoutineRecUpdateOne) SetNillableProgramRecID(u *uint64) *DailyRoutineRecUpdateOne {
	if u != nil {
		drruo.SetProgramRecID(*u)
	}
	return drruo
}

// ClearProgramRecID clears the value of the "program_rec_id" field.
func (drruo *DailyRoutineRecUpdateOne) ClearProgramRecID() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearProgramRecID()
	return drruo
}

// SetWeeklyRoutineRecID sets the "weekly_routine_rec_id" field.
func (drruo *DailyRoutineRecUpdateOne) SetWeeklyRoutineRecID(u uint64) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetWeeklyRoutineRecID(u)
	return drruo
}

// SetNillableWeeklyRoutineRecID sets the "weekly_routine_rec_id" field if the given value is not nil.
func (drruo *DailyRoutineRecUpdateOne) SetNillableWeeklyRoutineRecID(u *uint64) *DailyRoutineRecUpdateOne {
	if u != nil {
		drruo.SetWeeklyRoutineRecID(*u)
	}
	return drruo
}

// ClearWeeklyRoutineRecID clears the value of the "weekly_routine_rec_id" field.
func (drruo *DailyRoutineRecUpdateOne) ClearWeeklyRoutineRecID() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearWeeklyRoutineRecID()
	return drruo
}

// SetDailyRoutineID sets the "daily_routine_id" field.
func (drruo *DailyRoutineRecUpdateOne) SetDailyRoutineID(u uint64) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetDailyRoutineID(u)
	return drruo
}

// SetNillableDailyRoutineID sets the "daily_routine_id" field if the given value is not nil.
func (drruo *DailyRoutineRecUpdateOne) SetNillableDailyRoutineID(u *uint64) *DailyRoutineRecUpdateOne {
	if u != nil {
		drruo.SetDailyRoutineID(*u)
	}
	return drruo
}

// ClearDailyRoutineID clears the value of the "daily_routine_id" field.
func (drruo *DailyRoutineRecUpdateOne) ClearDailyRoutineID() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearDailyRoutineID()
	return drruo
}

// SetDate sets the "date" field.
func (drruo *DailyRoutineRecUpdateOne) SetDate(t time.Time) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetDate(t)
	return drruo
}

// SetStatus sets the "status" field.
func (drruo *DailyRoutineRecUpdateOne) SetStatus(d dailyroutinerec.Status) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetStatus(d)
	return drruo
}

// SetComment sets the "comment" field.
func (drruo *DailyRoutineRecUpdateOne) SetComment(s string) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetComment(s)
	return drruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (drruo *DailyRoutineRecUpdateOne) SetNillableComment(s *string) *DailyRoutineRecUpdateOne {
	if s != nil {
		drruo.SetComment(*s)
	}
	return drruo
}

// ClearComment clears the value of the "comment" field.
func (drruo *DailyRoutineRecUpdateOne) ClearComment() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearComment()
	return drruo
}

// SetUpdatedAt sets the "updated_at" field.
func (drruo *DailyRoutineRecUpdateOne) SetUpdatedAt(t time.Time) *DailyRoutineRecUpdateOne {
	drruo.mutation.SetUpdatedAt(t)
	return drruo
}

// SetDailyRoutine sets the "daily_routine" edge to the DailyRoutine entity.
func (drruo *DailyRoutineRecUpdateOne) SetDailyRoutine(d *DailyRoutine) *DailyRoutineRecUpdateOne {
	return drruo.SetDailyRoutineID(d.ID)
}

// SetProgramRec sets the "program_rec" edge to the ProgramRec entity.
func (drruo *DailyRoutineRecUpdateOne) SetProgramRec(p *ProgramRec) *DailyRoutineRecUpdateOne {
	return drruo.SetProgramRecID(p.ID)
}

// SetWeeklyRoutineRec sets the "weekly_routine_rec" edge to the WeeklyRoutineRec entity.
func (drruo *DailyRoutineRecUpdateOne) SetWeeklyRoutineRec(w *WeeklyRoutineRec) *DailyRoutineRecUpdateOne {
	return drruo.SetWeeklyRoutineRecID(w.ID)
}

// AddRoutineActRecIDs adds the "routine_act_recs" edge to the RoutineActRec entity by IDs.
func (drruo *DailyRoutineRecUpdateOne) AddRoutineActRecIDs(ids ...uint64) *DailyRoutineRecUpdateOne {
	drruo.mutation.AddRoutineActRecIDs(ids...)
	return drruo
}

// AddRoutineActRecs adds the "routine_act_recs" edges to the RoutineActRec entity.
func (drruo *DailyRoutineRecUpdateOne) AddRoutineActRecs(r ...*RoutineActRec) *DailyRoutineRecUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return drruo.AddRoutineActRecIDs(ids...)
}

// Mutation returns the DailyRoutineRecMutation object of the builder.
func (drruo *DailyRoutineRecUpdateOne) Mutation() *DailyRoutineRecMutation {
	return drruo.mutation
}

// ClearDailyRoutine clears the "daily_routine" edge to the DailyRoutine entity.
func (drruo *DailyRoutineRecUpdateOne) ClearDailyRoutine() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearDailyRoutine()
	return drruo
}

// ClearProgramRec clears the "program_rec" edge to the ProgramRec entity.
func (drruo *DailyRoutineRecUpdateOne) ClearProgramRec() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearProgramRec()
	return drruo
}

// ClearWeeklyRoutineRec clears the "weekly_routine_rec" edge to the WeeklyRoutineRec entity.
func (drruo *DailyRoutineRecUpdateOne) ClearWeeklyRoutineRec() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearWeeklyRoutineRec()
	return drruo
}

// ClearRoutineActRecs clears all "routine_act_recs" edges to the RoutineActRec entity.
func (drruo *DailyRoutineRecUpdateOne) ClearRoutineActRecs() *DailyRoutineRecUpdateOne {
	drruo.mutation.ClearRoutineActRecs()
	return drruo
}

// RemoveRoutineActRecIDs removes the "routine_act_recs" edge to RoutineActRec entities by IDs.
func (drruo *DailyRoutineRecUpdateOne) RemoveRoutineActRecIDs(ids ...uint64) *DailyRoutineRecUpdateOne {
	drruo.mutation.RemoveRoutineActRecIDs(ids...)
	return drruo
}

// RemoveRoutineActRecs removes "routine_act_recs" edges to RoutineActRec entities.
func (drruo *DailyRoutineRecUpdateOne) RemoveRoutineActRecs(r ...*RoutineActRec) *DailyRoutineRecUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return drruo.RemoveRoutineActRecIDs(ids...)
}

// Where appends a list predicates to the DailyRoutineRecUpdate builder.
func (drruo *DailyRoutineRecUpdateOne) Where(ps ...predicate.DailyRoutineRec) *DailyRoutineRecUpdateOne {
	drruo.mutation.Where(ps...)
	return drruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (drruo *DailyRoutineRecUpdateOne) Select(field string, fields ...string) *DailyRoutineRecUpdateOne {
	drruo.fields = append([]string{field}, fields...)
	return drruo
}

// Save executes the query and returns the updated DailyRoutineRec entity.
func (drruo *DailyRoutineRecUpdateOne) Save(ctx context.Context) (*DailyRoutineRec, error) {
	drruo.defaults()
	return withHooks(ctx, drruo.sqlSave, drruo.mutation, drruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (drruo *DailyRoutineRecUpdateOne) SaveX(ctx context.Context) *DailyRoutineRec {
	node, err := drruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (drruo *DailyRoutineRecUpdateOne) Exec(ctx context.Context) error {
	_, err := drruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drruo *DailyRoutineRecUpdateOne) ExecX(ctx context.Context) {
	if err := drruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (drruo *DailyRoutineRecUpdateOne) defaults() {
	if _, ok := drruo.mutation.UpdatedAt(); !ok {
		v := dailyroutinerec.UpdateDefaultUpdatedAt()
		drruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drruo *DailyRoutineRecUpdateOne) check() error {
	if v, ok := drruo.mutation.Status(); ok {
		if err := dailyroutinerec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "DailyRoutineRec.status": %w`, err)}
		}
	}
	return nil
}

func (drruo *DailyRoutineRecUpdateOne) sqlSave(ctx context.Context) (_node *DailyRoutineRec, err error) {
	if err := drruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dailyroutinerec.Table, dailyroutinerec.Columns, sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64))
	id, ok := drruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DailyRoutineRec.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := drruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailyroutinerec.FieldID)
		for _, f := range fields {
			if !dailyroutinerec.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dailyroutinerec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := drruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := drruo.mutation.Author(); ok {
		_spec.SetField(dailyroutinerec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := drruo.mutation.AddedAuthor(); ok {
		_spec.AddField(dailyroutinerec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := drruo.mutation.Date(); ok {
		_spec.SetField(dailyroutinerec.FieldDate, field.TypeTime, value)
	}
	if value, ok := drruo.mutation.Status(); ok {
		_spec.SetField(dailyroutinerec.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := drruo.mutation.Comment(); ok {
		_spec.SetField(dailyroutinerec.FieldComment, field.TypeString, value)
	}
	if drruo.mutation.CommentCleared() {
		_spec.ClearField(dailyroutinerec.FieldComment, field.TypeString)
	}
	if value, ok := drruo.mutation.UpdatedAt(); ok {
		_spec.SetField(dailyroutinerec.FieldUpdatedAt, field.TypeTime, value)
	}
	if drruo.mutation.DailyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.DailyRoutineTable,
			Columns: []string{dailyroutinerec.DailyRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drruo.mutation.DailyRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.DailyRoutineTable,
			Columns: []string{dailyroutinerec.DailyRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if drruo.mutation.ProgramRecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.ProgramRecTable,
			Columns: []string{dailyroutinerec.ProgramRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drruo.mutation.ProgramRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.ProgramRecTable,
			Columns: []string{dailyroutinerec.ProgramRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if drruo.mutation.WeeklyRoutineRecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.WeeklyRoutineRecTable,
			Columns: []string{dailyroutinerec.WeeklyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutinerec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drruo.mutation.WeeklyRoutineRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutinerec.WeeklyRoutineRecTable,
			Columns: []string{dailyroutinerec.WeeklyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutinerec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if drruo.mutation.RoutineActRecsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutinerec.RoutineActRecsTable,
			Columns: []string{dailyroutinerec.RoutineActRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drruo.mutation.RemovedRoutineActRecsIDs(); len(nodes) > 0 && !drruo.mutation.RoutineActRecsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutinerec.RoutineActRecsTable,
			Columns: []string{dailyroutinerec.RoutineActRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drruo.mutation.RoutineActRecsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutinerec.RoutineActRecsTable,
			Columns: []string{dailyroutinerec.RoutineActRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DailyRoutineRec{config: drruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, drruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyroutinerec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	drruo.mutation.done = true
	return _node, nil
}
