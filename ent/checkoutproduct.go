// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/checkout"
	"bongo/ent/checkoutproduct"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

// CheckoutProduct is the model entity for the CheckoutProduct schema.
type CheckoutProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// SellingPrice holds the value of the "selling_price" field.
	SellingPrice decimal.Decimal `json:"selling_price,omitempty"`
	// OfferPrice holds the value of the "offer_price" field.
	OfferPrice int `json:"offer_price,omitempty"`
	// Received holds the value of the "received" field.
	Received bool `json:"received,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckoutProductQuery when eager-loading is set.
	Edges                                      CheckoutProductEdges `json:"edges"`
	checkout_checkout_products                 *int
	seller_product_checkout_products           *int
	seller_product_variation_checkout_products *int
	user_checkout_products                     *int
	user_seller_checkout_products              *int
}

// CheckoutProductEdges holds the relations/edges for other nodes in the graph.
type CheckoutProductEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Checkout holds the value of the checkout edge.
	Checkout *Checkout `json:"checkout,omitempty"`
	// Seller holds the value of the seller edge.
	Seller *User `json:"seller,omitempty"`
	// SellerProduct holds the value of the seller_product edge.
	SellerProduct *SellerProduct `json:"seller_product,omitempty"`
	// SellerProductVariation holds the value of the seller_product_variation edge.
	SellerProductVariation *SellerProductVariation `json:"seller_product_variation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutProductEdges) UserOrErr() (*User, error) {
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

// CheckoutOrErr returns the Checkout value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutProductEdges) CheckoutOrErr() (*Checkout, error) {
	if e.loadedTypes[1] {
		if e.Checkout == nil {
			// The edge checkout was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: checkout.Label}
		}
		return e.Checkout, nil
	}
	return nil, &NotLoadedError{edge: "checkout"}
}

// SellerOrErr returns the Seller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutProductEdges) SellerOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Seller == nil {
			// The edge seller was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Seller, nil
	}
	return nil, &NotLoadedError{edge: "seller"}
}

// SellerProductOrErr returns the SellerProduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutProductEdges) SellerProductOrErr() (*SellerProduct, error) {
	if e.loadedTypes[3] {
		if e.SellerProduct == nil {
			// The edge seller_product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sellerproduct.Label}
		}
		return e.SellerProduct, nil
	}
	return nil, &NotLoadedError{edge: "seller_product"}
}

// SellerProductVariationOrErr returns the SellerProductVariation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutProductEdges) SellerProductVariationOrErr() (*SellerProductVariation, error) {
	if e.loadedTypes[4] {
		if e.SellerProductVariation == nil {
			// The edge seller_product_variation was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sellerproductvariation.Label}
		}
		return e.SellerProductVariation, nil
	}
	return nil, &NotLoadedError{edge: "seller_product_variation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CheckoutProduct) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case checkoutproduct.FieldSellingPrice:
			values[i] = new(decimal.Decimal)
		case checkoutproduct.FieldReceived:
			values[i] = new(sql.NullBool)
		case checkoutproduct.FieldID, checkoutproduct.FieldQuantity, checkoutproduct.FieldOfferPrice, checkoutproduct.FieldStatus:
			values[i] = new(sql.NullInt64)
		case checkoutproduct.FieldCreatedAt, checkoutproduct.FieldUpdatedAt, checkoutproduct.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case checkoutproduct.ForeignKeys[0]: // checkout_checkout_products
			values[i] = new(sql.NullInt64)
		case checkoutproduct.ForeignKeys[1]: // seller_product_checkout_products
			values[i] = new(sql.NullInt64)
		case checkoutproduct.ForeignKeys[2]: // seller_product_variation_checkout_products
			values[i] = new(sql.NullInt64)
		case checkoutproduct.ForeignKeys[3]: // user_checkout_products
			values[i] = new(sql.NullInt64)
		case checkoutproduct.ForeignKeys[4]: // user_seller_checkout_products
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CheckoutProduct", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CheckoutProduct fields.
func (cp *CheckoutProduct) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case checkoutproduct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cp.ID = int(value.Int64)
		case checkoutproduct.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				cp.Quantity = int(value.Int64)
			}
		case checkoutproduct.FieldSellingPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field selling_price", values[i])
			} else if value != nil {
				cp.SellingPrice = *value
			}
		case checkoutproduct.FieldOfferPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field offer_price", values[i])
			} else if value.Valid {
				cp.OfferPrice = int(value.Int64)
			}
		case checkoutproduct.FieldReceived:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field received", values[i])
			} else if value.Valid {
				cp.Received = value.Bool
			}
		case checkoutproduct.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cp.Status = int(value.Int64)
			}
		case checkoutproduct.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cp.CreatedAt = value.Time
			}
		case checkoutproduct.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cp.UpdatedAt = value.Time
			}
		case checkoutproduct.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cp.DeletedAt = new(time.Time)
				*cp.DeletedAt = value.Time
			}
		case checkoutproduct.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field checkout_checkout_products", value)
			} else if value.Valid {
				cp.checkout_checkout_products = new(int)
				*cp.checkout_checkout_products = int(value.Int64)
			}
		case checkoutproduct.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field seller_product_checkout_products", value)
			} else if value.Valid {
				cp.seller_product_checkout_products = new(int)
				*cp.seller_product_checkout_products = int(value.Int64)
			}
		case checkoutproduct.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field seller_product_variation_checkout_products", value)
			} else if value.Valid {
				cp.seller_product_variation_checkout_products = new(int)
				*cp.seller_product_variation_checkout_products = int(value.Int64)
			}
		case checkoutproduct.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_checkout_products", value)
			} else if value.Valid {
				cp.user_checkout_products = new(int)
				*cp.user_checkout_products = int(value.Int64)
			}
		case checkoutproduct.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_seller_checkout_products", value)
			} else if value.Valid {
				cp.user_seller_checkout_products = new(int)
				*cp.user_seller_checkout_products = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the CheckoutProduct entity.
func (cp *CheckoutProduct) QueryUser() *UserQuery {
	return (&CheckoutProductClient{config: cp.config}).QueryUser(cp)
}

// QueryCheckout queries the "checkout" edge of the CheckoutProduct entity.
func (cp *CheckoutProduct) QueryCheckout() *CheckoutQuery {
	return (&CheckoutProductClient{config: cp.config}).QueryCheckout(cp)
}

// QuerySeller queries the "seller" edge of the CheckoutProduct entity.
func (cp *CheckoutProduct) QuerySeller() *UserQuery {
	return (&CheckoutProductClient{config: cp.config}).QuerySeller(cp)
}

// QuerySellerProduct queries the "seller_product" edge of the CheckoutProduct entity.
func (cp *CheckoutProduct) QuerySellerProduct() *SellerProductQuery {
	return (&CheckoutProductClient{config: cp.config}).QuerySellerProduct(cp)
}

// QuerySellerProductVariation queries the "seller_product_variation" edge of the CheckoutProduct entity.
func (cp *CheckoutProduct) QuerySellerProductVariation() *SellerProductVariationQuery {
	return (&CheckoutProductClient{config: cp.config}).QuerySellerProductVariation(cp)
}

// Update returns a builder for updating this CheckoutProduct.
// Note that you need to call CheckoutProduct.Unwrap() before calling this method if this CheckoutProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (cp *CheckoutProduct) Update() *CheckoutProductUpdateOne {
	return (&CheckoutProductClient{config: cp.config}).UpdateOne(cp)
}

// Unwrap unwraps the CheckoutProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cp *CheckoutProduct) Unwrap() *CheckoutProduct {
	tx, ok := cp.config.driver.(*txDriver)
	if !ok {
		panic("ent: CheckoutProduct is not a transactional entity")
	}
	cp.config.driver = tx.drv
	return cp
}

// String implements the fmt.Stringer.
func (cp *CheckoutProduct) String() string {
	var builder strings.Builder
	builder.WriteString("CheckoutProduct(")
	builder.WriteString(fmt.Sprintf("id=%v", cp.ID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", cp.Quantity))
	builder.WriteString(", selling_price=")
	builder.WriteString(fmt.Sprintf("%v", cp.SellingPrice))
	builder.WriteString(", offer_price=")
	builder.WriteString(fmt.Sprintf("%v", cp.OfferPrice))
	builder.WriteString(", received=")
	builder.WriteString(fmt.Sprintf("%v", cp.Received))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", cp.Status))
	builder.WriteString(", created_at=")
	builder.WriteString(cp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cp.UpdatedAt.Format(time.ANSIC))
	if v := cp.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// CheckoutProducts is a parsable slice of CheckoutProduct.
type CheckoutProducts []*CheckoutProduct

func (cp CheckoutProducts) config(cfg config) {
	for _i := range cp {
		cp[_i].config = cfg
	}
}
