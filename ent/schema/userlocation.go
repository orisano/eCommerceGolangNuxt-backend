package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// UserLocation holds the schema definition for the UserLocation entity.
type UserLocation struct {
	ent.Schema
}

// Fields of the UserLocation.
func (UserLocation) Fields() []ent.Field {
	return []ent.Field {
		field.String("area"),
		field.String("street"),
		field.String("house"),
		field.String("post_office"),
		field.Int("post_code"),
		field.String("police_station"),
		field.String("city"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the UserLocation.
func (UserLocation) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("get_user",User.Type).Ref("user_locations").Unique(),
		edge.To("user_locations",Checkout.Type),
	}
}
