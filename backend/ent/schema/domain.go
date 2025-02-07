package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name").NotEmpty(),     // 域名
		field.String("provider").NotEmpty(), // 服务提供商
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("domains").Unique(),
	}
}
