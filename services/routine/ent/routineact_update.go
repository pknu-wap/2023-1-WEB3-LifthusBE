// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/act"
	"routine/ent/dailyroutine"
	"routine/ent/predicate"
	"routine/ent/routineact"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineActUpdate is the builder for updating RoutineAct entities.
type RoutineActUpdate struct {
	config
	hooks    []Hook
	mutation *RoutineActMutation
}

// Where appends a list predicates to the RoutineActUpdate builder.
func (rau *RoutineActUpdate) Where(ps ...predicate.RoutineAct) *RoutineActUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetDailyRoutineID sets the "daily_routine_id" field.
func (rau *RoutineActUpdate) SetDailyRoutineID(u uint64) *RoutineActUpdate {
	rau.mutation.ResetDailyRoutineID()
	rau.mutation.SetDailyRoutineID(u)
	return rau
}

// AddDailyRoutineID adds u to the "daily_routine_id" field.
func (rau *RoutineActUpdate) AddDailyRoutineID(u int64) *RoutineActUpdate {
	rau.mutation.AddDailyRoutineID(u)
	return rau
}

// SetActID sets the "act_id" field.
func (rau *RoutineActUpdate) SetActID(u uint64) *RoutineActUpdate {
	rau.mutation.ResetActID()
	rau.mutation.SetActID(u)
	return rau
}

// AddActID adds u to the "act_id" field.
func (rau *RoutineActUpdate) AddActID(u int64) *RoutineActUpdate {
	rau.mutation.AddActID(u)
	return rau
}

// SetOrder sets the "order" field.
func (rau *RoutineActUpdate) SetOrder(i int) *RoutineActUpdate {
	rau.mutation.ResetOrder()
	rau.mutation.SetOrder(i)
	return rau
}

// AddOrder adds i to the "order" field.
func (rau *RoutineActUpdate) AddOrder(i int) *RoutineActUpdate {
	rau.mutation.AddOrder(i)
	return rau
}

// SetReps sets the "reps" field.
func (rau *RoutineActUpdate) SetReps(i int) *RoutineActUpdate {
	rau.mutation.ResetReps()
	rau.mutation.SetReps(i)
	return rau
}

// SetNillableReps sets the "reps" field if the given value is not nil.
func (rau *RoutineActUpdate) SetNillableReps(i *int) *RoutineActUpdate {
	if i != nil {
		rau.SetReps(*i)
	}
	return rau
}

// AddReps adds i to the "reps" field.
func (rau *RoutineActUpdate) AddReps(i int) *RoutineActUpdate {
	rau.mutation.AddReps(i)
	return rau
}

// ClearReps clears the value of the "reps" field.
func (rau *RoutineActUpdate) ClearReps() *RoutineActUpdate {
	rau.mutation.ClearReps()
	return rau
}

// SetLap sets the "lap" field.
func (rau *RoutineActUpdate) SetLap(i int) *RoutineActUpdate {
	rau.mutation.ResetLap()
	rau.mutation.SetLap(i)
	return rau
}

// SetNillableLap sets the "lap" field if the given value is not nil.
func (rau *RoutineActUpdate) SetNillableLap(i *int) *RoutineActUpdate {
	if i != nil {
		rau.SetLap(*i)
	}
	return rau
}

// AddLap adds i to the "lap" field.
func (rau *RoutineActUpdate) AddLap(i int) *RoutineActUpdate {
	rau.mutation.AddLap(i)
	return rau
}

// ClearLap clears the value of the "lap" field.
func (rau *RoutineActUpdate) ClearLap() *RoutineActUpdate {
	rau.mutation.ClearLap()
	return rau
}

// SetActID sets the "act" edge to the Act entity by ID.
func (rau *RoutineActUpdate) SetActID(id uint64) *RoutineActUpdate {
	rau.mutation.SetActID(id)
	return rau
}

// SetAct sets the "act" edge to the Act entity.
func (rau *RoutineActUpdate) SetAct(a *Act) *RoutineActUpdate {
	return rau.SetActID(a.ID)
}

// SetDailyRoutineID sets the "daily_routine" edge to the DailyRoutine entity by ID.
func (rau *RoutineActUpdate) SetDailyRoutineID(id uint64) *RoutineActUpdate {
	rau.mutation.SetDailyRoutineID(id)
	return rau
}

// SetDailyRoutine sets the "daily_routine" edge to the DailyRoutine entity.
func (rau *RoutineActUpdate) SetDailyRoutine(d *DailyRoutine) *RoutineActUpdate {
	return rau.SetDailyRoutineID(d.ID)
}

// Mutation returns the RoutineActMutation object of the builder.
func (rau *RoutineActUpdate) Mutation() *RoutineActMutation {
	return rau.mutation
}

// ClearAct clears the "act" edge to the Act entity.
func (rau *RoutineActUpdate) ClearAct() *RoutineActUpdate {
	rau.mutation.ClearAct()
	return rau
}

// ClearDailyRoutine clears the "daily_routine" edge to the DailyRoutine entity.
func (rau *RoutineActUpdate) ClearDailyRoutine() *RoutineActUpdate {
	rau.mutation.ClearDailyRoutine()
	return rau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *RoutineActUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rau.sqlSave, rau.mutation, rau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rau *RoutineActUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *RoutineActUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *RoutineActUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rau *RoutineActUpdate) check() error {
	if v, ok := rau.mutation.Order(); ok {
		if err := routineact.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.order": %w`, err)}
		}
	}
	if v, ok := rau.mutation.Reps(); ok {
		if err := routineact.RepsValidator(v); err != nil {
			return &ValidationError{Name: "reps", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.reps": %w`, err)}
		}
	}
	if v, ok := rau.mutation.Lap(); ok {
		if err := routineact.LapValidator(v); err != nil {
			return &ValidationError{Name: "lap", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.lap": %w`, err)}
		}
	}
	if _, ok := rau.mutation.ActID(); rau.mutation.ActCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineAct.act"`)
	}
	if _, ok := rau.mutation.DailyRoutineID(); rau.mutation.DailyRoutineCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineAct.daily_routine"`)
	}
	return nil
}

func (rau *RoutineActUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(routineact.Table, routineact.Columns, sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64))
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.DailyRoutineID(); ok {
		_spec.SetField(routineact.FieldDailyRoutineID, field.TypeUint64, value)
	}
	if value, ok := rau.mutation.AddedDailyRoutineID(); ok {
		_spec.AddField(routineact.FieldDailyRoutineID, field.TypeUint64, value)
	}
	if value, ok := rau.mutation.ActID(); ok {
		_spec.SetField(routineact.FieldActID, field.TypeUint64, value)
	}
	if value, ok := rau.mutation.AddedActID(); ok {
		_spec.AddField(routineact.FieldActID, field.TypeUint64, value)
	}
	if value, ok := rau.mutation.Order(); ok {
		_spec.SetField(routineact.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rau.mutation.AddedOrder(); ok {
		_spec.AddField(routineact.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rau.mutation.Reps(); ok {
		_spec.SetField(routineact.FieldReps, field.TypeInt, value)
	}
	if value, ok := rau.mutation.AddedReps(); ok {
		_spec.AddField(routineact.FieldReps, field.TypeInt, value)
	}
	if rau.mutation.RepsCleared() {
		_spec.ClearField(routineact.FieldReps, field.TypeInt)
	}
	if value, ok := rau.mutation.Lap(); ok {
		_spec.SetField(routineact.FieldLap, field.TypeInt, value)
	}
	if value, ok := rau.mutation.AddedLap(); ok {
		_spec.AddField(routineact.FieldLap, field.TypeInt, value)
	}
	if rau.mutation.LapCleared() {
		_spec.ClearField(routineact.FieldLap, field.TypeInt)
	}
	if rau.mutation.ActCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.ActTable,
			Columns: []string{routineact.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.ActTable,
			Columns: []string{routineact.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rau.mutation.DailyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.DailyRoutineTable,
			Columns: []string{routineact.DailyRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.DailyRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.DailyRoutineTable,
			Columns: []string{routineact.DailyRoutineColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routineact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rau.mutation.done = true
	return n, nil
}

// RoutineActUpdateOne is the builder for updating a single RoutineAct entity.
type RoutineActUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoutineActMutation
}

// SetDailyRoutineID sets the "daily_routine_id" field.
func (rauo *RoutineActUpdateOne) SetDailyRoutineID(u uint64) *RoutineActUpdateOne {
	rauo.mutation.ResetDailyRoutineID()
	rauo.mutation.SetDailyRoutineID(u)
	return rauo
}

// AddDailyRoutineID adds u to the "daily_routine_id" field.
func (rauo *RoutineActUpdateOne) AddDailyRoutineID(u int64) *RoutineActUpdateOne {
	rauo.mutation.AddDailyRoutineID(u)
	return rauo
}

// SetActID sets the "act_id" field.
func (rauo *RoutineActUpdateOne) SetActID(u uint64) *RoutineActUpdateOne {
	rauo.mutation.ResetActID()
	rauo.mutation.SetActID(u)
	return rauo
}

// AddActID adds u to the "act_id" field.
func (rauo *RoutineActUpdateOne) AddActID(u int64) *RoutineActUpdateOne {
	rauo.mutation.AddActID(u)
	return rauo
}

// SetOrder sets the "order" field.
func (rauo *RoutineActUpdateOne) SetOrder(i int) *RoutineActUpdateOne {
	rauo.mutation.ResetOrder()
	rauo.mutation.SetOrder(i)
	return rauo
}

// AddOrder adds i to the "order" field.
func (rauo *RoutineActUpdateOne) AddOrder(i int) *RoutineActUpdateOne {
	rauo.mutation.AddOrder(i)
	return rauo
}

// SetReps sets the "reps" field.
func (rauo *RoutineActUpdateOne) SetReps(i int) *RoutineActUpdateOne {
	rauo.mutation.ResetReps()
	rauo.mutation.SetReps(i)
	return rauo
}

// SetNillableReps sets the "reps" field if the given value is not nil.
func (rauo *RoutineActUpdateOne) SetNillableReps(i *int) *RoutineActUpdateOne {
	if i != nil {
		rauo.SetReps(*i)
	}
	return rauo
}

// AddReps adds i to the "reps" field.
func (rauo *RoutineActUpdateOne) AddReps(i int) *RoutineActUpdateOne {
	rauo.mutation.AddReps(i)
	return rauo
}

// ClearReps clears the value of the "reps" field.
func (rauo *RoutineActUpdateOne) ClearReps() *RoutineActUpdateOne {
	rauo.mutation.ClearReps()
	return rauo
}

// SetLap sets the "lap" field.
func (rauo *RoutineActUpdateOne) SetLap(i int) *RoutineActUpdateOne {
	rauo.mutation.ResetLap()
	rauo.mutation.SetLap(i)
	return rauo
}

// SetNillableLap sets the "lap" field if the given value is not nil.
func (rauo *RoutineActUpdateOne) SetNillableLap(i *int) *RoutineActUpdateOne {
	if i != nil {
		rauo.SetLap(*i)
	}
	return rauo
}

// AddLap adds i to the "lap" field.
func (rauo *RoutineActUpdateOne) AddLap(i int) *RoutineActUpdateOne {
	rauo.mutation.AddLap(i)
	return rauo
}

// ClearLap clears the value of the "lap" field.
func (rauo *RoutineActUpdateOne) ClearLap() *RoutineActUpdateOne {
	rauo.mutation.ClearLap()
	return rauo
}

// SetActID sets the "act" edge to the Act entity by ID.
func (rauo *RoutineActUpdateOne) SetActID(id uint64) *RoutineActUpdateOne {
	rauo.mutation.SetActID(id)
	return rauo
}

// SetAct sets the "act" edge to the Act entity.
func (rauo *RoutineActUpdateOne) SetAct(a *Act) *RoutineActUpdateOne {
	return rauo.SetActID(a.ID)
}

// SetDailyRoutineID sets the "daily_routine" edge to the DailyRoutine entity by ID.
func (rauo *RoutineActUpdateOne) SetDailyRoutineID(id uint64) *RoutineActUpdateOne {
	rauo.mutation.SetDailyRoutineID(id)
	return rauo
}

// SetDailyRoutine sets the "daily_routine" edge to the DailyRoutine entity.
func (rauo *RoutineActUpdateOne) SetDailyRoutine(d *DailyRoutine) *RoutineActUpdateOne {
	return rauo.SetDailyRoutineID(d.ID)
}

// Mutation returns the RoutineActMutation object of the builder.
func (rauo *RoutineActUpdateOne) Mutation() *RoutineActMutation {
	return rauo.mutation
}

// ClearAct clears the "act" edge to the Act entity.
func (rauo *RoutineActUpdateOne) ClearAct() *RoutineActUpdateOne {
	rauo.mutation.ClearAct()
	return rauo
}

// ClearDailyRoutine clears the "daily_routine" edge to the DailyRoutine entity.
func (rauo *RoutineActUpdateOne) ClearDailyRoutine() *RoutineActUpdateOne {
	rauo.mutation.ClearDailyRoutine()
	return rauo
}

// Where appends a list predicates to the RoutineActUpdate builder.
func (rauo *RoutineActUpdateOne) Where(ps ...predicate.RoutineAct) *RoutineActUpdateOne {
	rauo.mutation.Where(ps...)
	return rauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *RoutineActUpdateOne) Select(field string, fields ...string) *RoutineActUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated RoutineAct entity.
func (rauo *RoutineActUpdateOne) Save(ctx context.Context) (*RoutineAct, error) {
	return withHooks(ctx, rauo.sqlSave, rauo.mutation, rauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *RoutineActUpdateOne) SaveX(ctx context.Context) *RoutineAct {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *RoutineActUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *RoutineActUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rauo *RoutineActUpdateOne) check() error {
	if v, ok := rauo.mutation.Order(); ok {
		if err := routineact.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.order": %w`, err)}
		}
	}
	if v, ok := rauo.mutation.Reps(); ok {
		if err := routineact.RepsValidator(v); err != nil {
			return &ValidationError{Name: "reps", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.reps": %w`, err)}
		}
	}
	if v, ok := rauo.mutation.Lap(); ok {
		if err := routineact.LapValidator(v); err != nil {
			return &ValidationError{Name: "lap", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.lap": %w`, err)}
		}
	}
	if _, ok := rauo.mutation.ActID(); rauo.mutation.ActCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineAct.act"`)
	}
	if _, ok := rauo.mutation.DailyRoutineID(); rauo.mutation.DailyRoutineCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineAct.daily_routine"`)
	}
	return nil
}

func (rauo *RoutineActUpdateOne) sqlSave(ctx context.Context) (_node *RoutineAct, err error) {
	if err := rauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(routineact.Table, routineact.Columns, sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64))
	id, ok := rauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoutineAct.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routineact.FieldID)
		for _, f := range fields {
			if !routineact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routineact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.DailyRoutineID(); ok {
		_spec.SetField(routineact.FieldDailyRoutineID, field.TypeUint64, value)
	}
	if value, ok := rauo.mutation.AddedDailyRoutineID(); ok {
		_spec.AddField(routineact.FieldDailyRoutineID, field.TypeUint64, value)
	}
	if value, ok := rauo.mutation.ActID(); ok {
		_spec.SetField(routineact.FieldActID, field.TypeUint64, value)
	}
	if value, ok := rauo.mutation.AddedActID(); ok {
		_spec.AddField(routineact.FieldActID, field.TypeUint64, value)
	}
	if value, ok := rauo.mutation.Order(); ok {
		_spec.SetField(routineact.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.AddedOrder(); ok {
		_spec.AddField(routineact.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.Reps(); ok {
		_spec.SetField(routineact.FieldReps, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.AddedReps(); ok {
		_spec.AddField(routineact.FieldReps, field.TypeInt, value)
	}
	if rauo.mutation.RepsCleared() {
		_spec.ClearField(routineact.FieldReps, field.TypeInt)
	}
	if value, ok := rauo.mutation.Lap(); ok {
		_spec.SetField(routineact.FieldLap, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.AddedLap(); ok {
		_spec.AddField(routineact.FieldLap, field.TypeInt, value)
	}
	if rauo.mutation.LapCleared() {
		_spec.ClearField(routineact.FieldLap, field.TypeInt)
	}
	if rauo.mutation.ActCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.ActTable,
			Columns: []string{routineact.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.ActTable,
			Columns: []string{routineact.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rauo.mutation.DailyRoutineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.DailyRoutineTable,
			Columns: []string{routineact.DailyRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.DailyRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.DailyRoutineTable,
			Columns: []string{routineact.DailyRoutineColumn},
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
	_node = &RoutineAct{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routineact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rauo.mutation.done = true
	return _node, nil
}
