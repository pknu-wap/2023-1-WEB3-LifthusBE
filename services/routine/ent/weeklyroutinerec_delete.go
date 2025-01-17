// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"routine/ent/predicate"
	"routine/ent/weeklyroutinerec"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WeeklyRoutineRecDelete is the builder for deleting a WeeklyRoutineRec entity.
type WeeklyRoutineRecDelete struct {
	config
	hooks    []Hook
	mutation *WeeklyRoutineRecMutation
}

// Where appends a list predicates to the WeeklyRoutineRecDelete builder.
func (wrrd *WeeklyRoutineRecDelete) Where(ps ...predicate.WeeklyRoutineRec) *WeeklyRoutineRecDelete {
	wrrd.mutation.Where(ps...)
	return wrrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wrrd *WeeklyRoutineRecDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wrrd.sqlExec, wrrd.mutation, wrrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wrrd *WeeklyRoutineRecDelete) ExecX(ctx context.Context) int {
	n, err := wrrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wrrd *WeeklyRoutineRecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(weeklyroutinerec.Table, sqlgraph.NewFieldSpec(weeklyroutinerec.FieldID, field.TypeUint64))
	if ps := wrrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wrrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wrrd.mutation.done = true
	return affected, err
}

// WeeklyRoutineRecDeleteOne is the builder for deleting a single WeeklyRoutineRec entity.
type WeeklyRoutineRecDeleteOne struct {
	wrrd *WeeklyRoutineRecDelete
}

// Where appends a list predicates to the WeeklyRoutineRecDelete builder.
func (wrrdo *WeeklyRoutineRecDeleteOne) Where(ps ...predicate.WeeklyRoutineRec) *WeeklyRoutineRecDeleteOne {
	wrrdo.wrrd.mutation.Where(ps...)
	return wrrdo
}

// Exec executes the deletion query.
func (wrrdo *WeeklyRoutineRecDeleteOne) Exec(ctx context.Context) error {
	n, err := wrrdo.wrrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{weeklyroutinerec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wrrdo *WeeklyRoutineRecDeleteOne) ExecX(ctx context.Context) {
	if err := wrrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
