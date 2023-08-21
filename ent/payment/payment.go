// Code generated by ent, DO NOT EDIT.

package payment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldPaymentMethod holds the string denoting the payment_method field in the database.
	FieldPaymentMethod = "payment_method"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldProcessedAt holds the string denoting the processed_at field in the database.
	FieldProcessedAt = "processed_at"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// Table holds the table name of the payment in the database.
	Table = "payments"
	// OrderTable is the table that holds the order relation/edge. The primary key declared below.
	OrderTable = "order_payments"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldPaymentMethod,
	FieldTransactionID,
	FieldStatus,
	FieldProcessedAt,
}

var (
	// OrderPrimaryKey and OrderColumn2 are the table columns denoting the
	// primary key for the order relation (M2M).
	OrderPrimaryKey = []string{"order_id", "payment_id"}
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
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(float64) error
	// PaymentMethodValidator is a validator for the "payment_method" field. It is called by the builders before save.
	PaymentMethodValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
	// DefaultProcessedAt holds the default value on creation for the "processed_at" field.
	DefaultProcessedAt func() time.Time
)

// OrderOption defines the ordering options for the Payment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByPaymentMethod orders the results by the payment_method field.
func ByPaymentMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentMethod, opts...).ToFunc()
}

// ByTransactionID orders the results by the transaction_id field.
func ByTransactionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByProcessedAt orders the results by the processed_at field.
func ByProcessedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProcessedAt, opts...).ToFunc()
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
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OrderTable, OrderPrimaryKey...),
	)
}