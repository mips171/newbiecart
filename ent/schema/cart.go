package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return nil
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cart_items", CartItem.Type),
		edge.From("customer", Customer.Type).
			Ref("cart").
			Unique(),
	}
}
