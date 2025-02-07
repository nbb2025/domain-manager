// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"domain-manager/ent/domain"
	"domain-manager/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainCreate is the builder for creating a Domain entity.
type DomainCreate struct {
	config
	mutation *DomainMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DomainCreate) SetName(s string) *DomainCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetProvider sets the "provider" field.
func (dc *DomainCreate) SetProvider(s string) *DomainCreate {
	dc.mutation.SetProvider(s)
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DomainCreate) SetCreatedAt(t time.Time) *DomainCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DomainCreate) SetUpdatedAt(t time.Time) *DomainCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetID sets the "id" field.
func (dc *DomainCreate) SetID(i int) *DomainCreate {
	dc.mutation.SetID(i)
	return dc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (dc *DomainCreate) SetUserID(id int) *DomainCreate {
	dc.mutation.SetUserID(id)
	return dc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (dc *DomainCreate) SetNillableUserID(id *int) *DomainCreate {
	if id != nil {
		dc = dc.SetUserID(*id)
	}
	return dc
}

// SetUser sets the "user" edge to the User entity.
func (dc *DomainCreate) SetUser(u *User) *DomainCreate {
	return dc.SetUserID(u.ID)
}

// Mutation returns the DomainMutation object of the builder.
func (dc *DomainCreate) Mutation() *DomainMutation {
	return dc.mutation
}

// Save creates the Domain in the database.
func (dc *DomainCreate) Save(ctx context.Context) (*Domain, error) {
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DomainCreate) SaveX(ctx context.Context) *Domain {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DomainCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DomainCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DomainCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Domain.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := domain.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Domain.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "Domain.provider"`)}
	}
	if v, ok := dc.mutation.Provider(); ok {
		if err := domain.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "Domain.provider": %w`, err)}
		}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Domain.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Domain.updated_at"`)}
	}
	if v, ok := dc.mutation.ID(); ok {
		if err := domain.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Domain.id": %w`, err)}
		}
	}
	return nil
}

func (dc *DomainCreate) sqlSave(ctx context.Context) (*Domain, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DomainCreate) createSpec() (*Domain, *sqlgraph.CreateSpec) {
	var (
		_node = &Domain{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(domain.Table, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(domain.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Provider(); ok {
		_spec.SetField(domain.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(domain.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(domain.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   domain.UserTable,
			Columns: []string{domain.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_domains = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DomainCreateBulk is the builder for creating many Domain entities in bulk.
type DomainCreateBulk struct {
	config
	err      error
	builders []*DomainCreate
}

// Save creates the Domain entities in the database.
func (dcb *DomainCreateBulk) Save(ctx context.Context) ([]*Domain, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Domain, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DomainMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DomainCreateBulk) SaveX(ctx context.Context) []*Domain {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DomainCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DomainCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
