// Code generated by ent, DO NOT EDIT.

package providerkey

import (
	"domain-manager/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldName, v))
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldProvider, v))
}

// AccessKey applies equality check predicate on the "access_key" field. It's identical to AccessKeyEQ.
func AccessKey(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldAccessKey, v))
}

// SecretKey applies equality check predicate on the "secret_key" field. It's identical to SecretKeyEQ.
func SecretKey(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldSecretKey, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContainsFold(FieldName, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldProvider, vs...))
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldProvider, v))
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldProvider, v))
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldProvider, v))
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldProvider, v))
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContains(FieldProvider, v))
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasPrefix(FieldProvider, v))
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasSuffix(FieldProvider, v))
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEqualFold(FieldProvider, v))
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContainsFold(FieldProvider, v))
}

// AccessKeyEQ applies the EQ predicate on the "access_key" field.
func AccessKeyEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldAccessKey, v))
}

// AccessKeyNEQ applies the NEQ predicate on the "access_key" field.
func AccessKeyNEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldAccessKey, v))
}

// AccessKeyIn applies the In predicate on the "access_key" field.
func AccessKeyIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldAccessKey, vs...))
}

// AccessKeyNotIn applies the NotIn predicate on the "access_key" field.
func AccessKeyNotIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldAccessKey, vs...))
}

// AccessKeyGT applies the GT predicate on the "access_key" field.
func AccessKeyGT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldAccessKey, v))
}

// AccessKeyGTE applies the GTE predicate on the "access_key" field.
func AccessKeyGTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldAccessKey, v))
}

// AccessKeyLT applies the LT predicate on the "access_key" field.
func AccessKeyLT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldAccessKey, v))
}

// AccessKeyLTE applies the LTE predicate on the "access_key" field.
func AccessKeyLTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldAccessKey, v))
}

// AccessKeyContains applies the Contains predicate on the "access_key" field.
func AccessKeyContains(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContains(FieldAccessKey, v))
}

// AccessKeyHasPrefix applies the HasPrefix predicate on the "access_key" field.
func AccessKeyHasPrefix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasPrefix(FieldAccessKey, v))
}

// AccessKeyHasSuffix applies the HasSuffix predicate on the "access_key" field.
func AccessKeyHasSuffix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasSuffix(FieldAccessKey, v))
}

// AccessKeyEqualFold applies the EqualFold predicate on the "access_key" field.
func AccessKeyEqualFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEqualFold(FieldAccessKey, v))
}

// AccessKeyContainsFold applies the ContainsFold predicate on the "access_key" field.
func AccessKeyContainsFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContainsFold(FieldAccessKey, v))
}

// SecretKeyEQ applies the EQ predicate on the "secret_key" field.
func SecretKeyEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldSecretKey, v))
}

// SecretKeyNEQ applies the NEQ predicate on the "secret_key" field.
func SecretKeyNEQ(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldSecretKey, v))
}

// SecretKeyIn applies the In predicate on the "secret_key" field.
func SecretKeyIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldSecretKey, vs...))
}

// SecretKeyNotIn applies the NotIn predicate on the "secret_key" field.
func SecretKeyNotIn(vs ...string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldSecretKey, vs...))
}

// SecretKeyGT applies the GT predicate on the "secret_key" field.
func SecretKeyGT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldSecretKey, v))
}

// SecretKeyGTE applies the GTE predicate on the "secret_key" field.
func SecretKeyGTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldSecretKey, v))
}

// SecretKeyLT applies the LT predicate on the "secret_key" field.
func SecretKeyLT(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldSecretKey, v))
}

// SecretKeyLTE applies the LTE predicate on the "secret_key" field.
func SecretKeyLTE(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldSecretKey, v))
}

// SecretKeyContains applies the Contains predicate on the "secret_key" field.
func SecretKeyContains(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContains(FieldSecretKey, v))
}

// SecretKeyHasPrefix applies the HasPrefix predicate on the "secret_key" field.
func SecretKeyHasPrefix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasPrefix(FieldSecretKey, v))
}

// SecretKeyHasSuffix applies the HasSuffix predicate on the "secret_key" field.
func SecretKeyHasSuffix(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldHasSuffix(FieldSecretKey, v))
}

// SecretKeyEqualFold applies the EqualFold predicate on the "secret_key" field.
func SecretKeyEqualFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEqualFold(FieldSecretKey, v))
}

// SecretKeyContainsFold applies the ContainsFold predicate on the "secret_key" field.
func SecretKeyContainsFold(v string) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldContainsFold(FieldSecretKey, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProviderKey {
	return predicate.ProviderKey(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ProviderKey {
	return predicate.ProviderKey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ProviderKey {
	return predicate.ProviderKey(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProviderKey) predicate.ProviderKey {
	return predicate.ProviderKey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProviderKey) predicate.ProviderKey {
	return predicate.ProviderKey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProviderKey) predicate.ProviderKey {
	return predicate.ProviderKey(sql.NotPredicates(p))
}
