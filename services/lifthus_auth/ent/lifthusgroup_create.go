// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"lifthus-auth/ent/lifthusgroup"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LifthusGroupCreate is the builder for creating a LifthusGroup entity.
type LifthusGroupCreate struct {
	config
	mutation *LifthusGroupMutation
	hooks    []Hook
}

// Mutation returns the LifthusGroupMutation object of the builder.
func (lgc *LifthusGroupCreate) Mutation() *LifthusGroupMutation {
	return lgc.mutation
}

// Save creates the LifthusGroup in the database.
func (lgc *LifthusGroupCreate) Save(ctx context.Context) (*LifthusGroup, error) {
	return withHooks[*LifthusGroup, LifthusGroupMutation](ctx, lgc.sqlSave, lgc.mutation, lgc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lgc *LifthusGroupCreate) SaveX(ctx context.Context) *LifthusGroup {
	v, err := lgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lgc *LifthusGroupCreate) Exec(ctx context.Context) error {
	_, err := lgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lgc *LifthusGroupCreate) ExecX(ctx context.Context) {
	if err := lgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lgc *LifthusGroupCreate) check() error {
	return nil
}

func (lgc *LifthusGroupCreate) sqlSave(ctx context.Context) (*LifthusGroup, error) {
	if err := lgc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lgc.mutation.id = &_node.ID
	lgc.mutation.done = true
	return _node, nil
}

func (lgc *LifthusGroupCreate) createSpec() (*LifthusGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &LifthusGroup{config: lgc.config}
		_spec = sqlgraph.NewCreateSpec(lifthusgroup.Table, sqlgraph.NewFieldSpec(lifthusgroup.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// LifthusGroupCreateBulk is the builder for creating many LifthusGroup entities in bulk.
type LifthusGroupCreateBulk struct {
	config
	builders []*LifthusGroupCreate
}

// Save creates the LifthusGroup entities in the database.
func (lgcb *LifthusGroupCreateBulk) Save(ctx context.Context) ([]*LifthusGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lgcb.builders))
	nodes := make([]*LifthusGroup, len(lgcb.builders))
	mutators := make([]Mutator, len(lgcb.builders))
	for i := range lgcb.builders {
		func(i int, root context.Context) {
			builder := lgcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LifthusGroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lgcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, lgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lgcb *LifthusGroupCreateBulk) SaveX(ctx context.Context) []*LifthusGroup {
	v, err := lgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lgcb *LifthusGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := lgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lgcb *LifthusGroupCreateBulk) ExecX(ctx context.Context) {
	if err := lgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
