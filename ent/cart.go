// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/cart"
	"bongo/ent/checkout"
	"bongo/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Cart is the model entity for the Cart schema.
type Cart struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartQuery when eager-loading is set.
	Edges      CartEdges `json:"edges"`
	user_carts *int
}

// CartEdges holds the relations/edges for other nodes in the graph.
type CartEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// CartProducts holds the value of the cart_products edge.
	CartProducts []*CartProduct `json:"cart_products,omitempty"`
	// Checkout holds the value of the checkout edge.
	Checkout *Checkout `json:"checkout,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CartProductsOrErr returns the CartProducts value or an error if the edge
// was not loaded in eager-loading.
func (e CartEdges) CartProductsOrErr() ([]*CartProduct, error) {
	if e.loadedTypes[1] {
		return e.CartProducts, nil
	}
	return nil, &NotLoadedError{edge: "cart_products"}
}

// CheckoutOrErr returns the Checkout value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartEdges) CheckoutOrErr() (*Checkout, error) {
	if e.loadedTypes[2] {
		if e.Checkout == nil {
			// The edge checkout was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: checkout.Label}
		}
		return e.Checkout, nil
	}
	return nil, &NotLoadedError{edge: "checkout"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cart) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			values[i] = new(sql.NullInt64)
		case cart.FieldSlug:
			values[i] = new(sql.NullString)
		case cart.FieldCreatedAt, cart.FieldUpdatedAt, cart.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case cart.ForeignKeys[0]: // user_carts
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cart", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cart fields.
func (c *Cart) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cart.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				c.Slug = value.String
			}
		case cart.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cart.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case cart.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case cart.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_carts", value)
			} else if value.Valid {
				c.user_carts = new(int)
				*c.user_carts = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Cart entity.
func (c *Cart) QueryUser() *UserQuery {
	return (&CartClient{config: c.config}).QueryUser(c)
}

// QueryCartProducts queries the "cart_products" edge of the Cart entity.
func (c *Cart) QueryCartProducts() *CartProductQuery {
	return (&CartClient{config: c.config}).QueryCartProducts(c)
}

// QueryCheckout queries the "checkout" edge of the Cart entity.
func (c *Cart) QueryCheckout() *CheckoutQuery {
	return (&CartClient{config: c.config}).QueryCheckout(c)
}

// Update returns a builder for updating this Cart.
// Note that you need to call Cart.Unwrap() before calling this method if this Cart
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cart) Update() *CartUpdateOne {
	return (&CartClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cart) Unwrap() *Cart {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cart is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cart) String() string {
	var builder strings.Builder
	builder.WriteString("Cart(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", slug=")
	builder.WriteString(c.Slug)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	if v := c.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Carts is a parsable slice of Cart.
type Carts []*Cart

func (c Carts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}