package schema

import (
	"time"

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
		field.String("sku").
			NotEmpty().Unique(),
		field.Text("description").
			Default(""),
		field.String("price").
			Default("0.00"),
		field.Int("stock_count").
			Default(0).
			NonNegative(),
		field.String("image_url").
			Default("https://via.placeholder.com/150"),
		field.Bool("is_active").
			Default(true),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// ... other fields
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
