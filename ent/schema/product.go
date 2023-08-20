package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Text("description"),
		field.Float("price").
			Positive(),
		field.Int("stock_count").
			Default(0).
			NonNegative(),
		field.String("image_url"),
	}

}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cart_items", CartItem.Type),
		edge.To("order_items", OrderItem.Type),
		edge.From("category", ProductCategory.Type).
			Ref("products"),
	}
}
