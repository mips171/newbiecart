// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/productcategory"
)

// ProductCategory is the model entity for the ProductCategory schema.
type ProductCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductCategoryQuery when eager-loading is set.
	Edges        ProductCategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductCategoryEdges holds the relations/edges for other nodes in the graph.
type ProductCategoryEdges struct {
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e ProductCategoryEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case productcategory.FieldName, productcategory.FieldDescription, productcategory.FieldImageURL:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductCategory fields.
func (pc *ProductCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		case productcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pc.Name = value.String
			}
		case productcategory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pc.Description = value.String
			}
		case productcategory.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				pc.ImageURL = value.String
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductCategory.
// This includes values selected through modifiers, order, etc.
func (pc *ProductCategory) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryProducts queries the "products" edge of the ProductCategory entity.
func (pc *ProductCategory) QueryProducts() *ProductQuery {
	return NewProductCategoryClient(pc.config).QueryProducts(pc)
}

// Update returns a builder for updating this ProductCategory.
// Note that you need to call ProductCategory.Unwrap() before calling this method if this ProductCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *ProductCategory) Update() *ProductCategoryUpdateOne {
	return NewProductCategoryClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the ProductCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *ProductCategory) Unwrap() *ProductCategory {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductCategory is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *ProductCategory) String() string {
	var builder strings.Builder
	builder.WriteString("ProductCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("name=")
	builder.WriteString(pc.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pc.Description)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(pc.ImageURL)
	builder.WriteByte(')')
	return builder.String()
}

// ProductCategories is a parsable slice of ProductCategory.
type ProductCategories []*ProductCategory
