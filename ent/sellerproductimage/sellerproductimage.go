// Code generated by entc, DO NOT EDIT.

package sellerproductimage

import (
	"time"
)

const (
	// Label holds the string label denoting the sellerproductimage type in the database.
	Label = "seller_product_image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisplay holds the string denoting the display field in the database.
	FieldDisplay = "display"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeSellerProduct holds the string denoting the seller_product edge name in mutations.
	EdgeSellerProduct = "seller_product"
	// Table holds the table name of the sellerproductimage in the database.
	Table = "seller_product_images"
	// SellerProductTable is the table that holds the seller_product relation/edge.
	SellerProductTable = "seller_product_images"
	// SellerProductInverseTable is the table name for the SellerProduct entity.
	// It exists in this package in order to avoid circular dependency with the "sellerproduct" package.
	SellerProductInverseTable = "seller_products"
	// SellerProductColumn is the table column denoting the seller_product relation/edge.
	SellerProductColumn = "seller_product_seller_product_images"
)

// Columns holds all SQL columns for sellerproductimage fields.
var Columns = []string{
	FieldID,
	FieldDisplay,
	FieldImage,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "seller_product_images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"seller_product_seller_product_images",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDisplay holds the default value on creation for the "display" field.
	DefaultDisplay bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)