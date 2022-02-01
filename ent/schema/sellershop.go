package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// SellerShop holds the schema definition for the SellerShop entity.
type SellerShop struct {
	ent.Schema
}

// Fields of the SellerShop.
func (SellerShop) Fields() []ent.Field {
	return []ent.Field {
		field.String("name"),
		field.String("slug"),
		field.String("contact_number").Match(regexp.MustCompile("(^(01)[3-9]\\d{8})$")),
		field.String("banner"),
		field.String("business_location"),
		field.String("tax_id"),
		field.Bool("active").Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerShop.
func (SellerShop) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("user",User.Type).Ref("seller_shops").Unique(),
		edge.From("admin",User.Type).Ref("approved_shops").Unique(),
		edge.From("get_shop_category",ShopCategory.Type).Ref("seller_shops").Unique(),
		edge.To("seller_products",SellerProduct.Type),
	}
}
