package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// SellerShopProduct holds the schema definition for the SellerShopProduct entity.
type SellerShopProduct struct {
	ent.Schema
}

// Fields of the SellerShopProduct.
func (SellerShopProduct) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerShopProduct.
func (SellerShopProduct) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("seller_shop",SellerShop.Type).Ref("seller_shop_products").Unique(),
		edge.From("seller_product",SellerProduct.Type).Ref("seller_shop_products").Unique(),
	}
}
