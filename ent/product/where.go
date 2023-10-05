// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// Sku applies equality check predicate on the "sku" field. It's identical to SkuEQ.
func Sku(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSku, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// StockCount applies equality check predicate on the "stock_count" field. It's identical to StockCountEQ.
func StockCount(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStockCount, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldImageURL, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldIsActive, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// SkuEQ applies the EQ predicate on the "sku" field.
func SkuEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSku, v))
}

// SkuNEQ applies the NEQ predicate on the "sku" field.
func SkuNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldSku, v))
}

// SkuIn applies the In predicate on the "sku" field.
func SkuIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldSku, vs...))
}

// SkuNotIn applies the NotIn predicate on the "sku" field.
func SkuNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldSku, vs...))
}

// SkuGT applies the GT predicate on the "sku" field.
func SkuGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldSku, v))
}

// SkuGTE applies the GTE predicate on the "sku" field.
func SkuGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldSku, v))
}

// SkuLT applies the LT predicate on the "sku" field.
func SkuLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldSku, v))
}

// SkuLTE applies the LTE predicate on the "sku" field.
func SkuLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldSku, v))
}

// SkuContains applies the Contains predicate on the "sku" field.
func SkuContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldSku, v))
}

// SkuHasPrefix applies the HasPrefix predicate on the "sku" field.
func SkuHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldSku, v))
}

// SkuHasSuffix applies the HasSuffix predicate on the "sku" field.
func SkuHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldSku, v))
}

// SkuEqualFold applies the EqualFold predicate on the "sku" field.
func SkuEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldSku, v))
}

// SkuContainsFold applies the ContainsFold predicate on the "sku" field.
func SkuContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldSku, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldDescription, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPrice, v))
}

// PriceContains applies the Contains predicate on the "price" field.
func PriceContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldPrice, v))
}

// PriceHasPrefix applies the HasPrefix predicate on the "price" field.
func PriceHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldPrice, v))
}

// PriceHasSuffix applies the HasSuffix predicate on the "price" field.
func PriceHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldPrice, v))
}

// PriceEqualFold applies the EqualFold predicate on the "price" field.
func PriceEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldPrice, v))
}

// PriceContainsFold applies the ContainsFold predicate on the "price" field.
func PriceContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldPrice, v))
}

// StockCountEQ applies the EQ predicate on the "stock_count" field.
func StockCountEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStockCount, v))
}

// StockCountNEQ applies the NEQ predicate on the "stock_count" field.
func StockCountNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldStockCount, v))
}

// StockCountIn applies the In predicate on the "stock_count" field.
func StockCountIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldStockCount, vs...))
}

// StockCountNotIn applies the NotIn predicate on the "stock_count" field.
func StockCountNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldStockCount, vs...))
}

// StockCountGT applies the GT predicate on the "stock_count" field.
func StockCountGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldStockCount, v))
}

// StockCountGTE applies the GTE predicate on the "stock_count" field.
func StockCountGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldStockCount, v))
}

// StockCountLT applies the LT predicate on the "stock_count" field.
func StockCountLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldStockCount, v))
}

// StockCountLTE applies the LTE predicate on the "stock_count" field.
func StockCountLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldStockCount, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldImageURL, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldIsActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasCartItems applies the HasEdge predicate on the "cart_items" edge.
func HasCartItems() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CartItemsTable, CartItemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCartItemsWith applies the HasEdge predicate on the "cart_items" edge with a given conditions (other predicates).
func HasCartItemsWith(preds ...predicate.CartItem) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newCartItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderItems applies the HasEdge predicate on the "order_items" edge.
func HasOrderItems() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrderItemsTable, OrderItemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderItemsWith applies the HasEdge predicate on the "order_items" edge with a given conditions (other predicates).
func HasOrderItemsWith(preds ...predicate.OrderItem) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newOrderItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategories applies the HasEdge predicate on the "categories" edge.
func HasCategories() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoriesTable, CategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriesWith applies the HasEdge predicate on the "categories" edge with a given conditions (other predicates).
func HasCategoriesWith(preds ...predicate.ProductCategory) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
