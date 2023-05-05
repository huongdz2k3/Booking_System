// Code generated by ent, DO NOT EDIT.

package flight

import (
	"flight-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldName, v))
}

// From applies equality check predicate on the "from" field. It's identical to FromEQ.
func From(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldFrom, v))
}

// To applies equality check predicate on the "to" field. It's identical to ToEQ.
func To(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldTo, v))
}

// DepartDate applies equality check predicate on the "depart_date" field. It's identical to DepartDateEQ.
func DepartDate(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartDate, v))
}

// DepartTime applies equality check predicate on the "depart_time" field. It's identical to DepartTimeEQ.
func DepartTime(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartTime, v))
}

// ReturnDate applies equality check predicate on the "return_date" field. It's identical to ReturnDateEQ.
func ReturnDate(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldReturnDate, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldStatus, v))
}

// AvailableSlots applies equality check predicate on the "available_slots" field. It's identical to AvailableSlotsEQ.
func AvailableSlots(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableSlots, v))
}

// FlightPlane applies equality check predicate on the "flight_plane" field. It's identical to FlightPlaneEQ.
func FlightPlane(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldFlightPlane, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldName, v))
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldFrom, v))
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldFrom, v))
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldFrom, vs...))
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldFrom, vs...))
}

// FromGT applies the GT predicate on the "from" field.
func FromGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldFrom, v))
}

// FromGTE applies the GTE predicate on the "from" field.
func FromGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldFrom, v))
}

// FromLT applies the LT predicate on the "from" field.
func FromLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldFrom, v))
}

// FromLTE applies the LTE predicate on the "from" field.
func FromLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldFrom, v))
}

// FromContains applies the Contains predicate on the "from" field.
func FromContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldFrom, v))
}

// FromHasPrefix applies the HasPrefix predicate on the "from" field.
func FromHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldFrom, v))
}

// FromHasSuffix applies the HasSuffix predicate on the "from" field.
func FromHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldFrom, v))
}

// FromEqualFold applies the EqualFold predicate on the "from" field.
func FromEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldFrom, v))
}

// FromContainsFold applies the ContainsFold predicate on the "from" field.
func FromContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldFrom, v))
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldTo, v))
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldTo, v))
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldTo, vs...))
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldTo, vs...))
}

// ToGT applies the GT predicate on the "to" field.
func ToGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldTo, v))
}

// ToGTE applies the GTE predicate on the "to" field.
func ToGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldTo, v))
}

// ToLT applies the LT predicate on the "to" field.
func ToLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldTo, v))
}

// ToLTE applies the LTE predicate on the "to" field.
func ToLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldTo, v))
}

// ToContains applies the Contains predicate on the "to" field.
func ToContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldTo, v))
}

// ToHasPrefix applies the HasPrefix predicate on the "to" field.
func ToHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldTo, v))
}

// ToHasSuffix applies the HasSuffix predicate on the "to" field.
func ToHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldTo, v))
}

// ToEqualFold applies the EqualFold predicate on the "to" field.
func ToEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldTo, v))
}

// ToContainsFold applies the ContainsFold predicate on the "to" field.
func ToContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldTo, v))
}

// DepartDateEQ applies the EQ predicate on the "depart_date" field.
func DepartDateEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartDate, v))
}

// DepartDateNEQ applies the NEQ predicate on the "depart_date" field.
func DepartDateNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldDepartDate, v))
}

// DepartDateIn applies the In predicate on the "depart_date" field.
func DepartDateIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldDepartDate, vs...))
}

// DepartDateNotIn applies the NotIn predicate on the "depart_date" field.
func DepartDateNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldDepartDate, vs...))
}

// DepartDateGT applies the GT predicate on the "depart_date" field.
func DepartDateGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldDepartDate, v))
}

// DepartDateGTE applies the GTE predicate on the "depart_date" field.
func DepartDateGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldDepartDate, v))
}

// DepartDateLT applies the LT predicate on the "depart_date" field.
func DepartDateLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldDepartDate, v))
}

// DepartDateLTE applies the LTE predicate on the "depart_date" field.
func DepartDateLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldDepartDate, v))
}

// DepartTimeEQ applies the EQ predicate on the "depart_time" field.
func DepartTimeEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartTime, v))
}

// DepartTimeNEQ applies the NEQ predicate on the "depart_time" field.
func DepartTimeNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldDepartTime, v))
}

// DepartTimeIn applies the In predicate on the "depart_time" field.
func DepartTimeIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldDepartTime, vs...))
}

// DepartTimeNotIn applies the NotIn predicate on the "depart_time" field.
func DepartTimeNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldDepartTime, vs...))
}

// DepartTimeGT applies the GT predicate on the "depart_time" field.
func DepartTimeGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldDepartTime, v))
}

// DepartTimeGTE applies the GTE predicate on the "depart_time" field.
func DepartTimeGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldDepartTime, v))
}

// DepartTimeLT applies the LT predicate on the "depart_time" field.
func DepartTimeLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldDepartTime, v))
}

// DepartTimeLTE applies the LTE predicate on the "depart_time" field.
func DepartTimeLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldDepartTime, v))
}

// ReturnDateEQ applies the EQ predicate on the "return_date" field.
func ReturnDateEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldReturnDate, v))
}

// ReturnDateNEQ applies the NEQ predicate on the "return_date" field.
func ReturnDateNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldReturnDate, v))
}

// ReturnDateIn applies the In predicate on the "return_date" field.
func ReturnDateIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldReturnDate, vs...))
}

// ReturnDateNotIn applies the NotIn predicate on the "return_date" field.
func ReturnDateNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldReturnDate, vs...))
}

// ReturnDateGT applies the GT predicate on the "return_date" field.
func ReturnDateGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldReturnDate, v))
}

// ReturnDateGTE applies the GTE predicate on the "return_date" field.
func ReturnDateGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldReturnDate, v))
}

// ReturnDateLT applies the LT predicate on the "return_date" field.
func ReturnDateLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldReturnDate, v))
}

// ReturnDateLTE applies the LTE predicate on the "return_date" field.
func ReturnDateLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldReturnDate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldStatus, v))
}

// AvailableSlotsEQ applies the EQ predicate on the "available_slots" field.
func AvailableSlotsEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableSlots, v))
}

// AvailableSlotsNEQ applies the NEQ predicate on the "available_slots" field.
func AvailableSlotsNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldAvailableSlots, v))
}

// AvailableSlotsIn applies the In predicate on the "available_slots" field.
func AvailableSlotsIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldAvailableSlots, vs...))
}

// AvailableSlotsNotIn applies the NotIn predicate on the "available_slots" field.
func AvailableSlotsNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldAvailableSlots, vs...))
}

// AvailableSlotsGT applies the GT predicate on the "available_slots" field.
func AvailableSlotsGT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldAvailableSlots, v))
}

// AvailableSlotsGTE applies the GTE predicate on the "available_slots" field.
func AvailableSlotsGTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldAvailableSlots, v))
}

// AvailableSlotsLT applies the LT predicate on the "available_slots" field.
func AvailableSlotsLT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldAvailableSlots, v))
}

// AvailableSlotsLTE applies the LTE predicate on the "available_slots" field.
func AvailableSlotsLTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldAvailableSlots, v))
}

// FlightPlaneEQ applies the EQ predicate on the "flight_plane" field.
func FlightPlaneEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldFlightPlane, v))
}

// FlightPlaneNEQ applies the NEQ predicate on the "flight_plane" field.
func FlightPlaneNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldFlightPlane, v))
}

// FlightPlaneIn applies the In predicate on the "flight_plane" field.
func FlightPlaneIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldFlightPlane, vs...))
}

// FlightPlaneNotIn applies the NotIn predicate on the "flight_plane" field.
func FlightPlaneNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldFlightPlane, vs...))
}

// FlightPlaneGT applies the GT predicate on the "flight_plane" field.
func FlightPlaneGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldFlightPlane, v))
}

// FlightPlaneGTE applies the GTE predicate on the "flight_plane" field.
func FlightPlaneGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldFlightPlane, v))
}

// FlightPlaneLT applies the LT predicate on the "flight_plane" field.
func FlightPlaneLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldFlightPlane, v))
}

// FlightPlaneLTE applies the LTE predicate on the "flight_plane" field.
func FlightPlaneLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldFlightPlane, v))
}

// FlightPlaneContains applies the Contains predicate on the "flight_plane" field.
func FlightPlaneContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldFlightPlane, v))
}

// FlightPlaneHasPrefix applies the HasPrefix predicate on the "flight_plane" field.
func FlightPlaneHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldFlightPlane, v))
}

// FlightPlaneHasSuffix applies the HasSuffix predicate on the "flight_plane" field.
func FlightPlaneHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldFlightPlane, v))
}

// FlightPlaneEqualFold applies the EqualFold predicate on the "flight_plane" field.
func FlightPlaneEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldFlightPlane, v))
}

// FlightPlaneContainsFold applies the ContainsFold predicate on the "flight_plane" field.
func FlightPlaneContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldFlightPlane, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		p(s.Not())
	})
}
