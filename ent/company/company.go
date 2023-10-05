// Code generated by ent, DO NOT EDIT.

package company

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBillingContact holds the string denoting the billing_contact field in the database.
	FieldBillingContact = "billing_contact"
	// FieldBillingEmail holds the string denoting the billing_email field in the database.
	FieldBillingEmail = "billing_email"
	// FieldBillingPhone holds the string denoting the billing_phone field in the database.
	FieldBillingPhone = "billing_phone"
	// FieldBillingAddress holds the string denoting the billing_address field in the database.
	FieldBillingAddress = "billing_address"
	// EdgeCustomers holds the string denoting the customers edge name in mutations.
	EdgeCustomers = "customers"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// CustomersTable is the table that holds the customers relation/edge.
	CustomersTable = "customers"
	// CustomersInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomersInverseTable = "customers"
	// CustomersColumn is the table column denoting the customers relation/edge.
	CustomersColumn = "company_customers"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "company_orders"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBillingContact,
	FieldBillingEmail,
	FieldBillingPhone,
	FieldBillingAddress,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// BillingContactValidator is a validator for the "billing_contact" field. It is called by the builders before save.
	BillingContactValidator func(string) error
	// BillingEmailValidator is a validator for the "billing_email" field. It is called by the builders before save.
	BillingEmailValidator func(string) error
	// BillingPhoneValidator is a validator for the "billing_phone" field. It is called by the builders before save.
	BillingPhoneValidator func(string) error
	// BillingAddressValidator is a validator for the "billing_address" field. It is called by the builders before save.
	BillingAddressValidator func(string) error
)

// OrderOption defines the ordering options for the Company queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBillingContact orders the results by the billing_contact field.
func ByBillingContact(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingContact, opts...).ToFunc()
}

// ByBillingEmail orders the results by the billing_email field.
func ByBillingEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingEmail, opts...).ToFunc()
}

// ByBillingPhone orders the results by the billing_phone field.
func ByBillingPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingPhone, opts...).ToFunc()
}

// ByBillingAddress orders the results by the billing_address field.
func ByBillingAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingAddress, opts...).ToFunc()
}

// ByCustomersCount orders the results by customers count.
func ByCustomersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCustomersStep(), opts...)
	}
}

// ByCustomers orders the results by customers terms.
func ByCustomers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrdersCount orders the results by orders count.
func ByOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdersStep(), opts...)
	}
}

// ByOrders orders the results by orders terms.
func ByOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCustomersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CustomersTable, CustomersColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
	)
}