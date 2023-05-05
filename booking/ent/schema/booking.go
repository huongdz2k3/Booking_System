package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.String("booking_code").NotEmpty(),
		field.Time("booking_date"),
		field.Time("cancel_date").Nillable(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Int("flight_id").NonNegative(),
		field.Int("customer_id").NonNegative(),
		field.Enum("status").Values("CANCEL", "SUCCESS", "PROCESS").Default("PROCESS"),
		field.String("customer_name").Nillable(),
		field.String("email").Nillable(),
		field.String("license_id").Nillable(),
		field.String("address").Nillable(),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return nil
}
