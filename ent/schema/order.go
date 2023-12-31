package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			NamedValues(
				"Pending", "PENDING",
				"In Progress", "IN_PROGRESS",
				"Completed", "COMPLETED",
				"Delivered", "DELIVERED",
				"Cancelled", "CANCELLED",
				"Returned", "RETURNED",
				"Refunded", "REFUNDED",
				"Failed", "FAILED",
				"On Hold", "ON_HOLD",
			).
			Default("PENDING"),
		field.Time("placed_at").
			Default(time.Now).
			Immutable(),
		field.Float("balance_due").
			Default(0).
			Min(0),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("orders"),
		edge.To("order_items", OrderItem.Type),
		edge.To("payments", Payment.Type),
		edge.From("processed_by", StaffMember.Type).
			Ref("processed_orders"),
		edge.From("company", Company.Type).Unique().
			Ref("orders"),
	}
}
