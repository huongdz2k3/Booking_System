// Code generated by ent, DO NOT EDIT.

package flight

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the flight type in the database.
	Label = "flight"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldDepartDate holds the string denoting the depart_date field in the database.
	FieldDepartDate = "depart_date"
	// FieldDepartTime holds the string denoting the depart_time field in the database.
	FieldDepartTime = "depart_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAvailableSlots holds the string denoting the available_slots field in the database.
	FieldAvailableSlots = "available_slots"
	// FieldReturnDate holds the string denoting the return_date field in the database.
	FieldReturnDate = "return_date"
	// FieldFlightPlane holds the string denoting the flight_plane field in the database.
	FieldFlightPlane = "flight_plane"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the flight in the database.
	Table = "flights"
)

// Columns holds all SQL columns for flight fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFrom,
	FieldTo,
	FieldDepartDate,
	FieldDepartTime,
	FieldStatus,
	FieldAvailableSlots,
	FieldReturnDate,
	FieldFlightPlane,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AvailableSlotsValidator is a validator for the "available_slots" field. It is called by the builders before save.
	AvailableSlotsValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusSCHEDULED is the default value of the Status enum.
const DefaultStatus = StatusSCHEDULED

// Status values.
const (
	StatusON_TIME   Status = "ON_TIME"
	StatusDELAYED   Status = "DELAYED"
	StatusCANCELLED Status = "CANCELLED"
	StatusSCHEDULED Status = "SCHEDULED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusON_TIME, StatusDELAYED, StatusCANCELLED, StatusSCHEDULED:
		return nil
	default:
		return fmt.Errorf("flight: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Flight queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFrom orders the results by the from field.
func ByFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrom, opts...).ToFunc()
}

// ByTo orders the results by the to field.
func ByTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTo, opts...).ToFunc()
}

// ByDepartDate orders the results by the depart_date field.
func ByDepartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartDate, opts...).ToFunc()
}

// ByDepartTime orders the results by the depart_time field.
func ByDepartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAvailableSlots orders the results by the available_slots field.
func ByAvailableSlots(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailableSlots, opts...).ToFunc()
}

// ByReturnDate orders the results by the return_date field.
func ByReturnDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReturnDate, opts...).ToFunc()
}

// ByFlightPlane orders the results by the flight_plane field.
func ByFlightPlane(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFlightPlane, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
