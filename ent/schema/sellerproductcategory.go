package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// SellerProductCategory holds the schema definition for the SellerProductCategory entity.
type SellerProductCategory struct {
	ent.Schema
}

// Fields of the SellerProductCategory.
func (SellerProductCategory) Fields() []ent.Field {
	return []ent.Field {
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerProductCategory.
func (SellerProductCategory) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("seller_product",SellerProduct.Type).Ref("seller_product_categories").Unique(),
		edge.From("category",Category.Type).Ref("product_categories").Unique(),
	}
}
