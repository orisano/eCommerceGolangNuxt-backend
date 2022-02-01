// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/sellerrequest"
	"bongo/ent/shopcategory"
	"bongo/ent/user"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerRequestUpdate is the builder for updating SellerRequest entities.
type SellerRequestUpdate struct {
	config
	hooks    []Hook
	mutation *SellerRequestMutation
}

// Where appends a list predicates to the SellerRequestUpdate builder.
func (sru *SellerRequestUpdate) Where(ps ...predicate.SellerRequest) *SellerRequestUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetSellerName sets the "seller_name" field.
func (sru *SellerRequestUpdate) SetSellerName(s string) *SellerRequestUpdate {
	sru.mutation.SetSellerName(s)
	return sru
}

// SetShopName sets the "shop_name" field.
func (sru *SellerRequestUpdate) SetShopName(s string) *SellerRequestUpdate {
	sru.mutation.SetShopName(s)
	return sru
}

// SetContactNumber sets the "contact_number" field.
func (sru *SellerRequestUpdate) SetContactNumber(s string) *SellerRequestUpdate {
	sru.mutation.SetContactNumber(s)
	return sru
}

// SetShopLocation sets the "shop_location" field.
func (sru *SellerRequestUpdate) SetShopLocation(s string) *SellerRequestUpdate {
	sru.mutation.SetShopLocation(s)
	return sru
}

// SetTaxID sets the "tax_id" field.
func (sru *SellerRequestUpdate) SetTaxID(s string) *SellerRequestUpdate {
	sru.mutation.SetTaxID(s)
	return sru
}

// SetAccepted sets the "accepted" field.
func (sru *SellerRequestUpdate) SetAccepted(b bool) *SellerRequestUpdate {
	sru.mutation.SetAccepted(b)
	return sru
}

// SetNillableAccepted sets the "accepted" field if the given value is not nil.
func (sru *SellerRequestUpdate) SetNillableAccepted(b *bool) *SellerRequestUpdate {
	if b != nil {
		sru.SetAccepted(*b)
	}
	return sru
}

// SetUpdatedAt sets the "updated_at" field.
func (sru *SellerRequestUpdate) SetUpdatedAt(t time.Time) *SellerRequestUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// SetDeletedAt sets the "deleted_at" field.
func (sru *SellerRequestUpdate) SetDeletedAt(t time.Time) *SellerRequestUpdate {
	sru.mutation.SetDeletedAt(t)
	return sru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sru *SellerRequestUpdate) SetNillableDeletedAt(t *time.Time) *SellerRequestUpdate {
	if t != nil {
		sru.SetDeletedAt(*t)
	}
	return sru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sru *SellerRequestUpdate) ClearDeletedAt() *SellerRequestUpdate {
	sru.mutation.ClearDeletedAt()
	return sru
}

// SetShopCategoryID sets the "shop_category" edge to the ShopCategory entity by ID.
func (sru *SellerRequestUpdate) SetShopCategoryID(id int) *SellerRequestUpdate {
	sru.mutation.SetShopCategoryID(id)
	return sru
}

// SetNillableShopCategoryID sets the "shop_category" edge to the ShopCategory entity by ID if the given value is not nil.
func (sru *SellerRequestUpdate) SetNillableShopCategoryID(id *int) *SellerRequestUpdate {
	if id != nil {
		sru = sru.SetShopCategoryID(*id)
	}
	return sru
}

// SetShopCategory sets the "shop_category" edge to the ShopCategory entity.
func (sru *SellerRequestUpdate) SetShopCategory(s *ShopCategory) *SellerRequestUpdate {
	return sru.SetShopCategoryID(s.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sru *SellerRequestUpdate) SetUserID(id int) *SellerRequestUpdate {
	sru.mutation.SetUserID(id)
	return sru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (sru *SellerRequestUpdate) SetNillableUserID(id *int) *SellerRequestUpdate {
	if id != nil {
		sru = sru.SetUserID(*id)
	}
	return sru
}

// SetUser sets the "user" edge to the User entity.
func (sru *SellerRequestUpdate) SetUser(u *User) *SellerRequestUpdate {
	return sru.SetUserID(u.ID)
}

// Mutation returns the SellerRequestMutation object of the builder.
func (sru *SellerRequestUpdate) Mutation() *SellerRequestMutation {
	return sru.mutation
}

// ClearShopCategory clears the "shop_category" edge to the ShopCategory entity.
func (sru *SellerRequestUpdate) ClearShopCategory() *SellerRequestUpdate {
	sru.mutation.ClearShopCategory()
	return sru
}

// ClearUser clears the "user" edge to the User entity.
func (sru *SellerRequestUpdate) ClearUser() *SellerRequestUpdate {
	sru.mutation.ClearUser()
	return sru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *SellerRequestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sru.defaults()
	if len(sru.hooks) == 0 {
		if err = sru.check(); err != nil {
			return 0, err
		}
		affected, err = sru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sru.check(); err != nil {
				return 0, err
			}
			sru.mutation = mutation
			affected, err = sru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sru.hooks) - 1; i >= 0; i-- {
			if sru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sru *SellerRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *SellerRequestUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *SellerRequestUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *SellerRequestUpdate) defaults() {
	if _, ok := sru.mutation.UpdatedAt(); !ok {
		v := sellerrequest.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *SellerRequestUpdate) check() error {
	if v, ok := sru.mutation.ContactNumber(); ok {
		if err := sellerrequest.ContactNumberValidator(v); err != nil {
			return &ValidationError{Name: "contact_number", err: fmt.Errorf("ent: validator failed for field \"contact_number\": %w", err)}
		}
	}
	return nil
}

func (sru *SellerRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sellerrequest.Table,
			Columns: sellerrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerrequest.FieldID,
			},
		},
	}
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.SellerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldSellerName,
		})
	}
	if value, ok := sru.mutation.ShopName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldShopName,
		})
	}
	if value, ok := sru.mutation.ContactNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldContactNumber,
		})
	}
	if value, ok := sru.mutation.ShopLocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldShopLocation,
		})
	}
	if value, ok := sru.mutation.TaxID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldTaxID,
		})
	}
	if value, ok := sru.mutation.Accepted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sellerrequest.FieldAccepted,
		})
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerrequest.FieldUpdatedAt,
		})
	}
	if value, ok := sru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerrequest.FieldDeletedAt,
		})
	}
	if sru.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sellerrequest.FieldDeletedAt,
		})
	}
	if sru.mutation.ShopCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.ShopCategoryTable,
			Columns: []string{sellerrequest.ShopCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.ShopCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.ShopCategoryTable,
			Columns: []string{sellerrequest.ShopCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.UserTable,
			Columns: []string{sellerrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.UserTable,
			Columns: []string{sellerrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sellerrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SellerRequestUpdateOne is the builder for updating a single SellerRequest entity.
type SellerRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SellerRequestMutation
}

// SetSellerName sets the "seller_name" field.
func (sruo *SellerRequestUpdateOne) SetSellerName(s string) *SellerRequestUpdateOne {
	sruo.mutation.SetSellerName(s)
	return sruo
}

// SetShopName sets the "shop_name" field.
func (sruo *SellerRequestUpdateOne) SetShopName(s string) *SellerRequestUpdateOne {
	sruo.mutation.SetShopName(s)
	return sruo
}

// SetContactNumber sets the "contact_number" field.
func (sruo *SellerRequestUpdateOne) SetContactNumber(s string) *SellerRequestUpdateOne {
	sruo.mutation.SetContactNumber(s)
	return sruo
}

// SetShopLocation sets the "shop_location" field.
func (sruo *SellerRequestUpdateOne) SetShopLocation(s string) *SellerRequestUpdateOne {
	sruo.mutation.SetShopLocation(s)
	return sruo
}

// SetTaxID sets the "tax_id" field.
func (sruo *SellerRequestUpdateOne) SetTaxID(s string) *SellerRequestUpdateOne {
	sruo.mutation.SetTaxID(s)
	return sruo
}

// SetAccepted sets the "accepted" field.
func (sruo *SellerRequestUpdateOne) SetAccepted(b bool) *SellerRequestUpdateOne {
	sruo.mutation.SetAccepted(b)
	return sruo
}

// SetNillableAccepted sets the "accepted" field if the given value is not nil.
func (sruo *SellerRequestUpdateOne) SetNillableAccepted(b *bool) *SellerRequestUpdateOne {
	if b != nil {
		sruo.SetAccepted(*b)
	}
	return sruo
}

// SetUpdatedAt sets the "updated_at" field.
func (sruo *SellerRequestUpdateOne) SetUpdatedAt(t time.Time) *SellerRequestUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// SetDeletedAt sets the "deleted_at" field.
func (sruo *SellerRequestUpdateOne) SetDeletedAt(t time.Time) *SellerRequestUpdateOne {
	sruo.mutation.SetDeletedAt(t)
	return sruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sruo *SellerRequestUpdateOne) SetNillableDeletedAt(t *time.Time) *SellerRequestUpdateOne {
	if t != nil {
		sruo.SetDeletedAt(*t)
	}
	return sruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sruo *SellerRequestUpdateOne) ClearDeletedAt() *SellerRequestUpdateOne {
	sruo.mutation.ClearDeletedAt()
	return sruo
}

// SetShopCategoryID sets the "shop_category" edge to the ShopCategory entity by ID.
func (sruo *SellerRequestUpdateOne) SetShopCategoryID(id int) *SellerRequestUpdateOne {
	sruo.mutation.SetShopCategoryID(id)
	return sruo
}

// SetNillableShopCategoryID sets the "shop_category" edge to the ShopCategory entity by ID if the given value is not nil.
func (sruo *SellerRequestUpdateOne) SetNillableShopCategoryID(id *int) *SellerRequestUpdateOne {
	if id != nil {
		sruo = sruo.SetShopCategoryID(*id)
	}
	return sruo
}

// SetShopCategory sets the "shop_category" edge to the ShopCategory entity.
func (sruo *SellerRequestUpdateOne) SetShopCategory(s *ShopCategory) *SellerRequestUpdateOne {
	return sruo.SetShopCategoryID(s.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sruo *SellerRequestUpdateOne) SetUserID(id int) *SellerRequestUpdateOne {
	sruo.mutation.SetUserID(id)
	return sruo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (sruo *SellerRequestUpdateOne) SetNillableUserID(id *int) *SellerRequestUpdateOne {
	if id != nil {
		sruo = sruo.SetUserID(*id)
	}
	return sruo
}

// SetUser sets the "user" edge to the User entity.
func (sruo *SellerRequestUpdateOne) SetUser(u *User) *SellerRequestUpdateOne {
	return sruo.SetUserID(u.ID)
}

// Mutation returns the SellerRequestMutation object of the builder.
func (sruo *SellerRequestUpdateOne) Mutation() *SellerRequestMutation {
	return sruo.mutation
}

// ClearShopCategory clears the "shop_category" edge to the ShopCategory entity.
func (sruo *SellerRequestUpdateOne) ClearShopCategory() *SellerRequestUpdateOne {
	sruo.mutation.ClearShopCategory()
	return sruo
}

// ClearUser clears the "user" edge to the User entity.
func (sruo *SellerRequestUpdateOne) ClearUser() *SellerRequestUpdateOne {
	sruo.mutation.ClearUser()
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *SellerRequestUpdateOne) Select(field string, fields ...string) *SellerRequestUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated SellerRequest entity.
func (sruo *SellerRequestUpdateOne) Save(ctx context.Context) (*SellerRequest, error) {
	var (
		err  error
		node *SellerRequest
	)
	sruo.defaults()
	if len(sruo.hooks) == 0 {
		if err = sruo.check(); err != nil {
			return nil, err
		}
		node, err = sruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sruo.check(); err != nil {
				return nil, err
			}
			sruo.mutation = mutation
			node, err = sruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sruo.hooks) - 1; i >= 0; i-- {
			if sruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *SellerRequestUpdateOne) SaveX(ctx context.Context) *SellerRequest {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *SellerRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *SellerRequestUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *SellerRequestUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdatedAt(); !ok {
		v := sellerrequest.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *SellerRequestUpdateOne) check() error {
	if v, ok := sruo.mutation.ContactNumber(); ok {
		if err := sellerrequest.ContactNumberValidator(v); err != nil {
			return &ValidationError{Name: "contact_number", err: fmt.Errorf("ent: validator failed for field \"contact_number\": %w", err)}
		}
	}
	return nil
}

func (sruo *SellerRequestUpdateOne) sqlSave(ctx context.Context) (_node *SellerRequest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sellerrequest.Table,
			Columns: sellerrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerrequest.FieldID,
			},
		},
	}
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SellerRequest.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sellerrequest.FieldID)
		for _, f := range fields {
			if !sellerrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sellerrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.SellerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldSellerName,
		})
	}
	if value, ok := sruo.mutation.ShopName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldShopName,
		})
	}
	if value, ok := sruo.mutation.ContactNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldContactNumber,
		})
	}
	if value, ok := sruo.mutation.ShopLocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldShopLocation,
		})
	}
	if value, ok := sruo.mutation.TaxID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerrequest.FieldTaxID,
		})
	}
	if value, ok := sruo.mutation.Accepted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sellerrequest.FieldAccepted,
		})
	}
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerrequest.FieldUpdatedAt,
		})
	}
	if value, ok := sruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerrequest.FieldDeletedAt,
		})
	}
	if sruo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sellerrequest.FieldDeletedAt,
		})
	}
	if sruo.mutation.ShopCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.ShopCategoryTable,
			Columns: []string{sellerrequest.ShopCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.ShopCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.ShopCategoryTable,
			Columns: []string{sellerrequest.ShopCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.UserTable,
			Columns: []string{sellerrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerrequest.UserTable,
			Columns: []string{sellerrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SellerRequest{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sellerrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}