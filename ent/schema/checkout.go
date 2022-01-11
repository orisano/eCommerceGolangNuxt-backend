package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"time"
)

// Checkout holds the schema definition for the Checkout entity.
type Checkout struct {
	ent.Schema
}

// Fields of the Checkout.
func (Checkout) Fields() []ent.Field {
	return []ent.Field{
		field.Float("total_price").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Bool("completed").Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}
// Edges of the Checkout.
func (Checkout) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("user",User.Type).Ref("checkouts").Unique(),
		edge.From("location",UserLocation.Type).Ref("user_locations").Unique(),
		edge.From("cart",Cart.Type).Ref("checkout").Unique(),
		edge.To("checkout_products",CheckoutProduct.Type),
	}
}
