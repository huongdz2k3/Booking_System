// Code generated by ent, DO NOT EDIT.

package ent

import (
	"customer/ent/flight"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Flight is the model entity for the Flight schema.
type Flight struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// DepartDate holds the value of the "depart_date" field.
	DepartDate time.Time `json:"depart_date,omitempty"`
	// DepartTime holds the value of the "depart_time" field.
	DepartTime time.Time `json:"depart_time,omitempty"`
	// Status holds the value of the "status" field.
	Status flight.Status `json:"status,omitempty"`
	// AvailableSlots holds the value of the "available_slots" field.
	AvailableSlots int `json:"available_slots,omitempty"`
	// ReturnDate holds the value of the "return_date" field.
	ReturnDate *time.Time `json:"return_date,omitempty"`
	// FlightPlane holds the value of the "flight_plane" field.
	FlightPlane string `json:"flight_plane,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flight) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case flight.FieldID, flight.FieldAvailableSlots:
			values[i] = new(sql.NullInt64)
		case flight.FieldName, flight.FieldFrom, flight.FieldTo, flight.FieldStatus, flight.FieldFlightPlane:
			values[i] = new(sql.NullString)
		case flight.FieldDepartDate, flight.FieldDepartTime, flight.FieldReturnDate, flight.FieldCreatedAt, flight.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Flight fields.
func (f *Flight) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flight.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case flight.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case flight.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				f.From = value.String
			}
		case flight.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				f.To = value.String
			}
		case flight.FieldDepartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field depart_date", values[i])
			} else if value.Valid {
				f.DepartDate = value.Time
			}
		case flight.FieldDepartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field depart_time", values[i])
			} else if value.Valid {
				f.DepartTime = value.Time
			}
		case flight.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = flight.Status(value.String)
			}
		case flight.FieldAvailableSlots:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field available_slots", values[i])
			} else if value.Valid {
				f.AvailableSlots = int(value.Int64)
			}
		case flight.FieldReturnDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field return_date", values[i])
			} else if value.Valid {
				f.ReturnDate = new(time.Time)
				*f.ReturnDate = value.Time
			}
		case flight.FieldFlightPlane:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_plane", values[i])
			} else if value.Valid {
				f.FlightPlane = value.String
			}
		case flight.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case flight.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Flight.
// This includes values selected through modifiers, order, etc.
func (f *Flight) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Flight.
// Note that you need to call Flight.Unwrap() before calling this method if this Flight
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Flight) Update() *FlightUpdateOne {
	return NewFlightClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Flight entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Flight) Unwrap() *Flight {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Flight is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Flight) String() string {
	var builder strings.Builder
	builder.WriteString("Flight(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(f.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(f.To)
	builder.WriteString(", ")
	builder.WriteString("depart_date=")
	builder.WriteString(f.DepartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("depart_time=")
	builder.WriteString(f.DepartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", f.Status))
	builder.WriteString(", ")
	builder.WriteString("available_slots=")
	builder.WriteString(fmt.Sprintf("%v", f.AvailableSlots))
	builder.WriteString(", ")
	if v := f.ReturnDate; v != nil {
		builder.WriteString("return_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("flight_plane=")
	builder.WriteString(f.FlightPlane)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Flights is a parsable slice of Flight.
type Flights []*Flight
