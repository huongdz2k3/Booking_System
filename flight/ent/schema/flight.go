package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("from"),
		field.String("to"),
		field.Time("depart_date"),
		field.Time("depart_time"),
		field.Enum("status").Values("ON_TIME", "DELAYED", "CANCELLED", "SCHEDULED").Default("SCHEDULED"),
		field.Int("available_slots").NonNegative(),
		field.Time("return_date").Nillable().Optional(),
		field.Enum("type").Values("ONE_WAY", "RETURN_TICKET"),
		field.String("flight_plane"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return nil
}
