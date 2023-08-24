// Code generated by ent, DO NOT EDIT.

package payment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldID, id))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldAmount, v))
}

// TransactionID applies equality check predicate on the "transaction_id" field. It's identical to TransactionIDEQ.
func TransactionID(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldTransactionID, v))
}

// ProcessedAt applies equality check predicate on the "processed_at" field. It's identical to ProcessedAtEQ.
func ProcessedAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldProcessedAt, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldAmount, v))
}

// PaymentMethodEQ applies the EQ predicate on the "payment_method" field.
func PaymentMethodEQ(v PaymentMethod) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentMethod, v))
}

// PaymentMethodNEQ applies the NEQ predicate on the "payment_method" field.
func PaymentMethodNEQ(v PaymentMethod) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaymentMethod, v))
}

// PaymentMethodIn applies the In predicate on the "payment_method" field.
func PaymentMethodIn(vs ...PaymentMethod) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaymentMethod, vs...))
}

// PaymentMethodNotIn applies the NotIn predicate on the "payment_method" field.
func PaymentMethodNotIn(vs ...PaymentMethod) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaymentMethod, vs...))
}

// TransactionIDEQ applies the EQ predicate on the "transaction_id" field.
func TransactionIDEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldTransactionID, v))
}

// TransactionIDNEQ applies the NEQ predicate on the "transaction_id" field.
func TransactionIDNEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldTransactionID, v))
}

// TransactionIDIn applies the In predicate on the "transaction_id" field.
func TransactionIDIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldTransactionID, vs...))
}

// TransactionIDNotIn applies the NotIn predicate on the "transaction_id" field.
func TransactionIDNotIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldTransactionID, vs...))
}

// TransactionIDGT applies the GT predicate on the "transaction_id" field.
func TransactionIDGT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldTransactionID, v))
}

// TransactionIDGTE applies the GTE predicate on the "transaction_id" field.
func TransactionIDGTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldTransactionID, v))
}

// TransactionIDLT applies the LT predicate on the "transaction_id" field.
func TransactionIDLT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldTransactionID, v))
}

// TransactionIDLTE applies the LTE predicate on the "transaction_id" field.
func TransactionIDLTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldTransactionID, v))
}

// TransactionIDContains applies the Contains predicate on the "transaction_id" field.
func TransactionIDContains(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContains(FieldTransactionID, v))
}

// TransactionIDHasPrefix applies the HasPrefix predicate on the "transaction_id" field.
func TransactionIDHasPrefix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasPrefix(FieldTransactionID, v))
}

// TransactionIDHasSuffix applies the HasSuffix predicate on the "transaction_id" field.
func TransactionIDHasSuffix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasSuffix(FieldTransactionID, v))
}

// TransactionIDEqualFold applies the EqualFold predicate on the "transaction_id" field.
func TransactionIDEqualFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEqualFold(FieldTransactionID, v))
}

// TransactionIDContainsFold applies the ContainsFold predicate on the "transaction_id" field.
func TransactionIDContainsFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContainsFold(FieldTransactionID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldStatus, vs...))
}

// ProcessedAtEQ applies the EQ predicate on the "processed_at" field.
func ProcessedAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldProcessedAt, v))
}

// ProcessedAtNEQ applies the NEQ predicate on the "processed_at" field.
func ProcessedAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldProcessedAt, v))
}

// ProcessedAtIn applies the In predicate on the "processed_at" field.
func ProcessedAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldProcessedAt, vs...))
}

// ProcessedAtNotIn applies the NotIn predicate on the "processed_at" field.
func ProcessedAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldProcessedAt, vs...))
}

// ProcessedAtGT applies the GT predicate on the "processed_at" field.
func ProcessedAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldProcessedAt, v))
}

// ProcessedAtGTE applies the GTE predicate on the "processed_at" field.
func ProcessedAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldProcessedAt, v))
}

// ProcessedAtLT applies the LT predicate on the "processed_at" field.
func ProcessedAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldProcessedAt, v))
}

// ProcessedAtLTE applies the LTE predicate on the "processed_at" field.
func ProcessedAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldProcessedAt, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrderTable, OrderPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
