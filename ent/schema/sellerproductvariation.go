package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"time"
)

// SellerProductVariation holds the schema definition for the SellerProductVariation entity.
type SellerProductVariation struct {
	ent.Schema
}

// Fields of the SellerProductVariation.
func (SellerProductVariation) Fields() []ent.Field {
	return []ent.Field {
		field.Float("product_price").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Float("selling_price").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Int("quantity"),
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

// Edges of the SellerProductVariation.
func (SellerProductVariation) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("seller_product",SellerProduct.Type).Ref("seller_product_variations").Unique(),
		edge.To("seller_product_variation_values",SellerProductVariationValues.Type),
		edge.To("cart_products",CartProduct.Type),
		edge.To("checkout_products",CheckoutProduct.Type),
	}
}
