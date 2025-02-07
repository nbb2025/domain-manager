package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProviderKey holds the schema definition for the ProviderKey entity.
type ProviderKey struct {
	ent.Schema
}

// Fields of the ProviderKey.
func (ProviderKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name").NotEmpty().Default("default"), // 密钥名称
		field.String("provider").NotEmpty(),                // 服务提供商（如：aliyun, cloudflare, tencent, dnspod）
		field.String("access_key").NotEmpty(),
		field.String("secret_key"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the ProviderKey.
func (ProviderKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("provider_keys").Unique(),
	}
}
