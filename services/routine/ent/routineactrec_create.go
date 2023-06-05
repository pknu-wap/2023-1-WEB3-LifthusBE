// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/act"
	"routine/ent/dailyroutinerec"
	"routine/ent/routineact"
	"routine/ent/routineactrec"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineActRecCreate is the builder for creating a RoutineActRec entity.
type RoutineActRecCreate struct {
	config
	mutation *RoutineActRecMutation
	hooks    []Hook
}

// SetAuthor sets the "author" field.
func (rarc *RoutineActRecCreate) SetAuthor(u uint64) *RoutineActRecCreate {
	rarc.mutation.SetAuthor(u)
	return rarc
}

// SetDailyRoutineRecID sets the "daily_routine_rec_id" field.
func (rarc *RoutineActRecCreate) SetDailyRoutineRecID(u uint64) *RoutineActRecCreate {
	rarc.mutation.SetDailyRoutineRecID(u)
	return rarc
}

// SetRoutineActID sets the "routine_act_id" field.
func (rarc *RoutineActRecCreate) SetRoutineActID(u uint64) *RoutineActRecCreate {
	rarc.mutation.SetRoutineActID(u)
	return rarc
}

// SetNillableRoutineActID sets the "routine_act_id" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableRoutineActID(u *uint64) *RoutineActRecCreate {
	if u != nil {
		rarc.SetRoutineActID(*u)
	}
	return rarc
}

// SetActID sets the "act_id" field.
func (rarc *RoutineActRecCreate) SetActID(u uint64) *RoutineActRecCreate {
	rarc.mutation.SetActID(u)
	return rarc
}

// SetOrder sets the "order" field.
func (rarc *RoutineActRecCreate) SetOrder(i int) *RoutineActRecCreate {
	rarc.mutation.SetOrder(i)
	return rarc
}

// SetReps sets the "reps" field.
func (rarc *RoutineActRecCreate) SetReps(i int) *RoutineActRecCreate {
	rarc.mutation.SetReps(i)
	return rarc
}

// SetNillableReps sets the "reps" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableReps(i *int) *RoutineActRecCreate {
	if i != nil {
		rarc.SetReps(*i)
	}
	return rarc
}

// SetLap sets the "lap" field.
func (rarc *RoutineActRecCreate) SetLap(i int) *RoutineActRecCreate {
	rarc.mutation.SetLap(i)
	return rarc
}

// SetNillableLap sets the "lap" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableLap(i *int) *RoutineActRecCreate {
	if i != nil {
		rarc.SetLap(*i)
	}
	return rarc
}

// SetCurrentReps sets the "current_reps" field.
func (rarc *RoutineActRecCreate) SetCurrentReps(i int) *RoutineActRecCreate {
	rarc.mutation.SetCurrentReps(i)
	return rarc
}

// SetNillableCurrentReps sets the "current_reps" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableCurrentReps(i *int) *RoutineActRecCreate {
	if i != nil {
		rarc.SetCurrentReps(*i)
	}
	return rarc
}

// SetCurrentLap sets the "current_lap" field.
func (rarc *RoutineActRecCreate) SetCurrentLap(i int) *RoutineActRecCreate {
	rarc.mutation.SetCurrentLap(i)
	return rarc
}

// SetNillableCurrentLap sets the "current_lap" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableCurrentLap(i *int) *RoutineActRecCreate {
	if i != nil {
		rarc.SetCurrentLap(*i)
	}
	return rarc
}

// SetStartedAt sets the "started_at" field.
func (rarc *RoutineActRecCreate) SetStartedAt(t time.Time) *RoutineActRecCreate {
	rarc.mutation.SetStartedAt(t)
	return rarc
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableStartedAt(t *time.Time) *RoutineActRecCreate {
	if t != nil {
		rarc.SetStartedAt(*t)
	}
	return rarc
}

// SetImage sets the "image" field.
func (rarc *RoutineActRecCreate) SetImage(s string) *RoutineActRecCreate {
	rarc.mutation.SetImage(s)
	return rarc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableImage(s *string) *RoutineActRecCreate {
	if s != nil {
		rarc.SetImage(*s)
	}
	return rarc
}

// SetComment sets the "comment" field.
func (rarc *RoutineActRecCreate) SetComment(s string) *RoutineActRecCreate {
	rarc.mutation.SetComment(s)
	return rarc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableComment(s *string) *RoutineActRecCreate {
	if s != nil {
		rarc.SetComment(*s)
	}
	return rarc
}

// SetStatus sets the "status" field.
func (rarc *RoutineActRecCreate) SetStatus(r routineactrec.Status) *RoutineActRecCreate {
	rarc.mutation.SetStatus(r)
	return rarc
}

// SetCreatedAt sets the "created_at" field.
func (rarc *RoutineActRecCreate) SetCreatedAt(t time.Time) *RoutineActRecCreate {
	rarc.mutation.SetCreatedAt(t)
	return rarc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableCreatedAt(t *time.Time) *RoutineActRecCreate {
	if t != nil {
		rarc.SetCreatedAt(*t)
	}
	return rarc
}

// SetUpdatedAt sets the "updated_at" field.
func (rarc *RoutineActRecCreate) SetUpdatedAt(t time.Time) *RoutineActRecCreate {
	rarc.mutation.SetUpdatedAt(t)
	return rarc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rarc *RoutineActRecCreate) SetNillableUpdatedAt(t *time.Time) *RoutineActRecCreate {
	if t != nil {
		rarc.SetUpdatedAt(*t)
	}
	return rarc
}

// SetID sets the "id" field.
func (rarc *RoutineActRecCreate) SetID(u uint64) *RoutineActRecCreate {
	rarc.mutation.SetID(u)
	return rarc
}

// SetDailyRoutineRec sets the "daily_routine_rec" edge to the DailyRoutineRec entity.
func (rarc *RoutineActRecCreate) SetDailyRoutineRec(d *DailyRoutineRec) *RoutineActRecCreate {
	return rarc.SetDailyRoutineRecID(d.ID)
}

// SetAct sets the "act" edge to the Act entity.
func (rarc *RoutineActRecCreate) SetAct(a *Act) *RoutineActRecCreate {
	return rarc.SetActID(a.ID)
}

// SetRoutineAct sets the "routine_act" edge to the RoutineAct entity.
func (rarc *RoutineActRecCreate) SetRoutineAct(r *RoutineAct) *RoutineActRecCreate {
	return rarc.SetRoutineActID(r.ID)
}

// Mutation returns the RoutineActRecMutation object of the builder.
func (rarc *RoutineActRecCreate) Mutation() *RoutineActRecMutation {
	return rarc.mutation
}

// Save creates the RoutineActRec in the database.
func (rarc *RoutineActRecCreate) Save(ctx context.Context) (*RoutineActRec, error) {
	rarc.defaults()
	return withHooks(ctx, rarc.sqlSave, rarc.mutation, rarc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rarc *RoutineActRecCreate) SaveX(ctx context.Context) *RoutineActRec {
	v, err := rarc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rarc *RoutineActRecCreate) Exec(ctx context.Context) error {
	_, err := rarc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rarc *RoutineActRecCreate) ExecX(ctx context.Context) {
	if err := rarc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rarc *RoutineActRecCreate) defaults() {
	if _, ok := rarc.mutation.CurrentReps(); !ok {
		v := routineactrec.DefaultCurrentReps
		rarc.mutation.SetCurrentReps(v)
	}
	if _, ok := rarc.mutation.CurrentLap(); !ok {
		v := routineactrec.DefaultCurrentLap
		rarc.mutation.SetCurrentLap(v)
	}
	if _, ok := rarc.mutation.CreatedAt(); !ok {
		v := routineactrec.DefaultCreatedAt()
		rarc.mutation.SetCreatedAt(v)
	}
	if _, ok := rarc.mutation.UpdatedAt(); !ok {
		v := routineactrec.DefaultUpdatedAt()
		rarc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rarc *RoutineActRecCreate) check() error {
	if _, ok := rarc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "RoutineActRec.author"`)}
	}
	if _, ok := rarc.mutation.DailyRoutineRecID(); !ok {
		return &ValidationError{Name: "daily_routine_rec_id", err: errors.New(`ent: missing required field "RoutineActRec.daily_routine_rec_id"`)}
	}
	if _, ok := rarc.mutation.ActID(); !ok {
		return &ValidationError{Name: "act_id", err: errors.New(`ent: missing required field "RoutineActRec.act_id"`)}
	}
	if _, ok := rarc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "RoutineActRec.order"`)}
	}
	if v, ok := rarc.mutation.Order(); ok {
		if err := routineactrec.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.order": %w`, err)}
		}
	}
	if v, ok := rarc.mutation.Reps(); ok {
		if err := routineactrec.RepsValidator(v); err != nil {
			return &ValidationError{Name: "reps", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.reps": %w`, err)}
		}
	}
	if v, ok := rarc.mutation.Lap(); ok {
		if err := routineactrec.LapValidator(v); err != nil {
			return &ValidationError{Name: "lap", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.lap": %w`, err)}
		}
	}
	if _, ok := rarc.mutation.CurrentReps(); !ok {
		return &ValidationError{Name: "current_reps", err: errors.New(`ent: missing required field "RoutineActRec.current_reps"`)}
	}
	if v, ok := rarc.mutation.CurrentReps(); ok {
		if err := routineactrec.CurrentRepsValidator(v); err != nil {
			return &ValidationError{Name: "current_reps", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.current_reps": %w`, err)}
		}
	}
	if _, ok := rarc.mutation.CurrentLap(); !ok {
		return &ValidationError{Name: "current_lap", err: errors.New(`ent: missing required field "RoutineActRec.current_lap"`)}
	}
	if v, ok := rarc.mutation.CurrentLap(); ok {
		if err := routineactrec.CurrentLapValidator(v); err != nil {
			return &ValidationError{Name: "current_lap", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.current_lap": %w`, err)}
		}
	}
	if _, ok := rarc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "RoutineActRec.status"`)}
	}
	if v, ok := rarc.mutation.Status(); ok {
		if err := routineactrec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.status": %w`, err)}
		}
	}
	if _, ok := rarc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RoutineActRec.created_at"`)}
	}
	if _, ok := rarc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RoutineActRec.updated_at"`)}
	}
	if _, ok := rarc.mutation.DailyRoutineRecID(); !ok {
		return &ValidationError{Name: "daily_routine_rec", err: errors.New(`ent: missing required edge "RoutineActRec.daily_routine_rec"`)}
	}
	if _, ok := rarc.mutation.ActID(); !ok {
		return &ValidationError{Name: "act", err: errors.New(`ent: missing required edge "RoutineActRec.act"`)}
	}
	return nil
}

func (rarc *RoutineActRecCreate) sqlSave(ctx context.Context) (*RoutineActRec, error) {
	if err := rarc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rarc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rarc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	rarc.mutation.id = &_node.ID
	rarc.mutation.done = true
	return _node, nil
}

func (rarc *RoutineActRecCreate) createSpec() (*RoutineActRec, *sqlgraph.CreateSpec) {
	var (
		_node = &RoutineActRec{config: rarc.config}
		_spec = sqlgraph.NewCreateSpec(routineactrec.Table, sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64))
	)
	if id, ok := rarc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rarc.mutation.Author(); ok {
		_spec.SetField(routineactrec.FieldAuthor, field.TypeUint64, value)
		_node.Author = value
	}
	if value, ok := rarc.mutation.Order(); ok {
		_spec.SetField(routineactrec.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := rarc.mutation.Reps(); ok {
		_spec.SetField(routineactrec.FieldReps, field.TypeInt, value)
		_node.Reps = &value
	}
	if value, ok := rarc.mutation.Lap(); ok {
		_spec.SetField(routineactrec.FieldLap, field.TypeInt, value)
		_node.Lap = &value
	}
	if value, ok := rarc.mutation.CurrentReps(); ok {
		_spec.SetField(routineactrec.FieldCurrentReps, field.TypeInt, value)
		_node.CurrentReps = value
	}
	if value, ok := rarc.mutation.CurrentLap(); ok {
		_spec.SetField(routineactrec.FieldCurrentLap, field.TypeInt, value)
		_node.CurrentLap = value
	}
	if value, ok := rarc.mutation.StartedAt(); ok {
		_spec.SetField(routineactrec.FieldStartedAt, field.TypeTime, value)
		_node.StartedAt = &value
	}
	if value, ok := rarc.mutation.Image(); ok {
		_spec.SetField(routineactrec.FieldImage, field.TypeString, value)
		_node.Image = &value
	}
	if value, ok := rarc.mutation.Comment(); ok {
		_spec.SetField(routineactrec.FieldComment, field.TypeString, value)
		_node.Comment = &value
	}
	if value, ok := rarc.mutation.Status(); ok {
		_spec.SetField(routineactrec.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := rarc.mutation.CreatedAt(); ok {
		_spec.SetField(routineactrec.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rarc.mutation.UpdatedAt(); ok {
		_spec.SetField(routineactrec.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := rarc.mutation.DailyRoutineRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.DailyRoutineRecTable,
			Columns: []string{routineactrec.DailyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DailyRoutineRecID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rarc.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.ActTable,
			Columns: []string{routineactrec.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rarc.mutation.RoutineActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.RoutineActTable,
			Columns: []string{routineactrec.RoutineActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoutineActID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoutineActRecCreateBulk is the builder for creating many RoutineActRec entities in bulk.
type RoutineActRecCreateBulk struct {
	config
	builders []*RoutineActRecCreate
}

// Save creates the RoutineActRec entities in the database.
func (rarcb *RoutineActRecCreateBulk) Save(ctx context.Context) ([]*RoutineActRec, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rarcb.builders))
	nodes := make([]*RoutineActRec, len(rarcb.builders))
	mutators := make([]Mutator, len(rarcb.builders))
	for i := range rarcb.builders {
		func(i int, root context.Context) {
			builder := rarcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoutineActRecMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rarcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rarcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rarcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rarcb *RoutineActRecCreateBulk) SaveX(ctx context.Context) []*RoutineActRec {
	v, err := rarcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rarcb *RoutineActRecCreateBulk) Exec(ctx context.Context) error {
	_, err := rarcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rarcb *RoutineActRecCreateBulk) ExecX(ctx context.Context) {
	if err := rarcb.Exec(ctx); err != nil {
		panic(err)
	}
}
