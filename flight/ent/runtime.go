// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flight-service/ent/flight"
	"flight-service/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	flightFields := schema.Flight{}.Fields()
	_ = flightFields
	// flightDescAvailableSlots is the schema descriptor for available_slots field.
	flightDescAvailableSlots := flightFields[7].Descriptor()
	// flight.AvailableSlotsValidator is a validator for the "available_slots" field. It is called by the builders before save.
	flight.AvailableSlotsValidator = flightDescAvailableSlots.Validators[0].(func(int) error)
	// flightDescCreatedAt is the schema descriptor for created_at field.
	flightDescCreatedAt := flightFields[9].Descriptor()
	// flight.DefaultCreatedAt holds the default value on creation for the created_at field.
	flight.DefaultCreatedAt = flightDescCreatedAt.Default.(time.Time)
	// flightDescUpdatedAt is the schema descriptor for updated_at field.
	flightDescUpdatedAt := flightFields[10].Descriptor()
	// flight.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	flight.DefaultUpdatedAt = flightDescUpdatedAt.Default.(time.Time)
}
