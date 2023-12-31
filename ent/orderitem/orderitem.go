// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orderitem type in the database.
	Label = "order_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUnitPrice holds the string denoting the unit_price field in the database.
	FieldUnitPrice = "unit_price"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// Table holds the table name of the orderitem in the database.
	Table = "order_items"
	// ProductTable is the table that holds the product relation/edge. The primary key declared below.
	ProductTable = "product_order_items"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// OrderTable is the table that holds the order relation/edge. The primary key declared below.
	OrderTable = "order_order_items"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
)

// Columns holds all SQL columns for orderitem fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
	FieldUnitPrice,
}

var (
	// ProductPrimaryKey and ProductColumn2 are the table columns denoting the
	// primary key for the product relation (M2M).
	ProductPrimaryKey = []string{"product_id", "order_item_id"}
	// OrderPrimaryKey and OrderColumn2 are the table columns denoting the
	// primary key for the order relation (M2M).
	OrderPrimaryKey = []string{"order_id", "order_item_id"}
)

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
	// DefaultQuantity holds the default value on creation for the "quantity" field.
	DefaultQuantity int
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
)

// OrderOption defines the ordering options for the OrderItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByUnitPrice orders the results by the unit_price field.
func ByUnitPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnitPrice, opts...).ToFunc()
}

// ByProductCount orders the results by product count.
func ByProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductStep(), opts...)
	}
}

// ByProduct orders the results by product terms.
func ByProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrderCount orders the results by order count.
func ByOrderCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderStep(), opts...)
	}
}

// ByOrder orders the results by order terms.
func ByOrder(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
	)
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OrderTable, OrderPrimaryKey...),
	)
}
