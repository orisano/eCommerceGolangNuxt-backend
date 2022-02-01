// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/attribute"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Attribute is the model entity for the Attribute schema.
type Attribute struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttributeQuery when eager-loading is set.
	Edges AttributeEdges `json:"edges"`
}

// AttributeEdges holds the relations/edges for other nodes in the graph.
type AttributeEdges struct {
	// Attribute holds the value of the attribute edge.
	Attribute []*SellerProductVariationValues `json:"attribute,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AttributeOrErr returns the Attribute value or an error if the edge
// was not loaded in eager-loading.
func (e AttributeEdges) AttributeOrErr() ([]*SellerProductVariationValues, error) {
	if e.loadedTypes[0] {
		return e.Attribute, nil
	}
	return nil, &NotLoadedError{edge: "attribute"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attribute) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case attribute.FieldID:
			values[i] = new(sql.NullInt64)
		case attribute.FieldName:
			values[i] = new(sql.NullString)
		case attribute.FieldCreatedAt, attribute.FieldUpdatedAt, attribute.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Attribute", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attribute fields.
func (a *Attribute) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attribute.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case attribute.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case attribute.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case attribute.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case attribute.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryAttribute queries the "attribute" edge of the Attribute entity.
func (a *Attribute) QueryAttribute() *SellerProductVariationValuesQuery {
	return (&AttributeClient{config: a.config}).QueryAttribute(a)
}

// Update returns a builder for updating this Attribute.
// Note that you need to call Attribute.Unwrap() before calling this method if this Attribute
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attribute) Update() *AttributeUpdateOne {
	return (&AttributeClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Attribute entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attribute) Unwrap() *Attribute {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attribute is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attribute) String() string {
	var builder strings.Builder
	builder.WriteString("Attribute(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	if v := a.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Attributes is a parsable slice of Attribute.
type Attributes []*Attribute

func (a Attributes) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
