// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/dailyroutine"
	"routine/ent/predicate"
	"routine/ent/program"
	"routine/ent/routineact"
	"routine/ent/weeklyroutine"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DailyRoutineUpdate is the builder for updating DailyRoutine entities.
type DailyRoutineUpdate struct {
	config
	hooks    []Hook
	mutation *DailyRoutineMutation
}

// Where appends a list predicates to the DailyRoutineUpdate builder.
func (dru *DailyRoutineUpdate) Where(ps ...predicate.DailyRoutine) *DailyRoutineUpdate {
	dru.mutation.Where(ps...)
	return dru
}

// SetProgramID sets the "program_id" field.
func (dru *DailyRoutineUpdate) SetProgramID(u uint64) *DailyRoutineUpdate {
	dru.mutation.ResetProgramID()
	dru.mutation.SetProgramID(u)
	return dru
}

// SetNillableProgramID sets the "program_id" field if the given value is not nil.
func (dru *DailyRoutineUpdate) SetNillableProgramID(u *uint64) *DailyRoutineUpdate {
	if u != nil {
		dru.SetProgramID(*u)
	}
	return dru
}

// AddProgramID adds u to the "program_id" field.
func (dru *DailyRoutineUpdate) AddProgramID(u int64) *DailyRoutineUpdate {
	dru.mutation.AddProgramID(u)
	return dru
}

// ClearProgramID clears the value of the "program_id" field.
func (dru *DailyRoutineUpdate) ClearProgramID() *DailyRoutineUpdate {
	dru.mutation.ClearProgramID()
	return dru
}

// SetWeekID sets the "week_id" field.
func (dru *DailyRoutineUpdate) SetWeekID(u uint64) *DailyRoutineUpdate {
	dru.mutation.ResetWeekID()
	dru.mutation.SetWeekID(u)
	return dru
}

// SetNillableWeekID sets the "week_id" field if the given value is not nil.
func (dru *DailyRoutineUpdate) SetNillableWeekID(u *uint64) *DailyRoutineUpdate {
	if u != nil {
		dru.SetWeekID(*u)
	}
	return dru
}

// AddWeekID adds u to the "week_id" field.
func (dru *DailyRoutineUpdate) AddWeekID(u int64) *DailyRoutineUpdate {
	dru.mutation.AddWeekID(u)
	return dru
}

// ClearWeekID clears the value of the "week_id" field.
func (dru *DailyRoutineUpdate) ClearWeekID() *DailyRoutineUpdate {
	dru.mutation.ClearWeekID()
	return dru
}

// SetDay sets the "day" field.
func (dru *DailyRoutineUpdate) SetDay(i int) *DailyRoutineUpdate {
	dru.mutation.ResetDay()
	dru.mutation.SetDay(i)
	return dru
}

// AddDay adds i to the "day" field.
func (dru *DailyRoutineUpdate) AddDay(i int) *DailyRoutineUpdate {
	dru.mutation.AddDay(i)
	return dru
}

// AddProgramIDs adds the "program" edge to the Program entity by IDs.
func (dru *DailyRoutineUpdate) AddProgramIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.AddProgramIDs(ids...)
	return dru
}

// AddProgram adds the "program" edges to the Program entity.
func (dru *DailyRoutineUpdate) AddProgram(p ...*Program) *DailyRoutineUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dru.AddProgramIDs(ids...)
}

// AddWeeklyRoutineIDs adds the "weekly_routine" edge to the WeeklyRoutine entity by IDs.
func (dru *DailyRoutineUpdate) AddWeeklyRoutineIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.AddWeeklyRoutineIDs(ids...)
	return dru
}

// AddWeeklyRoutine adds the "weekly_routine" edges to the WeeklyRoutine entity.
func (dru *DailyRoutineUpdate) AddWeeklyRoutine(w ...*WeeklyRoutine) *DailyRoutineUpdate {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return dru.AddWeeklyRoutineIDs(ids...)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (dru *DailyRoutineUpdate) AddRoutineActIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.AddRoutineActIDs(ids...)
	return dru
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (dru *DailyRoutineUpdate) AddRoutineActs(r ...*RoutineAct) *DailyRoutineUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dru.AddRoutineActIDs(ids...)
}

// Mutation returns the DailyRoutineMutation object of the builder.
func (dru *DailyRoutineUpdate) Mutation() *DailyRoutineMutation {
	return dru.mutation
}

// ClearProgram clears all "program" edges to the Program entity.
func (dru *DailyRoutineUpdate) ClearProgram() *DailyRoutineUpdate {
	dru.mutation.ClearProgram()
	return dru
}

// RemoveProgramIDs removes the "program" edge to Program entities by IDs.
func (dru *DailyRoutineUpdate) RemoveProgramIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.RemoveProgramIDs(ids...)
	return dru
}

// RemoveProgram removes "program" edges to Program entities.
func (dru *DailyRoutineUpdate) RemoveProgram(p ...*Program) *DailyRoutineUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dru.RemoveProgramIDs(ids...)
}

// ClearWeeklyRoutine clears all "weekly_routine" edges to the WeeklyRoutine entity.
func (dru *DailyRoutineUpdate) ClearWeeklyRoutine() *DailyRoutineUpdate {
	dru.mutation.ClearWeeklyRoutine()
	return dru
}

// RemoveWeeklyRoutineIDs removes the "weekly_routine" edge to WeeklyRoutine entities by IDs.
func (dru *DailyRoutineUpdate) RemoveWeeklyRoutineIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.RemoveWeeklyRoutineIDs(ids...)
	return dru
}

// RemoveWeeklyRoutine removes "weekly_routine" edges to WeeklyRoutine entities.
func (dru *DailyRoutineUpdate) RemoveWeeklyRoutine(w ...*WeeklyRoutine) *DailyRoutineUpdate {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return dru.RemoveWeeklyRoutineIDs(ids...)
}

// ClearRoutineActs clears all "routine_acts" edges to the RoutineAct entity.
func (dru *DailyRoutineUpdate) ClearRoutineActs() *DailyRoutineUpdate {
	dru.mutation.ClearRoutineActs()
	return dru
}

// RemoveRoutineActIDs removes the "routine_acts" edge to RoutineAct entities by IDs.
func (dru *DailyRoutineUpdate) RemoveRoutineActIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.RemoveRoutineActIDs(ids...)
	return dru
}

// RemoveRoutineActs removes "routine_acts" edges to RoutineAct entities.
func (dru *DailyRoutineUpdate) RemoveRoutineActs(r ...*RoutineAct) *DailyRoutineUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dru.RemoveRoutineActIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dru *DailyRoutineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dru.sqlSave, dru.mutation, dru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dru *DailyRoutineUpdate) SaveX(ctx context.Context) int {
	affected, err := dru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dru *DailyRoutineUpdate) Exec(ctx context.Context) error {
	_, err := dru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dru *DailyRoutineUpdate) ExecX(ctx context.Context) {
	if err := dru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dru *DailyRoutineUpdate) check() error {
	if v, ok := dru.mutation.Day(); ok {
		if err := dailyroutine.DayValidator(v); err != nil {
			return &ValidationError{Name: "day", err: fmt.Errorf(`ent: validator failed for field "DailyRoutine.day": %w`, err)}
		}
	}
	return nil
}

func (dru *DailyRoutineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dailyroutine.Table, dailyroutine.Columns, sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64))
	if ps := dru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dru.mutation.ProgramID(); ok {
		_spec.SetField(dailyroutine.FieldProgramID, field.TypeUint64, value)
	}
	if value, ok := dru.mutation.AddedProgramID(); ok {
		_spec.AddField(dailyroutine.FieldProgramID, field.TypeUint64, value)
	}
	if dru.mutation.ProgramIDCleared() {
		_spec.ClearField(dailyroutine.FieldProgramID, field.TypeUint64)
	}
	if value, ok := dru.mutation.WeekID(); ok {
		_spec.SetField(dailyroutine.FieldWeekID, field.TypeUint64, value)
	}
	if value, ok := dru.mutation.AddedWeekID(); ok {
		_spec.AddField(dailyroutine.FieldWeekID, field.TypeUint64, value)
	}
	if dru.mutation.WeekIDCleared() {
		_spec.ClearField(dailyroutine.FieldWeekID, field.TypeUint64)
	}
	if value, ok := dru.mutation.Day(); ok {
		_spec.SetField(dailyroutine.FieldDay, field.TypeInt, value)
	}
	if value, ok := dru.mutation.AddedDay(); ok {
		_spec.AddField(dailyroutine.FieldDay, field.TypeInt, value)
	}
	if dru.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.ProgramTable,
			Columns: dailyroutine.ProgramPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.RemovedProgramIDs(); len(nodes) > 0 && !dru.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.ProgramTable,
			Columns: dailyroutine.ProgramPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.ProgramTable,
			Columns: dailyroutine.ProgramPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dru.mutation.WeeklyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.WeeklyRoutineTable,
			Columns: dailyroutine.WeeklyRoutinePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.RemovedWeeklyRoutineIDs(); len(nodes) > 0 && !dru.mutation.WeeklyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.WeeklyRoutineTable,
			Columns: dailyroutine.WeeklyRoutinePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.WeeklyRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.WeeklyRoutineTable,
			Columns: dailyroutine.WeeklyRoutinePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dru.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.RemovedRoutineActsIDs(); len(nodes) > 0 && !dru.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyroutine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dru.mutation.done = true
	return n, nil
}

// DailyRoutineUpdateOne is the builder for updating a single DailyRoutine entity.
type DailyRoutineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DailyRoutineMutation
}

// SetProgramID sets the "program_id" field.
func (druo *DailyRoutineUpdateOne) SetProgramID(u uint64) *DailyRoutineUpdateOne {
	druo.mutation.ResetProgramID()
	druo.mutation.SetProgramID(u)
	return druo
}

// SetNillableProgramID sets the "program_id" field if the given value is not nil.
func (druo *DailyRoutineUpdateOne) SetNillableProgramID(u *uint64) *DailyRoutineUpdateOne {
	if u != nil {
		druo.SetProgramID(*u)
	}
	return druo
}

// AddProgramID adds u to the "program_id" field.
func (druo *DailyRoutineUpdateOne) AddProgramID(u int64) *DailyRoutineUpdateOne {
	druo.mutation.AddProgramID(u)
	return druo
}

// ClearProgramID clears the value of the "program_id" field.
func (druo *DailyRoutineUpdateOne) ClearProgramID() *DailyRoutineUpdateOne {
	druo.mutation.ClearProgramID()
	return druo
}

// SetWeekID sets the "week_id" field.
func (druo *DailyRoutineUpdateOne) SetWeekID(u uint64) *DailyRoutineUpdateOne {
	druo.mutation.ResetWeekID()
	druo.mutation.SetWeekID(u)
	return druo
}

// SetNillableWeekID sets the "week_id" field if the given value is not nil.
func (druo *DailyRoutineUpdateOne) SetNillableWeekID(u *uint64) *DailyRoutineUpdateOne {
	if u != nil {
		druo.SetWeekID(*u)
	}
	return druo
}

// AddWeekID adds u to the "week_id" field.
func (druo *DailyRoutineUpdateOne) AddWeekID(u int64) *DailyRoutineUpdateOne {
	druo.mutation.AddWeekID(u)
	return druo
}

// ClearWeekID clears the value of the "week_id" field.
func (druo *DailyRoutineUpdateOne) ClearWeekID() *DailyRoutineUpdateOne {
	druo.mutation.ClearWeekID()
	return druo
}

// SetDay sets the "day" field.
func (druo *DailyRoutineUpdateOne) SetDay(i int) *DailyRoutineUpdateOne {
	druo.mutation.ResetDay()
	druo.mutation.SetDay(i)
	return druo
}

// AddDay adds i to the "day" field.
func (druo *DailyRoutineUpdateOne) AddDay(i int) *DailyRoutineUpdateOne {
	druo.mutation.AddDay(i)
	return druo
}

// AddProgramIDs adds the "program" edge to the Program entity by IDs.
func (druo *DailyRoutineUpdateOne) AddProgramIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.AddProgramIDs(ids...)
	return druo
}

// AddProgram adds the "program" edges to the Program entity.
func (druo *DailyRoutineUpdateOne) AddProgram(p ...*Program) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return druo.AddProgramIDs(ids...)
}

// AddWeeklyRoutineIDs adds the "weekly_routine" edge to the WeeklyRoutine entity by IDs.
func (druo *DailyRoutineUpdateOne) AddWeeklyRoutineIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.AddWeeklyRoutineIDs(ids...)
	return druo
}

// AddWeeklyRoutine adds the "weekly_routine" edges to the WeeklyRoutine entity.
func (druo *DailyRoutineUpdateOne) AddWeeklyRoutine(w ...*WeeklyRoutine) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return druo.AddWeeklyRoutineIDs(ids...)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (druo *DailyRoutineUpdateOne) AddRoutineActIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.AddRoutineActIDs(ids...)
	return druo
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (druo *DailyRoutineUpdateOne) AddRoutineActs(r ...*RoutineAct) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return druo.AddRoutineActIDs(ids...)
}

// Mutation returns the DailyRoutineMutation object of the builder.
func (druo *DailyRoutineUpdateOne) Mutation() *DailyRoutineMutation {
	return druo.mutation
}

// ClearProgram clears all "program" edges to the Program entity.
func (druo *DailyRoutineUpdateOne) ClearProgram() *DailyRoutineUpdateOne {
	druo.mutation.ClearProgram()
	return druo
}

// RemoveProgramIDs removes the "program" edge to Program entities by IDs.
func (druo *DailyRoutineUpdateOne) RemoveProgramIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.RemoveProgramIDs(ids...)
	return druo
}

// RemoveProgram removes "program" edges to Program entities.
func (druo *DailyRoutineUpdateOne) RemoveProgram(p ...*Program) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return druo.RemoveProgramIDs(ids...)
}

// ClearWeeklyRoutine clears all "weekly_routine" edges to the WeeklyRoutine entity.
func (druo *DailyRoutineUpdateOne) ClearWeeklyRoutine() *DailyRoutineUpdateOne {
	druo.mutation.ClearWeeklyRoutine()
	return druo
}

// RemoveWeeklyRoutineIDs removes the "weekly_routine" edge to WeeklyRoutine entities by IDs.
func (druo *DailyRoutineUpdateOne) RemoveWeeklyRoutineIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.RemoveWeeklyRoutineIDs(ids...)
	return druo
}

// RemoveWeeklyRoutine removes "weekly_routine" edges to WeeklyRoutine entities.
func (druo *DailyRoutineUpdateOne) RemoveWeeklyRoutine(w ...*WeeklyRoutine) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return druo.RemoveWeeklyRoutineIDs(ids...)
}

// ClearRoutineActs clears all "routine_acts" edges to the RoutineAct entity.
func (druo *DailyRoutineUpdateOne) ClearRoutineActs() *DailyRoutineUpdateOne {
	druo.mutation.ClearRoutineActs()
	return druo
}

// RemoveRoutineActIDs removes the "routine_acts" edge to RoutineAct entities by IDs.
func (druo *DailyRoutineUpdateOne) RemoveRoutineActIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.RemoveRoutineActIDs(ids...)
	return druo
}

// RemoveRoutineActs removes "routine_acts" edges to RoutineAct entities.
func (druo *DailyRoutineUpdateOne) RemoveRoutineActs(r ...*RoutineAct) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return druo.RemoveRoutineActIDs(ids...)
}

// Where appends a list predicates to the DailyRoutineUpdate builder.
func (druo *DailyRoutineUpdateOne) Where(ps ...predicate.DailyRoutine) *DailyRoutineUpdateOne {
	druo.mutation.Where(ps...)
	return druo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (druo *DailyRoutineUpdateOne) Select(field string, fields ...string) *DailyRoutineUpdateOne {
	druo.fields = append([]string{field}, fields...)
	return druo
}

// Save executes the query and returns the updated DailyRoutine entity.
func (druo *DailyRoutineUpdateOne) Save(ctx context.Context) (*DailyRoutine, error) {
	return withHooks(ctx, druo.sqlSave, druo.mutation, druo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (druo *DailyRoutineUpdateOne) SaveX(ctx context.Context) *DailyRoutine {
	node, err := druo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (druo *DailyRoutineUpdateOne) Exec(ctx context.Context) error {
	_, err := druo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (druo *DailyRoutineUpdateOne) ExecX(ctx context.Context) {
	if err := druo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (druo *DailyRoutineUpdateOne) check() error {
	if v, ok := druo.mutation.Day(); ok {
		if err := dailyroutine.DayValidator(v); err != nil {
			return &ValidationError{Name: "day", err: fmt.Errorf(`ent: validator failed for field "DailyRoutine.day": %w`, err)}
		}
	}
	return nil
}

func (druo *DailyRoutineUpdateOne) sqlSave(ctx context.Context) (_node *DailyRoutine, err error) {
	if err := druo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dailyroutine.Table, dailyroutine.Columns, sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64))
	id, ok := druo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DailyRoutine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := druo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailyroutine.FieldID)
		for _, f := range fields {
			if !dailyroutine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dailyroutine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := druo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := druo.mutation.ProgramID(); ok {
		_spec.SetField(dailyroutine.FieldProgramID, field.TypeUint64, value)
	}
	if value, ok := druo.mutation.AddedProgramID(); ok {
		_spec.AddField(dailyroutine.FieldProgramID, field.TypeUint64, value)
	}
	if druo.mutation.ProgramIDCleared() {
		_spec.ClearField(dailyroutine.FieldProgramID, field.TypeUint64)
	}
	if value, ok := druo.mutation.WeekID(); ok {
		_spec.SetField(dailyroutine.FieldWeekID, field.TypeUint64, value)
	}
	if value, ok := druo.mutation.AddedWeekID(); ok {
		_spec.AddField(dailyroutine.FieldWeekID, field.TypeUint64, value)
	}
	if druo.mutation.WeekIDCleared() {
		_spec.ClearField(dailyroutine.FieldWeekID, field.TypeUint64)
	}
	if value, ok := druo.mutation.Day(); ok {
		_spec.SetField(dailyroutine.FieldDay, field.TypeInt, value)
	}
	if value, ok := druo.mutation.AddedDay(); ok {
		_spec.AddField(dailyroutine.FieldDay, field.TypeInt, value)
	}
	if druo.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.ProgramTable,
			Columns: dailyroutine.ProgramPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.RemovedProgramIDs(); len(nodes) > 0 && !druo.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.ProgramTable,
			Columns: dailyroutine.ProgramPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.ProgramTable,
			Columns: dailyroutine.ProgramPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if druo.mutation.WeeklyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.WeeklyRoutineTable,
			Columns: dailyroutine.WeeklyRoutinePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.RemovedWeeklyRoutineIDs(); len(nodes) > 0 && !druo.mutation.WeeklyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.WeeklyRoutineTable,
			Columns: dailyroutine.WeeklyRoutinePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.WeeklyRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dailyroutine.WeeklyRoutineTable,
			Columns: dailyroutine.WeeklyRoutinePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if druo.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.RemovedRoutineActsIDs(); len(nodes) > 0 && !druo.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DailyRoutine{config: druo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, druo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyroutine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	druo.mutation.done = true
	return _node, nil
}
