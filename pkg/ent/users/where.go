// Code generated by entc, DO NOT EDIT.

package users

import (
	"shutter/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Pseudo applies equality check predicate on the "pseudo" field. It's identical to PseudoEQ.
func Pseudo(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPseudo), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// Fname applies equality check predicate on the "fname" field. It's identical to FnameEQ.
func Fname(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFname), v))
	})
}

// Lname applies equality check predicate on the "lname" field. It's identical to LnameEQ.
func Lname(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLname), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// PseudoEQ applies the EQ predicate on the "pseudo" field.
func PseudoEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPseudo), v))
	})
}

// PseudoNEQ applies the NEQ predicate on the "pseudo" field.
func PseudoNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPseudo), v))
	})
}

// PseudoIn applies the In predicate on the "pseudo" field.
func PseudoIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPseudo), v...))
	})
}

// PseudoNotIn applies the NotIn predicate on the "pseudo" field.
func PseudoNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPseudo), v...))
	})
}

// PseudoGT applies the GT predicate on the "pseudo" field.
func PseudoGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPseudo), v))
	})
}

// PseudoGTE applies the GTE predicate on the "pseudo" field.
func PseudoGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPseudo), v))
	})
}

// PseudoLT applies the LT predicate on the "pseudo" field.
func PseudoLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPseudo), v))
	})
}

// PseudoLTE applies the LTE predicate on the "pseudo" field.
func PseudoLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPseudo), v))
	})
}

// PseudoContains applies the Contains predicate on the "pseudo" field.
func PseudoContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPseudo), v))
	})
}

// PseudoHasPrefix applies the HasPrefix predicate on the "pseudo" field.
func PseudoHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPseudo), v))
	})
}

// PseudoHasSuffix applies the HasSuffix predicate on the "pseudo" field.
func PseudoHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPseudo), v))
	})
}

// PseudoEqualFold applies the EqualFold predicate on the "pseudo" field.
func PseudoEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPseudo), v))
	})
}

// PseudoContainsFold applies the ContainsFold predicate on the "pseudo" field.
func PseudoContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPseudo), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// FnameEQ applies the EQ predicate on the "fname" field.
func FnameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFname), v))
	})
}

// FnameNEQ applies the NEQ predicate on the "fname" field.
func FnameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFname), v))
	})
}

// FnameIn applies the In predicate on the "fname" field.
func FnameIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFname), v...))
	})
}

// FnameNotIn applies the NotIn predicate on the "fname" field.
func FnameNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFname), v...))
	})
}

// FnameGT applies the GT predicate on the "fname" field.
func FnameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFname), v))
	})
}

// FnameGTE applies the GTE predicate on the "fname" field.
func FnameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFname), v))
	})
}

// FnameLT applies the LT predicate on the "fname" field.
func FnameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFname), v))
	})
}

// FnameLTE applies the LTE predicate on the "fname" field.
func FnameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFname), v))
	})
}

// FnameContains applies the Contains predicate on the "fname" field.
func FnameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFname), v))
	})
}

// FnameHasPrefix applies the HasPrefix predicate on the "fname" field.
func FnameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFname), v))
	})
}

// FnameHasSuffix applies the HasSuffix predicate on the "fname" field.
func FnameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFname), v))
	})
}

// FnameEqualFold applies the EqualFold predicate on the "fname" field.
func FnameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFname), v))
	})
}

// FnameContainsFold applies the ContainsFold predicate on the "fname" field.
func FnameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFname), v))
	})
}

// LnameEQ applies the EQ predicate on the "lname" field.
func LnameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLname), v))
	})
}

// LnameNEQ applies the NEQ predicate on the "lname" field.
func LnameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLname), v))
	})
}

// LnameIn applies the In predicate on the "lname" field.
func LnameIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLname), v...))
	})
}

// LnameNotIn applies the NotIn predicate on the "lname" field.
func LnameNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLname), v...))
	})
}

// LnameGT applies the GT predicate on the "lname" field.
func LnameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLname), v))
	})
}

// LnameGTE applies the GTE predicate on the "lname" field.
func LnameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLname), v))
	})
}

// LnameLT applies the LT predicate on the "lname" field.
func LnameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLname), v))
	})
}

// LnameLTE applies the LTE predicate on the "lname" field.
func LnameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLname), v))
	})
}

// LnameContains applies the Contains predicate on the "lname" field.
func LnameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLname), v))
	})
}

// LnameHasPrefix applies the HasPrefix predicate on the "lname" field.
func LnameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLname), v))
	})
}

// LnameHasSuffix applies the HasSuffix predicate on the "lname" field.
func LnameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLname), v))
	})
}

// LnameEqualFold applies the EqualFold predicate on the "lname" field.
func LnameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLname), v))
	})
}

// LnameContainsFold applies the ContainsFold predicate on the "lname" field.
func LnameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLname), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
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
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		p(s.Not())
	})
}
