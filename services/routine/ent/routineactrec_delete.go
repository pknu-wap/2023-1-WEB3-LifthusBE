// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"routine/ent/predicate"
	"routine/ent/routineactrec"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineActRecDelete is the builder for deleting a RoutineActRec entity.
type RoutineActRecDelete struct {
	config
	hooks    []Hook
	mutation *RoutineActRecMutation
}

// Where appends a list predicates to the RoutineActRecDelete builder.
func (rard *RoutineActRecDelete) Where(ps ...predicate.RoutineActRec) *RoutineActRecDelete {
	rard.mutation.Where(ps...)
	return rard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rard *RoutineActRecDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rard.sqlExec, rard.mutation, rard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rard *RoutineActRecDelete) ExecX(ctx context.Context) int {
	n, err := rard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rard *RoutineActRecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(routineactrec.Table, sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64))
	if ps := rard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rard.mutation.done = true
	return affected, err
}

// RoutineActRecDeleteOne is the builder for deleting a single RoutineActRec entity.
type RoutineActRecDeleteOne struct {
	rard *RoutineActRecDelete
}

// Where appends a list predicates to the RoutineActRecDelete builder.
func (rardo *RoutineActRecDeleteOne) Where(ps ...predicate.RoutineActRec) *RoutineActRecDeleteOne {
	rardo.rard.mutation.Where(ps...)
	return rardo
}

// Exec executes the deletion query.
func (rardo *RoutineActRecDeleteOne) Exec(ctx context.Context) error {
	n, err := rardo.rard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{routineactrec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rardo *RoutineActRecDeleteOne) ExecX(ctx context.Context) {
	if err := rardo.Exec(ctx); err != nil {
		panic(err)
	}
}
