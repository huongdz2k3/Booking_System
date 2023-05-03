package schema

import (
	"entgo.io/contrib/entproto"
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
		field.String("name").NotEmpty().MinLen(2).MaxLen(8).
			Annotations(
				entproto.Field(2),
			),
		field.String("phone_number").NotEmpty().MinLen(10).MaxLen(10).
			Annotations(
				entproto.Field(3),
			),
		field.String("email").NotEmpty().Unique().
			Annotations(
				entproto.Field(4),
			),
		field.String("license_id").Unique().MinLen(10).
			Annotations(
				entproto.Field(5),
			),
		field.String("address").NotEmpty().
			Annotations(
				entproto.Field(6),
			),
		field.Int("membership_number").NonNegative().Optional().
			Annotations(
				entproto.Field(7),
			),
		field.Bool("is_active").Default(true).
			Annotations(
				entproto.Field(8),
			),
		field.String("password").MinLen(8).NotEmpty().
			Annotations(
				entproto.Field(9),
			),
		field.Time("create_at").Default(time.Now()).
			Annotations(
				entproto.Field(10),
			),
		field.Time("update_at").Default(time.Now()).
			Annotations(
				entproto.Field(11),
			),
		field.String("role").NotEmpty().
			Annotations(
				entproto.Field(12),
			),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(), entproto.Service(),
	}
}
