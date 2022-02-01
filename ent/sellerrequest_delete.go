// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/sellerrequest"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerRequestDelete is the builder for deleting a SellerRequest entity.
type SellerRequestDelete struct {
	config
	hooks    []Hook
	mutation *SellerRequestMutation
}

// Where appends a list predicates to the SellerRequestDelete builder.
func (srd *SellerRequestDelete) Where(ps ...predicate.SellerRequest) *SellerRequestDelete {
	srd.mutation.Where(ps...)
	return srd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srd *SellerRequestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(srd.hooks) == 0 {
		affected, err = srd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			srd.mutation = mutation
			affected, err = srd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(srd.hooks) - 1; i >= 0; i-- {
			if srd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = srd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, srd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (srd *SellerRequestDelete) ExecX(ctx context.Context) int {
	n, err := srd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srd *SellerRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sellerrequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerrequest.FieldID,
			},
		},
	}
	if ps := srd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, srd.driver, _spec)
}

// SellerRequestDeleteOne is the builder for deleting a single SellerRequest entity.
type SellerRequestDeleteOne struct {
	srd *SellerRequestDelete
}

// Exec executes the deletion query.
func (srdo *SellerRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := srdo.srd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sellerrequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srdo *SellerRequestDeleteOne) ExecX(ctx context.Context) {
	srdo.srd.ExecX(ctx)
}