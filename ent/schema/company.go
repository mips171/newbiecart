package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("billing_contact").
			NotEmpty().
			Comment("Name of the person to contact for billing"),
		field.String("billing_email").
			NotEmpty(),
		field.String("billing_phone").
			NotEmpty(),
		field.String("billing_address").
			NotEmpty(),
		field.String("tax_identifier").
			Comment("Usually government-issued tax ID or business ID such as ABN in Australia"),
		// ... Other company-specific fields ...
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customers", Customer.Type), // This is the reference edge for the 'company' edge in Customer entity.
		edge.To("orders", Order.Type),       // One company can have multiple orders.
	}
}
