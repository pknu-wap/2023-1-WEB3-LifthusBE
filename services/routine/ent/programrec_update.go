// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/predicate"
	"routine/ent/programrec"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramRecUpdate is the builder for updating ProgramRec entities.
type ProgramRecUpdate struct {
	config
	hooks    []Hook
	mutation *ProgramRecMutation
}

// Where appends a list predicates to the ProgramRecUpdate builder.
func (pru *ProgramRecUpdate) Where(ps ...predicate.ProgramRec) *ProgramRecUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetAuthor sets the "author" field.
func (pru *ProgramRecUpdate) SetAuthor(u uint64) *ProgramRecUpdate {
	pru.mutation.ResetAuthor()
	pru.mutation.SetAuthor(u)
	return pru
}

// AddAuthor adds u to the "author" field.
func (pru *ProgramRecUpdate) AddAuthor(u int64) *ProgramRecUpdate {
	pru.mutation.AddAuthor(u)
	return pru
}

// SetProgramID sets the "program_id" field.
func (pru *ProgramRecUpdate) SetProgramID(u uint64) *ProgramRecUpdate {
	pru.mutation.ResetProgramID()
	pru.mutation.SetProgramID(u)
	return pru
}

// AddProgramID adds u to the "program_id" field.
func (pru *ProgramRecUpdate) AddProgramID(u int64) *ProgramRecUpdate {
	pru.mutation.AddProgramID(u)
	return pru
}

// SetStartDate sets the "start_date" field.
func (pru *ProgramRecUpdate) SetStartDate(t time.Time) *ProgramRecUpdate {
	pru.mutation.SetStartDate(t)
	return pru
}

// SetEndDate sets the "end_date" field.
func (pru *ProgramRecUpdate) SetEndDate(t time.Time) *ProgramRecUpdate {
	pru.mutation.SetEndDate(t)
	return pru
}

// SetStatus sets the "status" field.
func (pru *ProgramRecUpdate) SetStatus(pr programrec.Status) *ProgramRecUpdate {
	pru.mutation.SetStatus(pr)
	return pru
}

// SetComment sets the "comment" field.
func (pru *ProgramRecUpdate) SetComment(s string) *ProgramRecUpdate {
	pru.mutation.SetComment(s)
	return pru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (pru *ProgramRecUpdate) SetNillableComment(s *string) *ProgramRecUpdate {
	if s != nil {
		pru.SetComment(*s)
	}
	return pru
}

// ClearComment clears the value of the "comment" field.
func (pru *ProgramRecUpdate) ClearComment() *ProgramRecUpdate {
	pru.mutation.ClearComment()
	return pru
}

// SetUpdatedAt sets the "updated_at" field.
func (pru *ProgramRecUpdate) SetUpdatedAt(t time.Time) *ProgramRecUpdate {
	pru.mutation.SetUpdatedAt(t)
	return pru
}

// Mutation returns the ProgramRecMutation object of the builder.
func (pru *ProgramRecUpdate) Mutation() *ProgramRecMutation {
	return pru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *ProgramRecUpdate) Save(ctx context.Context) (int, error) {
	pru.defaults()
	return withHooks(ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *ProgramRecUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *ProgramRecUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *ProgramRecUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pru *ProgramRecUpdate) defaults() {
	if _, ok := pru.mutation.UpdatedAt(); !ok {
		v := programrec.UpdateDefaultUpdatedAt()
		pru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pru *ProgramRecUpdate) check() error {
	if v, ok := pru.mutation.Status(); ok {
		if err := programrec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProgramRec.status": %w`, err)}
		}
	}
	return nil
}

func (pru *ProgramRecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(programrec.Table, programrec.Columns, sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pru.mutation.Author(); ok {
		_spec.SetField(programrec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := pru.mutation.AddedAuthor(); ok {
		_spec.AddField(programrec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := pru.mutation.ProgramID(); ok {
		_spec.SetField(programrec.FieldProgramID, field.TypeUint64, value)
	}
	if value, ok := pru.mutation.AddedProgramID(); ok {
		_spec.AddField(programrec.FieldProgramID, field.TypeUint64, value)
	}
	if value, ok := pru.mutation.StartDate(); ok {
		_spec.SetField(programrec.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := pru.mutation.EndDate(); ok {
		_spec.SetField(programrec.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := pru.mutation.Status(); ok {
		_spec.SetField(programrec.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pru.mutation.Comment(); ok {
		_spec.SetField(programrec.FieldComment, field.TypeString, value)
	}
	if pru.mutation.CommentCleared() {
		_spec.ClearField(programrec.FieldComment, field.TypeString)
	}
	if value, ok := pru.mutation.UpdatedAt(); ok {
		_spec.SetField(programrec.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programrec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// ProgramRecUpdateOne is the builder for updating a single ProgramRec entity.
type ProgramRecUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProgramRecMutation
}

// SetAuthor sets the "author" field.
func (pruo *ProgramRecUpdateOne) SetAuthor(u uint64) *ProgramRecUpdateOne {
	pruo.mutation.ResetAuthor()
	pruo.mutation.SetAuthor(u)
	return pruo
}

// AddAuthor adds u to the "author" field.
func (pruo *ProgramRecUpdateOne) AddAuthor(u int64) *ProgramRecUpdateOne {
	pruo.mutation.AddAuthor(u)
	return pruo
}

// SetProgramID sets the "program_id" field.
func (pruo *ProgramRecUpdateOne) SetProgramID(u uint64) *ProgramRecUpdateOne {
	pruo.mutation.ResetProgramID()
	pruo.mutation.SetProgramID(u)
	return pruo
}

// AddProgramID adds u to the "program_id" field.
func (pruo *ProgramRecUpdateOne) AddProgramID(u int64) *ProgramRecUpdateOne {
	pruo.mutation.AddProgramID(u)
	return pruo
}

// SetStartDate sets the "start_date" field.
func (pruo *ProgramRecUpdateOne) SetStartDate(t time.Time) *ProgramRecUpdateOne {
	pruo.mutation.SetStartDate(t)
	return pruo
}

// SetEndDate sets the "end_date" field.
func (pruo *ProgramRecUpdateOne) SetEndDate(t time.Time) *ProgramRecUpdateOne {
	pruo.mutation.SetEndDate(t)
	return pruo
}

// SetStatus sets the "status" field.
func (pruo *ProgramRecUpdateOne) SetStatus(pr programrec.Status) *ProgramRecUpdateOne {
	pruo.mutation.SetStatus(pr)
	return pruo
}

// SetComment sets the "comment" field.
func (pruo *ProgramRecUpdateOne) SetComment(s string) *ProgramRecUpdateOne {
	pruo.mutation.SetComment(s)
	return pruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (pruo *ProgramRecUpdateOne) SetNillableComment(s *string) *ProgramRecUpdateOne {
	if s != nil {
		pruo.SetComment(*s)
	}
	return pruo
}

// ClearComment clears the value of the "comment" field.
func (pruo *ProgramRecUpdateOne) ClearComment() *ProgramRecUpdateOne {
	pruo.mutation.ClearComment()
	return pruo
}

// SetUpdatedAt sets the "updated_at" field.
func (pruo *ProgramRecUpdateOne) SetUpdatedAt(t time.Time) *ProgramRecUpdateOne {
	pruo.mutation.SetUpdatedAt(t)
	return pruo
}

// Mutation returns the ProgramRecMutation object of the builder.
func (pruo *ProgramRecUpdateOne) Mutation() *ProgramRecMutation {
	return pruo.mutation
}

// Where appends a list predicates to the ProgramRecUpdate builder.
func (pruo *ProgramRecUpdateOne) Where(ps ...predicate.ProgramRec) *ProgramRecUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *ProgramRecUpdateOne) Select(field string, fields ...string) *ProgramRecUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated ProgramRec entity.
func (pruo *ProgramRecUpdateOne) Save(ctx context.Context) (*ProgramRec, error) {
	pruo.defaults()
	return withHooks(ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *ProgramRecUpdateOne) SaveX(ctx context.Context) *ProgramRec {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *ProgramRecUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *ProgramRecUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pruo *ProgramRecUpdateOne) defaults() {
	if _, ok := pruo.mutation.UpdatedAt(); !ok {
		v := programrec.UpdateDefaultUpdatedAt()
		pruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pruo *ProgramRecUpdateOne) check() error {
	if v, ok := pruo.mutation.Status(); ok {
		if err := programrec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProgramRec.status": %w`, err)}
		}
	}
	return nil
}

func (pruo *ProgramRecUpdateOne) sqlSave(ctx context.Context) (_node *ProgramRec, err error) {
	if err := pruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(programrec.Table, programrec.Columns, sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProgramRec.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, programrec.FieldID)
		for _, f := range fields {
			if !programrec.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != programrec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pruo.mutation.Author(); ok {
		_spec.SetField(programrec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := pruo.mutation.AddedAuthor(); ok {
		_spec.AddField(programrec.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := pruo.mutation.ProgramID(); ok {
		_spec.SetField(programrec.FieldProgramID, field.TypeUint64, value)
	}
	if value, ok := pruo.mutation.AddedProgramID(); ok {
		_spec.AddField(programrec.FieldProgramID, field.TypeUint64, value)
	}
	if value, ok := pruo.mutation.StartDate(); ok {
		_spec.SetField(programrec.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := pruo.mutation.EndDate(); ok {
		_spec.SetField(programrec.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := pruo.mutation.Status(); ok {
		_spec.SetField(programrec.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pruo.mutation.Comment(); ok {
		_spec.SetField(programrec.FieldComment, field.TypeString, value)
	}
	if pruo.mutation.CommentCleared() {
		_spec.ClearField(programrec.FieldComment, field.TypeString)
	}
	if value, ok := pruo.mutation.UpdatedAt(); ok {
		_spec.SetField(programrec.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &ProgramRec{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programrec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}
