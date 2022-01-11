package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// SellerProductImage holds the schema definition for the SellerProductImage entity.
type SellerProductImage struct {
	ent.Schema
}

// Fields of the SellerProductImage.
func (SellerProductImage) Fields() []ent.Field {
	return []ent.Field {
		field.Bool("display").Default(false),
		field.String("image"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerProductImage.
func (SellerProductImage) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("seller_product",SellerProduct.Type).Ref("seller_product_images").Unique(),
	}
}
