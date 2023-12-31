// Code generated by ent, DO NOT EDIT.

package cartitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the cartitem type in the database.
	Label = "cart_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// EdgeCart holds the string denoting the cart edge name in mutations.
	EdgeCart = "cart"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// Table holds the table name of the cartitem in the database.
	Table = "cart_items"
	// CartTable is the table that holds the cart relation/edge. The primary key declared below.
	CartTable = "cart_cart_items"
	// CartInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartInverseTable = "carts"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "cart_items"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "cart_item_product"
)

// Columns holds all SQL columns for cartitem fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cart_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"cart_item_product",
	"product_cart_items",
}

var (
	// CartPrimaryKey and CartColumn2 are the table columns denoting the
	// primary key for the cart relation (M2M).
	CartPrimaryKey = []string{"cart_id", "cart_item_id"}
)

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
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
)

// OrderOption defines the ordering options for the CartItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByCartCount orders the results by cart count.
func ByCartCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCartStep(), opts...)
	}
}

// ByCart orders the results by cart terms.
func ByCart(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCartStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductField orders the results by product field.
func ByProductField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), sql.OrderByField(field, opts...))
	}
}
func newCartStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CartInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CartTable, CartPrimaryKey...),
	)
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProductTable, ProductColumn),
	)
}
