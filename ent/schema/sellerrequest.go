package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// SellerRequest holds the schema definition for the SellerRequest entity.
type SellerRequest struct {
	ent.Schema
}

// Fields of the SellerRequest.
func (SellerRequest) Fields() []ent.Field {
	return []ent.Field {
		field.String("seller_name"),
		field.String("shop_name"),
		field.String("contact_number").Unique().Match(regexp.MustCompile("(^(01)[3-9]\\d{8})$")),
		field.String("shop_location"),
		field.String("tax_id"),
		field.Bool("accepted").Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerRequest.
func (SellerRequest) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("shop_category",ShopCategory.Type).Ref("seller_requests").Unique(),
		edge.From("user",User.Type).Ref("seller_requests").Unique(),
	}
}
