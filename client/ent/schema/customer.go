package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MinLen(2).MaxLen(8),
		field.String("phone_number").NotEmpty().MinLen(10).MaxLen(10),
		field.String("email").NotEmpty().Unique(),
		field.String("license_id").Unique().MinLen(10),
		field.String("address").NotEmpty(),
		field.Int("membership_number").NonNegative().Nillable(),
		field.Bool("is_active").Default(true),
		field.String("password").MinLen(8).NotEmpty(),
		field.Time("create_at").Default(time.Now()),
		field.Time("update_at").Default(time.Now()),
		field.Enum("role").Values("SUBSCRIBER", "ADMIN").Default("SUBSCRIBER"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
