// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/cartitem"
)

// CartItem is the model entity for the CartItem schema.
type CartItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartItemQuery when eager-loading is set.
	Edges        CartItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CartItemEdges holds the relations/edges for other nodes in the graph.
type CartItemEdges struct {
	// Cart holds the value of the cart edge.
	Cart []*Cart `json:"cart,omitempty"`
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CartOrErr returns the Cart value or an error if the edge
// was not loaded in eager-loading.
func (e CartItemEdges) CartOrErr() ([]*Cart, error) {
	if e.loadedTypes[0] {
		return e.Cart, nil
	}
	return nil, &NotLoadedError{edge: "cart"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e CartItemEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CartItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cartitem.FieldID, cartitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CartItem fields.
func (ci *CartItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cartitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ci.ID = int(value.Int64)
		case cartitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				ci.Quantity = int(value.Int64)
			}
		default:
			ci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CartItem.
// This includes values selected through modifiers, order, etc.
func (ci *CartItem) Value(name string) (ent.Value, error) {
	return ci.selectValues.Get(name)
}

// QueryCart queries the "cart" edge of the CartItem entity.
func (ci *CartItem) QueryCart() *CartQuery {
	return NewCartItemClient(ci.config).QueryCart(ci)
}

// QueryProduct queries the "product" edge of the CartItem entity.
func (ci *CartItem) QueryProduct() *ProductQuery {
	return NewCartItemClient(ci.config).QueryProduct(ci)
}

// Update returns a builder for updating this CartItem.
// Note that you need to call CartItem.Unwrap() before calling this method if this CartItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CartItem) Update() *CartItemUpdateOne {
	return NewCartItemClient(ci.config).UpdateOne(ci)
}

// Unwrap unwraps the CartItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *CartItem) Unwrap() *CartItem {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: CartItem is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CartItem) String() string {
	var builder strings.Builder
	builder.WriteString("CartItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", ci.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// CartItems is a parsable slice of CartItem.
type CartItems []*CartItem
