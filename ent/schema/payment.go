package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount").
			Positive(),
		field.Enum("payment_method").
			NamedValues(
				"Credit Card", "CREDIT_CARD",
				"PayPal", "PAYPAL",
				"Bank Transfer", "BANK_TRANSFER",
				"Debit Card", "DEBIT_CARD",
				"Cash", "CASH",
				"Check", "CHECK",
				"Mobile Payment", "MOBILE_PAYMENT",
				"Other", "OTHER",
			).
			Default("CREDIT_CARD"),
		field.String("transaction_id").
			Unique(), // Transaction ID returned from the payment gateway or system.
		field.Enum("status").
			NamedValues(
				"Pending", "PENDING",
				"Completed", "COMPLETED",
				"Failed", "FAILED",
				"Refunded", "REFUNDED",
				"Disputed", "DISPUTED",
				"Chargeback", "CHARGEBACK",
			).
			Default("PENDING"),
		field.Time("processed_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("payments"),
	}
}
