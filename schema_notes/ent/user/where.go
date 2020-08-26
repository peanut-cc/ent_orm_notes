// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/peanut-cc/ent_orm_notes/schema_notes/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RequiredName applies equality check predicate on the "required_name" field. It's identical to RequiredNameEQ.
func RequiredName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequiredName), v))
	})
}

// OptionalName applies equality check predicate on the "optional_name" field. It's identical to OptionalNameEQ.
func OptionalName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOptionalName), v))
	})
}

// NilableName applies equality check predicate on the "nilable_name" field. It's identical to NilableNameEQ.
func NilableName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNilableName), v))
	})
}

// NilableName2 applies equality check predicate on the "nilable_name2" field. It's identical to NilableName2EQ.
func NilableName2(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNilableName2), v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// Age2 applies equality check predicate on the "age2" field. It's identical to Age2EQ.
func Age2(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge2), v))
	})
}

// RequiredNameEQ applies the EQ predicate on the "required_name" field.
func RequiredNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequiredName), v))
	})
}

// RequiredNameNEQ applies the NEQ predicate on the "required_name" field.
func RequiredNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRequiredName), v))
	})
}

// RequiredNameIn applies the In predicate on the "required_name" field.
func RequiredNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRequiredName), v...))
	})
}

// RequiredNameNotIn applies the NotIn predicate on the "required_name" field.
func RequiredNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRequiredName), v...))
	})
}

// RequiredNameGT applies the GT predicate on the "required_name" field.
func RequiredNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRequiredName), v))
	})
}

// RequiredNameGTE applies the GTE predicate on the "required_name" field.
func RequiredNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRequiredName), v))
	})
}

// RequiredNameLT applies the LT predicate on the "required_name" field.
func RequiredNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRequiredName), v))
	})
}

// RequiredNameLTE applies the LTE predicate on the "required_name" field.
func RequiredNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRequiredName), v))
	})
}

// RequiredNameContains applies the Contains predicate on the "required_name" field.
func RequiredNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRequiredName), v))
	})
}

// RequiredNameHasPrefix applies the HasPrefix predicate on the "required_name" field.
func RequiredNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRequiredName), v))
	})
}

// RequiredNameHasSuffix applies the HasSuffix predicate on the "required_name" field.
func RequiredNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRequiredName), v))
	})
}

// RequiredNameEqualFold applies the EqualFold predicate on the "required_name" field.
func RequiredNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRequiredName), v))
	})
}

// RequiredNameContainsFold applies the ContainsFold predicate on the "required_name" field.
func RequiredNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRequiredName), v))
	})
}

// OptionalNameEQ applies the EQ predicate on the "optional_name" field.
func OptionalNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOptionalName), v))
	})
}

// OptionalNameNEQ applies the NEQ predicate on the "optional_name" field.
func OptionalNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOptionalName), v))
	})
}

// OptionalNameIn applies the In predicate on the "optional_name" field.
func OptionalNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOptionalName), v...))
	})
}

// OptionalNameNotIn applies the NotIn predicate on the "optional_name" field.
func OptionalNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOptionalName), v...))
	})
}

// OptionalNameGT applies the GT predicate on the "optional_name" field.
func OptionalNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOptionalName), v))
	})
}

// OptionalNameGTE applies the GTE predicate on the "optional_name" field.
func OptionalNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOptionalName), v))
	})
}

// OptionalNameLT applies the LT predicate on the "optional_name" field.
func OptionalNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOptionalName), v))
	})
}

// OptionalNameLTE applies the LTE predicate on the "optional_name" field.
func OptionalNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOptionalName), v))
	})
}

// OptionalNameContains applies the Contains predicate on the "optional_name" field.
func OptionalNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOptionalName), v))
	})
}

// OptionalNameHasPrefix applies the HasPrefix predicate on the "optional_name" field.
func OptionalNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOptionalName), v))
	})
}

// OptionalNameHasSuffix applies the HasSuffix predicate on the "optional_name" field.
func OptionalNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOptionalName), v))
	})
}

// OptionalNameIsNil applies the IsNil predicate on the "optional_name" field.
func OptionalNameIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOptionalName)))
	})
}

// OptionalNameNotNil applies the NotNil predicate on the "optional_name" field.
func OptionalNameNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOptionalName)))
	})
}

// OptionalNameEqualFold applies the EqualFold predicate on the "optional_name" field.
func OptionalNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOptionalName), v))
	})
}

// OptionalNameContainsFold applies the ContainsFold predicate on the "optional_name" field.
func OptionalNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOptionalName), v))
	})
}

// NilableNameEQ applies the EQ predicate on the "nilable_name" field.
func NilableNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNilableName), v))
	})
}

// NilableNameNEQ applies the NEQ predicate on the "nilable_name" field.
func NilableNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNilableName), v))
	})
}

// NilableNameIn applies the In predicate on the "nilable_name" field.
func NilableNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNilableName), v...))
	})
}

// NilableNameNotIn applies the NotIn predicate on the "nilable_name" field.
func NilableNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNilableName), v...))
	})
}

// NilableNameGT applies the GT predicate on the "nilable_name" field.
func NilableNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNilableName), v))
	})
}

// NilableNameGTE applies the GTE predicate on the "nilable_name" field.
func NilableNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNilableName), v))
	})
}

// NilableNameLT applies the LT predicate on the "nilable_name" field.
func NilableNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNilableName), v))
	})
}

// NilableNameLTE applies the LTE predicate on the "nilable_name" field.
func NilableNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNilableName), v))
	})
}

// NilableNameContains applies the Contains predicate on the "nilable_name" field.
func NilableNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNilableName), v))
	})
}

// NilableNameHasPrefix applies the HasPrefix predicate on the "nilable_name" field.
func NilableNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNilableName), v))
	})
}

// NilableNameHasSuffix applies the HasSuffix predicate on the "nilable_name" field.
func NilableNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNilableName), v))
	})
}

// NilableNameIsNil applies the IsNil predicate on the "nilable_name" field.
func NilableNameIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNilableName)))
	})
}

// NilableNameNotNil applies the NotNil predicate on the "nilable_name" field.
func NilableNameNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNilableName)))
	})
}

// NilableNameEqualFold applies the EqualFold predicate on the "nilable_name" field.
func NilableNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNilableName), v))
	})
}

// NilableNameContainsFold applies the ContainsFold predicate on the "nilable_name" field.
func NilableNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNilableName), v))
	})
}

// NilableName2EQ applies the EQ predicate on the "nilable_name2" field.
func NilableName2EQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNilableName2), v))
	})
}

// NilableName2NEQ applies the NEQ predicate on the "nilable_name2" field.
func NilableName2NEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNilableName2), v))
	})
}

// NilableName2In applies the In predicate on the "nilable_name2" field.
func NilableName2In(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNilableName2), v...))
	})
}

// NilableName2NotIn applies the NotIn predicate on the "nilable_name2" field.
func NilableName2NotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNilableName2), v...))
	})
}

// NilableName2GT applies the GT predicate on the "nilable_name2" field.
func NilableName2GT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNilableName2), v))
	})
}

// NilableName2GTE applies the GTE predicate on the "nilable_name2" field.
func NilableName2GTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNilableName2), v))
	})
}

// NilableName2LT applies the LT predicate on the "nilable_name2" field.
func NilableName2LT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNilableName2), v))
	})
}

// NilableName2LTE applies the LTE predicate on the "nilable_name2" field.
func NilableName2LTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNilableName2), v))
	})
}

// NilableName2Contains applies the Contains predicate on the "nilable_name2" field.
func NilableName2Contains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNilableName2), v))
	})
}

// NilableName2HasPrefix applies the HasPrefix predicate on the "nilable_name2" field.
func NilableName2HasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNilableName2), v))
	})
}

// NilableName2HasSuffix applies the HasSuffix predicate on the "nilable_name2" field.
func NilableName2HasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNilableName2), v))
	})
}

// NilableName2IsNil applies the IsNil predicate on the "nilable_name2" field.
func NilableName2IsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNilableName2)))
	})
}

// NilableName2NotNil applies the NotNil predicate on the "nilable_name2" field.
func NilableName2NotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNilableName2)))
	})
}

// NilableName2EqualFold applies the EqualFold predicate on the "nilable_name2" field.
func NilableName2EqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNilableName2), v))
	})
}

// NilableName2ContainsFold applies the ContainsFold predicate on the "nilable_name2" field.
func NilableName2ContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNilableName2), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// AgeIsNil applies the IsNil predicate on the "age" field.
func AgeIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAge)))
	})
}

// AgeNotNil applies the NotNil predicate on the "age" field.
func AgeNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAge)))
	})
}

// Age2EQ applies the EQ predicate on the "age2" field.
func Age2EQ(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge2), v))
	})
}

// Age2NEQ applies the NEQ predicate on the "age2" field.
func Age2NEQ(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge2), v))
	})
}

// Age2In applies the In predicate on the "age2" field.
func Age2In(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge2), v...))
	})
}

// Age2NotIn applies the NotIn predicate on the "age2" field.
func Age2NotIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge2), v...))
	})
}

// Age2GT applies the GT predicate on the "age2" field.
func Age2GT(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge2), v))
	})
}

// Age2GTE applies the GTE predicate on the "age2" field.
func Age2GTE(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge2), v))
	})
}

// Age2LT applies the LT predicate on the "age2" field.
func Age2LT(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge2), v))
	})
}

// Age2LTE applies the LTE predicate on the "age2" field.
func Age2LTE(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge2), v))
	})
}

// Age2IsNil applies the IsNil predicate on the "age2" field.
func Age2IsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAge2)))
	})
}

// Age2NotNil applies the NotNil predicate on the "age2" field.
func Age2NotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAge2)))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
