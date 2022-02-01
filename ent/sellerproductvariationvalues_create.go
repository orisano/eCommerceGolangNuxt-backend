// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/attribute"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/sellerproductvariationvalues"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerProductVariationValuesCreate is the builder for creating a SellerProductVariationValues entity.
type SellerProductVariationValuesCreate struct {
	config
	mutation *SellerProductVariationValuesMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (spvvc *SellerProductVariationValuesCreate) SetName(s string) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetName(s)
	return spvvc
}

// SetDescription sets the "description" field.
func (spvvc *SellerProductVariationValuesCreate) SetDescription(s string) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetDescription(s)
	return spvvc
}

// SetCreatedAt sets the "created_at" field.
func (spvvc *SellerProductVariationValuesCreate) SetCreatedAt(t time.Time) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetCreatedAt(t)
	return spvvc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spvvc *SellerProductVariationValuesCreate) SetNillableCreatedAt(t *time.Time) *SellerProductVariationValuesCreate {
	if t != nil {
		spvvc.SetCreatedAt(*t)
	}
	return spvvc
}

// SetUpdatedAt sets the "updated_at" field.
func (spvvc *SellerProductVariationValuesCreate) SetUpdatedAt(t time.Time) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetUpdatedAt(t)
	return spvvc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (spvvc *SellerProductVariationValuesCreate) SetNillableUpdatedAt(t *time.Time) *SellerProductVariationValuesCreate {
	if t != nil {
		spvvc.SetUpdatedAt(*t)
	}
	return spvvc
}

// SetDeletedAt sets the "deleted_at" field.
func (spvvc *SellerProductVariationValuesCreate) SetDeletedAt(t time.Time) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetDeletedAt(t)
	return spvvc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spvvc *SellerProductVariationValuesCreate) SetNillableDeletedAt(t *time.Time) *SellerProductVariationValuesCreate {
	if t != nil {
		spvvc.SetDeletedAt(*t)
	}
	return spvvc
}

// SetAttributeID sets the "attribute" edge to the Attribute entity by ID.
func (spvvc *SellerProductVariationValuesCreate) SetAttributeID(id int) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetAttributeID(id)
	return spvvc
}

// SetNillableAttributeID sets the "attribute" edge to the Attribute entity by ID if the given value is not nil.
func (spvvc *SellerProductVariationValuesCreate) SetNillableAttributeID(id *int) *SellerProductVariationValuesCreate {
	if id != nil {
		spvvc = spvvc.SetAttributeID(*id)
	}
	return spvvc
}

// SetAttribute sets the "attribute" edge to the Attribute entity.
func (spvvc *SellerProductVariationValuesCreate) SetAttribute(a *Attribute) *SellerProductVariationValuesCreate {
	return spvvc.SetAttributeID(a.ID)
}

// SetSellerProductVariationID sets the "seller_product_variation" edge to the SellerProductVariation entity by ID.
func (spvvc *SellerProductVariationValuesCreate) SetSellerProductVariationID(id int) *SellerProductVariationValuesCreate {
	spvvc.mutation.SetSellerProductVariationID(id)
	return spvvc
}

// SetNillableSellerProductVariationID sets the "seller_product_variation" edge to the SellerProductVariation entity by ID if the given value is not nil.
func (spvvc *SellerProductVariationValuesCreate) SetNillableSellerProductVariationID(id *int) *SellerProductVariationValuesCreate {
	if id != nil {
		spvvc = spvvc.SetSellerProductVariationID(*id)
	}
	return spvvc
}

// SetSellerProductVariation sets the "seller_product_variation" edge to the SellerProductVariation entity.
func (spvvc *SellerProductVariationValuesCreate) SetSellerProductVariation(s *SellerProductVariation) *SellerProductVariationValuesCreate {
	return spvvc.SetSellerProductVariationID(s.ID)
}

// Mutation returns the SellerProductVariationValuesMutation object of the builder.
func (spvvc *SellerProductVariationValuesCreate) Mutation() *SellerProductVariationValuesMutation {
	return spvvc.mutation
}

// Save creates the SellerProductVariationValues in the database.
func (spvvc *SellerProductVariationValuesCreate) Save(ctx context.Context) (*SellerProductVariationValues, error) {
	var (
		err  error
		node *SellerProductVariationValues
	)
	spvvc.defaults()
	if len(spvvc.hooks) == 0 {
		if err = spvvc.check(); err != nil {
			return nil, err
		}
		node, err = spvvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductVariationValuesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spvvc.check(); err != nil {
				return nil, err
			}
			spvvc.mutation = mutation
			if node, err = spvvc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spvvc.hooks) - 1; i >= 0; i-- {
			if spvvc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spvvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spvvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spvvc *SellerProductVariationValuesCreate) SaveX(ctx context.Context) *SellerProductVariationValues {
	v, err := spvvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spvvc *SellerProductVariationValuesCreate) Exec(ctx context.Context) error {
	_, err := spvvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvvc *SellerProductVariationValuesCreate) ExecX(ctx context.Context) {
	if err := spvvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spvvc *SellerProductVariationValuesCreate) defaults() {
	if _, ok := spvvc.mutation.CreatedAt(); !ok {
		v := sellerproductvariationvalues.DefaultCreatedAt()
		spvvc.mutation.SetCreatedAt(v)
	}
	if _, ok := spvvc.mutation.UpdatedAt(); !ok {
		v := sellerproductvariationvalues.DefaultUpdatedAt()
		spvvc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spvvc *SellerProductVariationValuesCreate) check() error {
	if _, ok := spvvc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := spvvc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := spvvc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := spvvc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (spvvc *SellerProductVariationValuesCreate) sqlSave(ctx context.Context) (*SellerProductVariationValues, error) {
	_node, _spec := spvvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spvvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (spvvc *SellerProductVariationValuesCreate) createSpec() (*SellerProductVariationValues, *sqlgraph.CreateSpec) {
	var (
		_node = &SellerProductVariationValues{config: spvvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sellerproductvariationvalues.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductvariationvalues.FieldID,
			},
		}
	)
	if value, ok := spvvc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerproductvariationvalues.FieldName,
		})
		_node.Name = value
	}
	if value, ok := spvvc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerproductvariationvalues.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := spvvc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := spvvc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := spvvc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := spvvc.mutation.AttributeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.AttributeTable,
			Columns: []string{sellerproductvariationvalues.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.attribute_attribute = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spvvc.mutation.SellerProductVariationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.SellerProductVariationTable,
			Columns: []string{sellerproductvariationvalues.SellerProductVariationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproductvariation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.seller_product_variation_seller_product_variation_values = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SellerProductVariationValuesCreateBulk is the builder for creating many SellerProductVariationValues entities in bulk.
type SellerProductVariationValuesCreateBulk struct {
	config
	builders []*SellerProductVariationValuesCreate
}

// Save creates the SellerProductVariationValues entities in the database.
func (spvvcb *SellerProductVariationValuesCreateBulk) Save(ctx context.Context) ([]*SellerProductVariationValues, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spvvcb.builders))
	nodes := make([]*SellerProductVariationValues, len(spvvcb.builders))
	mutators := make([]Mutator, len(spvvcb.builders))
	for i := range spvvcb.builders {
		func(i int, root context.Context) {
			builder := spvvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SellerProductVariationValuesMutation)
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
					_, err = mutators[i+1].Mutate(root, spvvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spvvcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, spvvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spvvcb *SellerProductVariationValuesCreateBulk) SaveX(ctx context.Context) []*SellerProductVariationValues {
	v, err := spvvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spvvcb *SellerProductVariationValuesCreateBulk) Exec(ctx context.Context) error {
	_, err := spvvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvvcb *SellerProductVariationValuesCreateBulk) ExecX(ctx context.Context) {
	if err := spvvcb.Exec(ctx); err != nil {
		panic(err)
	}
}