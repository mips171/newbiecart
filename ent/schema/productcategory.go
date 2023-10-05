package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductCategory holds the schema definition for the ProductCategory entity.
type ProductCategory struct {
	ent.Schema
}

// Fields of the ProductCategory.
func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(), // Name should be unique.
		field.Text("description").
			Default(""),
		field.String("image_url").
			Default("https://via.placeholder.com/150"),
		// ... other fields
	}
}

// Edges of the ProductCategory.
func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type), // Reverse edge.
		// ... other edges
	}
}
