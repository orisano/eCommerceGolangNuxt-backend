package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// ShopCategory holds the schema definition for the ShopCategory entity.
type ShopCategory struct {
	ent.Schema
}

// Fields of the ShopCategory.
func (ShopCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("slug"),
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

// Edges of the ShopCategory.
func (ShopCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", Category.Type),
		edge.To("seller_requests", SellerRequest.Type),
		edge.To("seller_shops", SellerShop.Type),
	}
}
