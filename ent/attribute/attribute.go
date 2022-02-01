// Code generated by entc, DO NOT EDIT.

package attribute

import (
	"time"
)

const (
	// Label holds the string label denoting the attribute type in the database.
	Label = "attribute"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeAttribute holds the string denoting the attribute edge name in mutations.
	EdgeAttribute = "attribute"
	// Table holds the table name of the attribute in the database.
	Table = "attributes"
	// AttributeTable is the table that holds the attribute relation/edge.
	AttributeTable = "seller_product_variation_values"
	// AttributeInverseTable is the table name for the SellerProductVariationValues entity.
	// It exists in this package in order to avoid circular dependency with the "sellerproductvariationvalues" package.
	AttributeInverseTable = "seller_product_variation_values"
	// AttributeColumn is the table column denoting the attribute relation/edge.
	AttributeColumn = "attribute_attribute"
)

// Columns holds all SQL columns for attribute fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)