// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"domain-manager/ent/predicate"
	"domain-manager/ent/providerkey"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProviderKeyDelete is the builder for deleting a ProviderKey entity.
type ProviderKeyDelete struct {
	config
	hooks    []Hook
	mutation *ProviderKeyMutation
}

// Where appends a list predicates to the ProviderKeyDelete builder.
func (pkd *ProviderKeyDelete) Where(ps ...predicate.ProviderKey) *ProviderKeyDelete {
	pkd.mutation.Where(ps...)
	return pkd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pkd *ProviderKeyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pkd.sqlExec, pkd.mutation, pkd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pkd *ProviderKeyDelete) ExecX(ctx context.Context) int {
	n, err := pkd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pkd *ProviderKeyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(providerkey.Table, sqlgraph.NewFieldSpec(providerkey.FieldID, field.TypeInt))
	if ps := pkd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pkd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pkd.mutation.done = true
	return affected, err
}

// ProviderKeyDeleteOne is the builder for deleting a single ProviderKey entity.
type ProviderKeyDeleteOne struct {
	pkd *ProviderKeyDelete
}

// Where appends a list predicates to the ProviderKeyDelete builder.
func (pkdo *ProviderKeyDeleteOne) Where(ps ...predicate.ProviderKey) *ProviderKeyDeleteOne {
	pkdo.pkd.mutation.Where(ps...)
	return pkdo
}

// Exec executes the deletion query.
func (pkdo *ProviderKeyDeleteOne) Exec(ctx context.Context) error {
	n, err := pkdo.pkd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{providerkey.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pkdo *ProviderKeyDeleteOne) ExecX(ctx context.Context) {
	if err := pkdo.Exec(ctx); err != nil {
		panic(err)
	}
}
