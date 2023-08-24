package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StaffMember holds the schema definition for the StaffMember entity.
type StaffMember struct {
	ent.Schema
}

// Fields of the StaffMember.
func (StaffMember) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("password").
			Sensitive(), // Hide in logs
		field.Enum("role").
			NamedValues(
				"Administrator", "ADMINISTRATOR",
				"Manager", "MANAGER",
				"Sales Associate", "SALES_ASSOCIATE",
				"Support Agent", "SUPPORT_AGENT",
				"Technician", "TECHNICIAN",
				"Cashier", "CASHIER",
				"Inventory Specialist", "INVENTORY_SPECIALIST",
				"Other", "OTHER",
			).Default("SALES_ASSOCIATE"),
	}
}

// Edges of the StaffMember.
func (StaffMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("processed_orders", Order.Type),
	}
}
