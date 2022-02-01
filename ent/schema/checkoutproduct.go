package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"time"
)

// CheckoutProduct holds the schema definition for the CheckoutProduct entity.
type CheckoutProduct struct {
	ent.Schema
}

// Fields of the CheckoutProduct.
func (CheckoutProduct) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
		field.Float("selling_price").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Int("offer_price").Default(0),
		field.Bool("received").Default(false),
		field.Int("status").Default(0),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the CheckoutProduct.
func (CheckoutProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("checkout_products").Unique(),

		edge.From("checkout", Checkout.Type).Ref("checkout_products").Unique(),
		edge.From("seller", User.Type).Ref("seller_checkout_products").Unique(),

		edge.From("seller_product", SellerProduct.Type).Ref("checkout_products").Unique(),
		edge.From("seller_product_variation", SellerProductVariation.Type).Ref("checkout_products").Unique(),
	}
}
