// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/staffmember"
)

// StaffMemberCreate is the builder for creating a StaffMember entity.
type StaffMemberCreate struct {
	config
	mutation *StaffMemberMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (smc *StaffMemberCreate) SetName(s string) *StaffMemberCreate {
	smc.mutation.SetName(s)
	return smc
}

// SetEmail sets the "email" field.
func (smc *StaffMemberCreate) SetEmail(s string) *StaffMemberCreate {
	smc.mutation.SetEmail(s)
	return smc
}

// SetPassword sets the "password" field.
func (smc *StaffMemberCreate) SetPassword(s string) *StaffMemberCreate {
	smc.mutation.SetPassword(s)
	return smc
}

// SetRole sets the "role" field.
func (smc *StaffMemberCreate) SetRole(s staffmember.Role) *StaffMemberCreate {
	smc.mutation.SetRole(s)
	return smc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (smc *StaffMemberCreate) SetNillableRole(s *staffmember.Role) *StaffMemberCreate {
	if s != nil {
		smc.SetRole(*s)
	}
	return smc
}

// SetStatus sets the "status" field.
func (smc *StaffMemberCreate) SetStatus(s staffmember.Status) *StaffMemberCreate {
	smc.mutation.SetStatus(s)
	return smc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smc *StaffMemberCreate) SetNillableStatus(s *staffmember.Status) *StaffMemberCreate {
	if s != nil {
		smc.SetStatus(*s)
	}
	return smc
}

// AddProcessedOrderIDs adds the "processed_orders" edge to the Order entity by IDs.
func (smc *StaffMemberCreate) AddProcessedOrderIDs(ids ...int) *StaffMemberCreate {
	smc.mutation.AddProcessedOrderIDs(ids...)
	return smc
}

// AddProcessedOrders adds the "processed_orders" edges to the Order entity.
func (smc *StaffMemberCreate) AddProcessedOrders(o ...*Order) *StaffMemberCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return smc.AddProcessedOrderIDs(ids...)
}

// Mutation returns the StaffMemberMutation object of the builder.
func (smc *StaffMemberCreate) Mutation() *StaffMemberMutation {
	return smc.mutation
}

// Save creates the StaffMember in the database.
func (smc *StaffMemberCreate) Save(ctx context.Context) (*StaffMember, error) {
	smc.defaults()
	return withHooks(ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *StaffMemberCreate) SaveX(ctx context.Context) *StaffMember {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *StaffMemberCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *StaffMemberCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smc *StaffMemberCreate) defaults() {
	if _, ok := smc.mutation.Role(); !ok {
		v := staffmember.DefaultRole
		smc.mutation.SetRole(v)
	}
	if _, ok := smc.mutation.Status(); !ok {
		v := staffmember.DefaultStatus
		smc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *StaffMemberCreate) check() error {
	if _, ok := smc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "StaffMember.name"`)}
	}
	if v, ok := smc.mutation.Name(); ok {
		if err := staffmember.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "StaffMember.name": %w`, err)}
		}
	}
	if _, ok := smc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "StaffMember.email"`)}
	}
	if v, ok := smc.mutation.Email(); ok {
		if err := staffmember.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "StaffMember.email": %w`, err)}
		}
	}
	if _, ok := smc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "StaffMember.password"`)}
	}
	if _, ok := smc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "StaffMember.role"`)}
	}
	if v, ok := smc.mutation.Role(); ok {
		if err := staffmember.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "StaffMember.role": %w`, err)}
		}
	}
	if _, ok := smc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "StaffMember.status"`)}
	}
	if v, ok := smc.mutation.Status(); ok {
		if err := staffmember.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "StaffMember.status": %w`, err)}
		}
	}
	return nil
}

func (smc *StaffMemberCreate) sqlSave(ctx context.Context) (*StaffMember, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *StaffMemberCreate) createSpec() (*StaffMember, *sqlgraph.CreateSpec) {
	var (
		_node = &StaffMember{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(staffmember.Table, sqlgraph.NewFieldSpec(staffmember.FieldID, field.TypeInt))
	)
	if value, ok := smc.mutation.Name(); ok {
		_spec.SetField(staffmember.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := smc.mutation.Email(); ok {
		_spec.SetField(staffmember.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := smc.mutation.Password(); ok {
		_spec.SetField(staffmember.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := smc.mutation.Role(); ok {
		_spec.SetField(staffmember.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := smc.mutation.Status(); ok {
		_spec.SetField(staffmember.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := smc.mutation.ProcessedOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffmember.ProcessedOrdersTable,
			Columns: staffmember.ProcessedOrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StaffMemberCreateBulk is the builder for creating many StaffMember entities in bulk.
type StaffMemberCreateBulk struct {
	config
	err      error
	builders []*StaffMemberCreate
}

// Save creates the StaffMember entities in the database.
func (smcb *StaffMemberCreateBulk) Save(ctx context.Context) ([]*StaffMember, error) {
	if smcb.err != nil {
		return nil, smcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*StaffMember, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StaffMemberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *StaffMemberCreateBulk) SaveX(ctx context.Context) []*StaffMember {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *StaffMemberCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *StaffMemberCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
