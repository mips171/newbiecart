// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/cart"
	"github.com/mikestefanello/pagoda/ent/cartitem"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/product"
)

// CartItemUpdate is the builder for updating CartItem entities.
type CartItemUpdate struct {
	config
	hooks    []Hook
	mutation *CartItemMutation
}

// Where appends a list predicates to the CartItemUpdate builder.
func (ciu *CartItemUpdate) Where(ps ...predicate.CartItem) *CartItemUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetQuantity sets the "quantity" field.
func (ciu *CartItemUpdate) SetQuantity(i int) *CartItemUpdate {
	ciu.mutation.ResetQuantity()
	ciu.mutation.SetQuantity(i)
	return ciu
}

// AddQuantity adds i to the "quantity" field.
func (ciu *CartItemUpdate) AddQuantity(i int) *CartItemUpdate {
	ciu.mutation.AddQuantity(i)
	return ciu
}

// AddCartIDs adds the "cart" edge to the Cart entity by IDs.
func (ciu *CartItemUpdate) AddCartIDs(ids ...int) *CartItemUpdate {
	ciu.mutation.AddCartIDs(ids...)
	return ciu
}

// AddCart adds the "cart" edges to the Cart entity.
func (ciu *CartItemUpdate) AddCart(c ...*Cart) *CartItemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ciu.AddCartIDs(ids...)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (ciu *CartItemUpdate) SetProductID(id int) *CartItemUpdate {
	ciu.mutation.SetProductID(id)
	return ciu
}

// SetProduct sets the "product" edge to the Product entity.
func (ciu *CartItemUpdate) SetProduct(p *Product) *CartItemUpdate {
	return ciu.SetProductID(p.ID)
}

// Mutation returns the CartItemMutation object of the builder.
func (ciu *CartItemUpdate) Mutation() *CartItemMutation {
	return ciu.mutation
}

// ClearCart clears all "cart" edges to the Cart entity.
func (ciu *CartItemUpdate) ClearCart() *CartItemUpdate {
	ciu.mutation.ClearCart()
	return ciu
}

// RemoveCartIDs removes the "cart" edge to Cart entities by IDs.
func (ciu *CartItemUpdate) RemoveCartIDs(ids ...int) *CartItemUpdate {
	ciu.mutation.RemoveCartIDs(ids...)
	return ciu
}

// RemoveCart removes "cart" edges to Cart entities.
func (ciu *CartItemUpdate) RemoveCart(c ...*Cart) *CartItemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ciu.RemoveCartIDs(ids...)
}

// ClearProduct clears the "product" edge to the Product entity.
func (ciu *CartItemUpdate) ClearProduct() *CartItemUpdate {
	ciu.mutation.ClearProduct()
	return ciu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CartItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CartItemUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CartItemUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CartItemUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *CartItemUpdate) check() error {
	if v, ok := ciu.mutation.Quantity(); ok {
		if err := cartitem.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "CartItem.quantity": %w`, err)}
		}
	}
	if _, ok := ciu.mutation.ProductID(); ciu.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CartItem.product"`)
	}
	return nil
}

func (ciu *CartItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ciu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeInt))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.Quantity(); ok {
		_spec.SetField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := ciu.mutation.AddedQuantity(); ok {
		_spec.AddField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if ciu.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: cartitem.CartPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.RemovedCartIDs(); len(nodes) > 0 && !ciu.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: cartitem.CartPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: cartitem.CartPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cartitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// CartItemUpdateOne is the builder for updating a single CartItem entity.
type CartItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartItemMutation
}

// SetQuantity sets the "quantity" field.
func (ciuo *CartItemUpdateOne) SetQuantity(i int) *CartItemUpdateOne {
	ciuo.mutation.ResetQuantity()
	ciuo.mutation.SetQuantity(i)
	return ciuo
}

// AddQuantity adds i to the "quantity" field.
func (ciuo *CartItemUpdateOne) AddQuantity(i int) *CartItemUpdateOne {
	ciuo.mutation.AddQuantity(i)
	return ciuo
}

// AddCartIDs adds the "cart" edge to the Cart entity by IDs.
func (ciuo *CartItemUpdateOne) AddCartIDs(ids ...int) *CartItemUpdateOne {
	ciuo.mutation.AddCartIDs(ids...)
	return ciuo
}

// AddCart adds the "cart" edges to the Cart entity.
func (ciuo *CartItemUpdateOne) AddCart(c ...*Cart) *CartItemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ciuo.AddCartIDs(ids...)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (ciuo *CartItemUpdateOne) SetProductID(id int) *CartItemUpdateOne {
	ciuo.mutation.SetProductID(id)
	return ciuo
}

// SetProduct sets the "product" edge to the Product entity.
func (ciuo *CartItemUpdateOne) SetProduct(p *Product) *CartItemUpdateOne {
	return ciuo.SetProductID(p.ID)
}

// Mutation returns the CartItemMutation object of the builder.
func (ciuo *CartItemUpdateOne) Mutation() *CartItemMutation {
	return ciuo.mutation
}

// ClearCart clears all "cart" edges to the Cart entity.
func (ciuo *CartItemUpdateOne) ClearCart() *CartItemUpdateOne {
	ciuo.mutation.ClearCart()
	return ciuo
}

// RemoveCartIDs removes the "cart" edge to Cart entities by IDs.
func (ciuo *CartItemUpdateOne) RemoveCartIDs(ids ...int) *CartItemUpdateOne {
	ciuo.mutation.RemoveCartIDs(ids...)
	return ciuo
}

// RemoveCart removes "cart" edges to Cart entities.
func (ciuo *CartItemUpdateOne) RemoveCart(c ...*Cart) *CartItemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ciuo.RemoveCartIDs(ids...)
}

// ClearProduct clears the "product" edge to the Product entity.
func (ciuo *CartItemUpdateOne) ClearProduct() *CartItemUpdateOne {
	ciuo.mutation.ClearProduct()
	return ciuo
}

// Where appends a list predicates to the CartItemUpdate builder.
func (ciuo *CartItemUpdateOne) Where(ps ...predicate.CartItem) *CartItemUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CartItemUpdateOne) Select(field string, fields ...string) *CartItemUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CartItem entity.
func (ciuo *CartItemUpdateOne) Save(ctx context.Context) (*CartItem, error) {
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CartItemUpdateOne) SaveX(ctx context.Context) *CartItem {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CartItemUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CartItemUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *CartItemUpdateOne) check() error {
	if v, ok := ciuo.mutation.Quantity(); ok {
		if err := cartitem.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "CartItem.quantity": %w`, err)}
		}
	}
	if _, ok := ciuo.mutation.ProductID(); ciuo.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CartItem.product"`)
	}
	return nil
}

func (ciuo *CartItemUpdateOne) sqlSave(ctx context.Context) (_node *CartItem, err error) {
	if err := ciuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeInt))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CartItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cartitem.FieldID)
		for _, f := range fields {
			if !cartitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cartitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.Quantity(); ok {
		_spec.SetField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := ciuo.mutation.AddedQuantity(); ok {
		_spec.AddField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if ciuo.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: cartitem.CartPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.RemovedCartIDs(); len(nodes) > 0 && !ciuo.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: cartitem.CartPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: cartitem.CartPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CartItem{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cartitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
