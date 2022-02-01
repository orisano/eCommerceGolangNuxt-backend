package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("phone_number").Match(regexp.MustCompile("(^(01)[3-9]\\d{8})$")).Unique(),
		field.String("password").Sensitive(),
		field.Bool("admin").Default(false),
		field.Bool("staff").Default(false),
		field.Bool("seller").Default(false),
		field.Bool("active").Default(false),
		field.String("admin_user_name").Sensitive().Nillable().Optional(),
		field.String("admin_user_token").Sensitive().Nillable().Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("seller_requests", SellerRequest.Type),
		edge.To("seller_shops", SellerShop.Type),
		edge.To("seller_products", SellerProduct.Type),
		edge.To("carts", Cart.Type),
		edge.To("user_locations", UserLocation.Type),
		edge.To("checkouts", Checkout.Type),
		edge.To("checkout_products", CheckoutProduct.Type),

		edge.To("seller_checkout_products", CheckoutProduct.Type),
		edge.To("approved_shops", SellerShop.Type),
	}
}
