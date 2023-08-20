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
		field.String("payment_method").
			NotEmpty(), // Could be values like "Credit Card", "PayPal", "Bank Transfer", etc.
		field.String("transaction_id").
			Unique(), // Transaction ID returned from the payment gateway or system.
		field.String("status").
			Default("pending"). // Typical values: "pending", "completed", "failed", etc.
			NotEmpty(),
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
