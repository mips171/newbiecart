// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// PlacedAt applies equality check predicate on the "placed_at" field. It's identical to PlacedAtEQ.
func PlacedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPlacedAt, v))
}

// BalanceDue applies equality check predicate on the "balance_due" field. It's identical to BalanceDueEQ.
func BalanceDue(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldBalanceDue, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatus, vs...))
}

// PlacedAtEQ applies the EQ predicate on the "placed_at" field.
func PlacedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPlacedAt, v))
}

// PlacedAtNEQ applies the NEQ predicate on the "placed_at" field.
func PlacedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldPlacedAt, v))
}

// PlacedAtIn applies the In predicate on the "placed_at" field.
func PlacedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldPlacedAt, vs...))
}

// PlacedAtNotIn applies the NotIn predicate on the "placed_at" field.
func PlacedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldPlacedAt, vs...))
}

// PlacedAtGT applies the GT predicate on the "placed_at" field.
func PlacedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldPlacedAt, v))
}

// PlacedAtGTE applies the GTE predicate on the "placed_at" field.
func PlacedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldPlacedAt, v))
}

// PlacedAtLT applies the LT predicate on the "placed_at" field.
func PlacedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldPlacedAt, v))
}

// PlacedAtLTE applies the LTE predicate on the "placed_at" field.
func PlacedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldPlacedAt, v))
}

// BalanceDueEQ applies the EQ predicate on the "balance_due" field.
func BalanceDueEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldBalanceDue, v))
}

// BalanceDueNEQ applies the NEQ predicate on the "balance_due" field.
func BalanceDueNEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldBalanceDue, v))
}

// BalanceDueIn applies the In predicate on the "balance_due" field.
func BalanceDueIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldBalanceDue, vs...))
}

// BalanceDueNotIn applies the NotIn predicate on the "balance_due" field.
func BalanceDueNotIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldBalanceDue, vs...))
}

// BalanceDueGT applies the GT predicate on the "balance_due" field.
func BalanceDueGT(v float64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldBalanceDue, v))
}

// BalanceDueGTE applies the GTE predicate on the "balance_due" field.
func BalanceDueGTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldBalanceDue, v))
}

// BalanceDueLT applies the LT predicate on the "balance_due" field.
func BalanceDueLT(v float64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldBalanceDue, v))
}

// BalanceDueLTE applies the LTE predicate on the "balance_due" field.
func BalanceDueLTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldBalanceDue, v))
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CustomerTable, CustomerPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderItems applies the HasEdge predicate on the "order_items" edge.
func HasOrderItems() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrderItemsTable, OrderItemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderItemsWith applies the HasEdge predicate on the "order_items" edge with a given conditions (other predicates).
func HasOrderItemsWith(preds ...predicate.OrderItem) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newOrderItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayments applies the HasEdge predicate on the "payments" edge.
func HasPayments() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PaymentsTable, PaymentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentsWith applies the HasEdge predicate on the "payments" edge with a given conditions (other predicates).
func HasPaymentsWith(preds ...predicate.Payment) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newPaymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProcessedBy applies the HasEdge predicate on the "processed_by" edge.
func HasProcessedBy() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProcessedByTable, ProcessedByPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProcessedByWith applies the HasEdge predicate on the "processed_by" edge with a given conditions (other predicates).
func HasProcessedByWith(preds ...predicate.StaffMember) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newProcessedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}
