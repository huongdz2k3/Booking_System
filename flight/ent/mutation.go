// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flight-service/ent/flight"
	"flight-service/ent/predicate"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFlight = "Flight"
)

// FlightMutation represents an operation that mutates the Flight nodes in the graph.
type FlightMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	name               *string
	from               *string
	to                 *string
	depart_date        *time.Time
	depart_time        *time.Time
	status             *flight.Status
	available_slots    *int
	addavailable_slots *int
	return_date        *time.Time
	_type              *flight.Type
	flight_plane       *string
	created_at         *time.Time
	updated_at         *time.Time
	clearedFields      map[string]struct{}
	done               bool
	oldValue           func(context.Context) (*Flight, error)
	predicates         []predicate.Flight
}

var _ ent.Mutation = (*FlightMutation)(nil)

// flightOption allows management of the mutation configuration using functional options.
type flightOption func(*FlightMutation)

// newFlightMutation creates new mutation for the Flight entity.
func newFlightMutation(c config, op Op, opts ...flightOption) *FlightMutation {
	m := &FlightMutation{
		config:        c,
		op:            op,
		typ:           TypeFlight,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFlightID sets the ID field of the mutation.
func withFlightID(id int) flightOption {
	return func(m *FlightMutation) {
		var (
			err   error
			once  sync.Once
			value *Flight
		)
		m.oldValue = func(ctx context.Context) (*Flight, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Flight.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFlight sets the old Flight of the mutation.
func withFlight(node *Flight) flightOption {
	return func(m *FlightMutation) {
		m.oldValue = func(context.Context) (*Flight, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FlightMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FlightMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FlightMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *FlightMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Flight.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *FlightMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FlightMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FlightMutation) ResetName() {
	m.name = nil
}

// SetFrom sets the "from" field.
func (m *FlightMutation) SetFrom(s string) {
	m.from = &s
}

// From returns the value of the "from" field in the mutation.
func (m *FlightMutation) From() (r string, exists bool) {
	v := m.from
	if v == nil {
		return
	}
	return *v, true
}

// OldFrom returns the old "from" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldFrom(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFrom is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFrom requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFrom: %w", err)
	}
	return oldValue.From, nil
}

// ResetFrom resets all changes to the "from" field.
func (m *FlightMutation) ResetFrom() {
	m.from = nil
}

// SetTo sets the "to" field.
func (m *FlightMutation) SetTo(s string) {
	m.to = &s
}

// To returns the value of the "to" field in the mutation.
func (m *FlightMutation) To() (r string, exists bool) {
	v := m.to
	if v == nil {
		return
	}
	return *v, true
}

// OldTo returns the old "to" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldTo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTo: %w", err)
	}
	return oldValue.To, nil
}

// ResetTo resets all changes to the "to" field.
func (m *FlightMutation) ResetTo() {
	m.to = nil
}

// SetDepartDate sets the "depart_date" field.
func (m *FlightMutation) SetDepartDate(t time.Time) {
	m.depart_date = &t
}

// DepartDate returns the value of the "depart_date" field in the mutation.
func (m *FlightMutation) DepartDate() (r time.Time, exists bool) {
	v := m.depart_date
	if v == nil {
		return
	}
	return *v, true
}

// OldDepartDate returns the old "depart_date" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldDepartDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDepartDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDepartDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDepartDate: %w", err)
	}
	return oldValue.DepartDate, nil
}

// ResetDepartDate resets all changes to the "depart_date" field.
func (m *FlightMutation) ResetDepartDate() {
	m.depart_date = nil
}

// SetDepartTime sets the "depart_time" field.
func (m *FlightMutation) SetDepartTime(t time.Time) {
	m.depart_time = &t
}

// DepartTime returns the value of the "depart_time" field in the mutation.
func (m *FlightMutation) DepartTime() (r time.Time, exists bool) {
	v := m.depart_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDepartTime returns the old "depart_time" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldDepartTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDepartTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDepartTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDepartTime: %w", err)
	}
	return oldValue.DepartTime, nil
}

// ResetDepartTime resets all changes to the "depart_time" field.
func (m *FlightMutation) ResetDepartTime() {
	m.depart_time = nil
}

// SetStatus sets the "status" field.
func (m *FlightMutation) SetStatus(f flight.Status) {
	m.status = &f
}

// Status returns the value of the "status" field in the mutation.
func (m *FlightMutation) Status() (r flight.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldStatus(ctx context.Context) (v flight.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *FlightMutation) ResetStatus() {
	m.status = nil
}

// SetAvailableSlots sets the "available_slots" field.
func (m *FlightMutation) SetAvailableSlots(i int) {
	m.available_slots = &i
	m.addavailable_slots = nil
}

// AvailableSlots returns the value of the "available_slots" field in the mutation.
func (m *FlightMutation) AvailableSlots() (r int, exists bool) {
	v := m.available_slots
	if v == nil {
		return
	}
	return *v, true
}

// OldAvailableSlots returns the old "available_slots" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldAvailableSlots(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvailableSlots is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvailableSlots requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvailableSlots: %w", err)
	}
	return oldValue.AvailableSlots, nil
}

// AddAvailableSlots adds i to the "available_slots" field.
func (m *FlightMutation) AddAvailableSlots(i int) {
	if m.addavailable_slots != nil {
		*m.addavailable_slots += i
	} else {
		m.addavailable_slots = &i
	}
}

// AddedAvailableSlots returns the value that was added to the "available_slots" field in this mutation.
func (m *FlightMutation) AddedAvailableSlots() (r int, exists bool) {
	v := m.addavailable_slots
	if v == nil {
		return
	}
	return *v, true
}

// ResetAvailableSlots resets all changes to the "available_slots" field.
func (m *FlightMutation) ResetAvailableSlots() {
	m.available_slots = nil
	m.addavailable_slots = nil
}

// SetReturnDate sets the "return_date" field.
func (m *FlightMutation) SetReturnDate(t time.Time) {
	m.return_date = &t
}

// ReturnDate returns the value of the "return_date" field in the mutation.
func (m *FlightMutation) ReturnDate() (r time.Time, exists bool) {
	v := m.return_date
	if v == nil {
		return
	}
	return *v, true
}

// OldReturnDate returns the old "return_date" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldReturnDate(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReturnDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReturnDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReturnDate: %w", err)
	}
	return oldValue.ReturnDate, nil
}

// ClearReturnDate clears the value of the "return_date" field.
func (m *FlightMutation) ClearReturnDate() {
	m.return_date = nil
	m.clearedFields[flight.FieldReturnDate] = struct{}{}
}

// ReturnDateCleared returns if the "return_date" field was cleared in this mutation.
func (m *FlightMutation) ReturnDateCleared() bool {
	_, ok := m.clearedFields[flight.FieldReturnDate]
	return ok
}

// ResetReturnDate resets all changes to the "return_date" field.
func (m *FlightMutation) ResetReturnDate() {
	m.return_date = nil
	delete(m.clearedFields, flight.FieldReturnDate)
}

// SetType sets the "type" field.
func (m *FlightMutation) SetType(f flight.Type) {
	m._type = &f
}

// GetType returns the value of the "type" field in the mutation.
func (m *FlightMutation) GetType() (r flight.Type, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldType(ctx context.Context) (v flight.Type, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *FlightMutation) ResetType() {
	m._type = nil
}

// SetFlightPlane sets the "flight_plane" field.
func (m *FlightMutation) SetFlightPlane(s string) {
	m.flight_plane = &s
}

// FlightPlane returns the value of the "flight_plane" field in the mutation.
func (m *FlightMutation) FlightPlane() (r string, exists bool) {
	v := m.flight_plane
	if v == nil {
		return
	}
	return *v, true
}

// OldFlightPlane returns the old "flight_plane" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldFlightPlane(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFlightPlane is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFlightPlane requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFlightPlane: %w", err)
	}
	return oldValue.FlightPlane, nil
}

// ResetFlightPlane resets all changes to the "flight_plane" field.
func (m *FlightMutation) ResetFlightPlane() {
	m.flight_plane = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *FlightMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *FlightMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *FlightMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *FlightMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *FlightMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Flight entity.
// If the Flight object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FlightMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *FlightMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the FlightMutation builder.
func (m *FlightMutation) Where(ps ...predicate.Flight) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the FlightMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *FlightMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Flight, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *FlightMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *FlightMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Flight).
func (m *FlightMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FlightMutation) Fields() []string {
	fields := make([]string, 0, 12)
	if m.name != nil {
		fields = append(fields, flight.FieldName)
	}
	if m.from != nil {
		fields = append(fields, flight.FieldFrom)
	}
	if m.to != nil {
		fields = append(fields, flight.FieldTo)
	}
	if m.depart_date != nil {
		fields = append(fields, flight.FieldDepartDate)
	}
	if m.depart_time != nil {
		fields = append(fields, flight.FieldDepartTime)
	}
	if m.status != nil {
		fields = append(fields, flight.FieldStatus)
	}
	if m.available_slots != nil {
		fields = append(fields, flight.FieldAvailableSlots)
	}
	if m.return_date != nil {
		fields = append(fields, flight.FieldReturnDate)
	}
	if m._type != nil {
		fields = append(fields, flight.FieldType)
	}
	if m.flight_plane != nil {
		fields = append(fields, flight.FieldFlightPlane)
	}
	if m.created_at != nil {
		fields = append(fields, flight.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, flight.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FlightMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case flight.FieldName:
		return m.Name()
	case flight.FieldFrom:
		return m.From()
	case flight.FieldTo:
		return m.To()
	case flight.FieldDepartDate:
		return m.DepartDate()
	case flight.FieldDepartTime:
		return m.DepartTime()
	case flight.FieldStatus:
		return m.Status()
	case flight.FieldAvailableSlots:
		return m.AvailableSlots()
	case flight.FieldReturnDate:
		return m.ReturnDate()
	case flight.FieldType:
		return m.GetType()
	case flight.FieldFlightPlane:
		return m.FlightPlane()
	case flight.FieldCreatedAt:
		return m.CreatedAt()
	case flight.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FlightMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case flight.FieldName:
		return m.OldName(ctx)
	case flight.FieldFrom:
		return m.OldFrom(ctx)
	case flight.FieldTo:
		return m.OldTo(ctx)
	case flight.FieldDepartDate:
		return m.OldDepartDate(ctx)
	case flight.FieldDepartTime:
		return m.OldDepartTime(ctx)
	case flight.FieldStatus:
		return m.OldStatus(ctx)
	case flight.FieldAvailableSlots:
		return m.OldAvailableSlots(ctx)
	case flight.FieldReturnDate:
		return m.OldReturnDate(ctx)
	case flight.FieldType:
		return m.OldType(ctx)
	case flight.FieldFlightPlane:
		return m.OldFlightPlane(ctx)
	case flight.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case flight.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Flight field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FlightMutation) SetField(name string, value ent.Value) error {
	switch name {
	case flight.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case flight.FieldFrom:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFrom(v)
		return nil
	case flight.FieldTo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTo(v)
		return nil
	case flight.FieldDepartDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDepartDate(v)
		return nil
	case flight.FieldDepartTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDepartTime(v)
		return nil
	case flight.FieldStatus:
		v, ok := value.(flight.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case flight.FieldAvailableSlots:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvailableSlots(v)
		return nil
	case flight.FieldReturnDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReturnDate(v)
		return nil
	case flight.FieldType:
		v, ok := value.(flight.Type)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case flight.FieldFlightPlane:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFlightPlane(v)
		return nil
	case flight.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case flight.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Flight field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FlightMutation) AddedFields() []string {
	var fields []string
	if m.addavailable_slots != nil {
		fields = append(fields, flight.FieldAvailableSlots)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FlightMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case flight.FieldAvailableSlots:
		return m.AddedAvailableSlots()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FlightMutation) AddField(name string, value ent.Value) error {
	switch name {
	case flight.FieldAvailableSlots:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAvailableSlots(v)
		return nil
	}
	return fmt.Errorf("unknown Flight numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FlightMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(flight.FieldReturnDate) {
		fields = append(fields, flight.FieldReturnDate)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FlightMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FlightMutation) ClearField(name string) error {
	switch name {
	case flight.FieldReturnDate:
		m.ClearReturnDate()
		return nil
	}
	return fmt.Errorf("unknown Flight nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FlightMutation) ResetField(name string) error {
	switch name {
	case flight.FieldName:
		m.ResetName()
		return nil
	case flight.FieldFrom:
		m.ResetFrom()
		return nil
	case flight.FieldTo:
		m.ResetTo()
		return nil
	case flight.FieldDepartDate:
		m.ResetDepartDate()
		return nil
	case flight.FieldDepartTime:
		m.ResetDepartTime()
		return nil
	case flight.FieldStatus:
		m.ResetStatus()
		return nil
	case flight.FieldAvailableSlots:
		m.ResetAvailableSlots()
		return nil
	case flight.FieldReturnDate:
		m.ResetReturnDate()
		return nil
	case flight.FieldType:
		m.ResetType()
		return nil
	case flight.FieldFlightPlane:
		m.ResetFlightPlane()
		return nil
	case flight.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case flight.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Flight field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FlightMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FlightMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FlightMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FlightMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FlightMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FlightMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FlightMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Flight unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FlightMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Flight edge %s", name)
}
