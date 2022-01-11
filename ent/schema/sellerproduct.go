package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"time"
)

// SellerProduct holds the schema definition for the SellerProduct entity.
type SellerProduct struct {
	ent.Schema
}

// Fields of the SellerProduct.
func (SellerProduct) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("slug").Unique(),
		field.Float("selling_price").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Float("product_price").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Int("quantity"),
		field.Bool("active").Default(true),
		field.String("description").Nillable().Optional(),
		field.Int("offer_price").Nillable().Optional(),
		field.Time("offer_price_start").Nillable().Optional(),
		field.Time("offer_price_end").Nillable().Optional(),
		field.String("next_stock").Nillable().Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerProduct.
func (SellerProduct) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("brand",Brand.Type).Ref("brand").Unique(),
		edge.From("user",User.Type).Ref("seller_products").Unique(),
		edge.To("seller_product_images",SellerProductImage.Type),
		edge.To("seller_product_categories",SellerProductCategory.Type),
		edge.To("cart_products",CartProduct.Type),
		edge.To("checkout_products",CheckoutProduct.Type),
		edge.To("seller_product_variations",SellerProductVariation.Type),
		edge.To("seller_shop_products",SellerShopProduct.Type),
	}
}
