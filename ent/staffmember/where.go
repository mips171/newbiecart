// Code generated by ent, DO NOT EDIT.

package staffmember

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldPassword, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldContainsFold(FieldPassword, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNotIn(FieldRole, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.StaffMember {
	return predicate.StaffMember(sql.FieldNotIn(FieldStatus, vs...))
}

// HasProcessedOrders applies the HasEdge predicate on the "processed_orders" edge.
func HasProcessedOrders() predicate.StaffMember {
	return predicate.StaffMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ProcessedOrdersTable, ProcessedOrdersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProcessedOrdersWith applies the HasEdge predicate on the "processed_orders" edge with a given conditions (other predicates).
func HasProcessedOrdersWith(preds ...predicate.Order) predicate.StaffMember {
	return predicate.StaffMember(func(s *sql.Selector) {
		step := newProcessedOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StaffMember) predicate.StaffMember {
	return predicate.StaffMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StaffMember) predicate.StaffMember {
	return predicate.StaffMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StaffMember) predicate.StaffMember {
	return predicate.StaffMember(func(s *sql.Selector) {
		p(s.Not())
	})
}
