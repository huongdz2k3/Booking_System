package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entproto.Field(2)),
		field.String("from").Annotations(entproto.Field(3)),
		field.String("to").Annotations(entproto.Field(4)),
		field.Time("depart_date").Annotations(entproto.Field(5)),
		field.Time("depart_time").Annotations(entproto.Field(6)),
		field.Time("return_date").Annotations(entproto.Field(12)).Nillable(),
		field.String("status").Annotations(entproto.Field(7)),
		field.Int("available_slots").NonNegative().Annotations(entproto.Field(8)),
		field.String("flight_plane").Annotations(entproto.Field(9)),
		field.Time("created_at").Default(time.Now()).Annotations(entproto.Field(10)),
		field.Time("updated_at").Default(time.Now()).Annotations(entproto.Field(11)),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return nil
}
