package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// CartProduct holds the schema definition for the CartProduct entity.
type CartProduct struct {
	ent.Schema
}

// Fields of the CartProduct.
func (CartProduct) Fields() []ent.Field {
	return []ent.Field {
		field.Int("quantity"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the CartProduct.
func (CartProduct) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("seller_product",SellerProduct.Type).Ref("cart_products").Unique(),
		edge.From("cart",Cart.Type).Ref("cart_products").Unique(),
		edge.From("seller_product_variation",SellerProductVariation.Type).Ref("cart_products").Unique(),
	}
}
