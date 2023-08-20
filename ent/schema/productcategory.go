package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type ProductCategory struct {
	ent.Schema
}

// Fields of the Category.
func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}

}

// Edges of the ProductCategory.
func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}
