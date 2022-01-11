package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// SellerProductVariationValues holds the schema definition for the SellerProductVariationValues entity.
type SellerProductVariationValues struct {
	ent.Schema
}

// Fields of the SellerProductVariationValues.
func (SellerProductVariationValues) Fields() []ent.Field {
	return []ent.Field {
		field.String("name"),
		field.String("description"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the SellerProductVariationValues.
func (SellerProductVariationValues) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("attribute",Attribute.Type).Ref("attribute").Unique(),
		edge.From("seller_product_variation",SellerProductVariation.Type).Ref("seller_product_variation_values").Unique(),
	}
}
