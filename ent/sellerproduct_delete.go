// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/sellerproduct"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerProductDelete is the builder for deleting a SellerProduct entity.
type SellerProductDelete struct {
	config
	hooks    []Hook
	mutation *SellerProductMutation
}

// Where appends a list predicates to the SellerProductDelete builder.
func (spd *SellerProductDelete) Where(ps ...predicate.SellerProduct) *SellerProductDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *SellerProductDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spd.hooks) == 0 {
		affected, err = spd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spd.mutation = mutation
			affected, err = spd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spd.hooks) - 1; i >= 0; i-- {
			if spd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *SellerProductDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *SellerProductDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sellerproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproduct.FieldID,
			},
		},
	}
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
}

// SellerProductDeleteOne is the builder for deleting a single SellerProduct entity.
type SellerProductDeleteOne struct {
	spd *SellerProductDelete
}

// Exec executes the deletion query.
func (spdo *SellerProductDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sellerproduct.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *SellerProductDeleteOne) ExecX(ctx context.Context) {
	spdo.spd.ExecX(ctx)
}