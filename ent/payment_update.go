// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/payment"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PaymentUpdate) AddAmount(f float64) *PaymentUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetPaymentMethod sets the "payment_method" field.
func (pu *PaymentUpdate) SetPaymentMethod(pm payment.PaymentMethod) *PaymentUpdate {
	pu.mutation.SetPaymentMethod(pm)
	return pu
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillablePaymentMethod(pm *payment.PaymentMethod) *PaymentUpdate {
	if pm != nil {
		pu.SetPaymentMethod(*pm)
	}
	return pu
}

// SetTransactionID sets the "transaction_id" field.
func (pu *PaymentUpdate) SetTransactionID(s string) *PaymentUpdate {
	pu.mutation.SetTransactionID(s)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PaymentUpdate) SetStatus(pa payment.Status) *PaymentUpdate {
	pu.mutation.SetStatus(pa)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableStatus(pa *payment.Status) *PaymentUpdate {
	if pa != nil {
		pu.SetStatus(*pa)
	}
	return pu
}

// AddOrderIDs adds the "order" edge to the Order entity by IDs.
func (pu *PaymentUpdate) AddOrderIDs(ids ...int) *PaymentUpdate {
	pu.mutation.AddOrderIDs(ids...)
	return pu
}

// AddOrder adds the "order" edges to the Order entity.
func (pu *PaymentUpdate) AddOrder(o ...*Order) *PaymentUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddOrderIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearOrder clears all "order" edges to the Order entity.
func (pu *PaymentUpdate) ClearOrder() *PaymentUpdate {
	pu.mutation.ClearOrder()
	return pu
}

// RemoveOrderIDs removes the "order" edge to Order entities by IDs.
func (pu *PaymentUpdate) RemoveOrderIDs(ids ...int) *PaymentUpdate {
	pu.mutation.RemoveOrderIDs(ids...)
	return pu
}

// RemoveOrder removes "order" edges to Order entities.
func (pu *PaymentUpdate) RemoveOrder(o ...*Order) *PaymentUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.PaymentMethod(); ok {
		if err := payment.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Payment.payment_method": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.PaymentMethod(); ok {
		_spec.SetField(payment.FieldPaymentMethod, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.TransactionID(); ok {
		_spec.SetField(payment.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   payment.OrderTable,
			Columns: payment.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedOrderIDs(); len(nodes) > 0 && !pu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   payment.OrderTable,
			Columns: payment.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   payment.OrderTable,
			Columns: payment.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetPaymentMethod sets the "payment_method" field.
func (puo *PaymentUpdateOne) SetPaymentMethod(pm payment.PaymentMethod) *PaymentUpdateOne {
	puo.mutation.SetPaymentMethod(pm)
	return puo
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillablePaymentMethod(pm *payment.PaymentMethod) *PaymentUpdateOne {
	if pm != nil {
		puo.SetPaymentMethod(*pm)
	}
	return puo
}

// SetTransactionID sets the "transaction_id" field.
func (puo *PaymentUpdateOne) SetTransactionID(s string) *PaymentUpdateOne {
	puo.mutation.SetTransactionID(s)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PaymentUpdateOne) SetStatus(pa payment.Status) *PaymentUpdateOne {
	puo.mutation.SetStatus(pa)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableStatus(pa *payment.Status) *PaymentUpdateOne {
	if pa != nil {
		puo.SetStatus(*pa)
	}
	return puo
}

// AddOrderIDs adds the "order" edge to the Order entity by IDs.
func (puo *PaymentUpdateOne) AddOrderIDs(ids ...int) *PaymentUpdateOne {
	puo.mutation.AddOrderIDs(ids...)
	return puo
}

// AddOrder adds the "order" edges to the Order entity.
func (puo *PaymentUpdateOne) AddOrder(o ...*Order) *PaymentUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddOrderIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearOrder clears all "order" edges to the Order entity.
func (puo *PaymentUpdateOne) ClearOrder() *PaymentUpdateOne {
	puo.mutation.ClearOrder()
	return puo
}

// RemoveOrderIDs removes the "order" edge to Order entities by IDs.
func (puo *PaymentUpdateOne) RemoveOrderIDs(ids ...int) *PaymentUpdateOne {
	puo.mutation.RemoveOrderIDs(ids...)
	return puo
}

// RemoveOrder removes "order" edges to Order entities.
func (puo *PaymentUpdateOne) RemoveOrder(o ...*Order) *PaymentUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveOrderIDs(ids...)
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.PaymentMethod(); ok {
		if err := payment.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Payment.payment_method": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.PaymentMethod(); ok {
		_spec.SetField(payment.FieldPaymentMethod, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.TransactionID(); ok {
		_spec.SetField(payment.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   payment.OrderTable,
			Columns: payment.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedOrderIDs(); len(nodes) > 0 && !puo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   payment.OrderTable,
			Columns: payment.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   payment.OrderTable,
			Columns: payment.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
