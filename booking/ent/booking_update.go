// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking/ent/booking"
	"booking/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBookingCode sets the "booking_code" field.
func (bu *BookingUpdate) SetBookingCode(s string) *BookingUpdate {
	bu.mutation.SetBookingCode(s)
	return bu
}

// SetBookingDate sets the "booking_date" field.
func (bu *BookingUpdate) SetBookingDate(t time.Time) *BookingUpdate {
	bu.mutation.SetBookingDate(t)
	return bu
}

// SetCancelDate sets the "cancel_date" field.
func (bu *BookingUpdate) SetCancelDate(t time.Time) *BookingUpdate {
	bu.mutation.SetCancelDate(t)
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BookingUpdate) SetCreatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableCreatedAt(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookingUpdate) SetUpdatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableUpdatedAt(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// SetFlightID sets the "flight_id" field.
func (bu *BookingUpdate) SetFlightID(i int) *BookingUpdate {
	bu.mutation.ResetFlightID()
	bu.mutation.SetFlightID(i)
	return bu
}

// AddFlightID adds i to the "flight_id" field.
func (bu *BookingUpdate) AddFlightID(i int) *BookingUpdate {
	bu.mutation.AddFlightID(i)
	return bu
}

// SetCustomerID sets the "customer_id" field.
func (bu *BookingUpdate) SetCustomerID(i int) *BookingUpdate {
	bu.mutation.ResetCustomerID()
	bu.mutation.SetCustomerID(i)
	return bu
}

// AddCustomerID adds i to the "customer_id" field.
func (bu *BookingUpdate) AddCustomerID(i int) *BookingUpdate {
	bu.mutation.AddCustomerID(i)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BookingUpdate) SetStatus(b booking.Status) *BookingUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableStatus(b *booking.Status) *BookingUpdate {
	if b != nil {
		bu.SetStatus(*b)
	}
	return bu
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if v, ok := bu.mutation.BookingCode(); ok {
		if err := booking.BookingCodeValidator(v); err != nil {
			return &ValidationError{Name: "booking_code", err: fmt.Errorf(`ent: validator failed for field "Booking.booking_code": %w`, err)}
		}
	}
	if v, ok := bu.mutation.FlightID(); ok {
		if err := booking.FlightIDValidator(v); err != nil {
			return &ValidationError{Name: "flight_id", err: fmt.Errorf(`ent: validator failed for field "Booking.flight_id": %w`, err)}
		}
	}
	if v, ok := bu.mutation.CustomerID(); ok {
		if err := booking.CustomerIDValidator(v); err != nil {
			return &ValidationError{Name: "customer_id", err: fmt.Errorf(`ent: validator failed for field "Booking.customer_id": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	return nil
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BookingCode(); ok {
		_spec.SetField(booking.FieldBookingCode, field.TypeString, value)
	}
	if value, ok := bu.mutation.BookingDate(); ok {
		_spec.SetField(booking.FieldBookingDate, field.TypeTime, value)
	}
	if value, ok := bu.mutation.CancelDate(); ok {
		_spec.SetField(booking.FieldCancelDate, field.TypeTime, value)
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.FlightID(); ok {
		_spec.SetField(booking.FieldFlightID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedFlightID(); ok {
		_spec.AddField(booking.FieldFlightID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.CustomerID(); ok {
		_spec.SetField(booking.FieldCustomerID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedCustomerID(); ok {
		_spec.AddField(booking.FieldCustomerID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetBookingCode sets the "booking_code" field.
func (buo *BookingUpdateOne) SetBookingCode(s string) *BookingUpdateOne {
	buo.mutation.SetBookingCode(s)
	return buo
}

// SetBookingDate sets the "booking_date" field.
func (buo *BookingUpdateOne) SetBookingDate(t time.Time) *BookingUpdateOne {
	buo.mutation.SetBookingDate(t)
	return buo
}

// SetCancelDate sets the "cancel_date" field.
func (buo *BookingUpdateOne) SetCancelDate(t time.Time) *BookingUpdateOne {
	buo.mutation.SetCancelDate(t)
	return buo
}

// SetCreatedAt sets the "created_at" field.
func (buo *BookingUpdateOne) SetCreatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCreatedAt(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookingUpdateOne) SetUpdatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableUpdatedAt(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// SetFlightID sets the "flight_id" field.
func (buo *BookingUpdateOne) SetFlightID(i int) *BookingUpdateOne {
	buo.mutation.ResetFlightID()
	buo.mutation.SetFlightID(i)
	return buo
}

// AddFlightID adds i to the "flight_id" field.
func (buo *BookingUpdateOne) AddFlightID(i int) *BookingUpdateOne {
	buo.mutation.AddFlightID(i)
	return buo
}

// SetCustomerID sets the "customer_id" field.
func (buo *BookingUpdateOne) SetCustomerID(i int) *BookingUpdateOne {
	buo.mutation.ResetCustomerID()
	buo.mutation.SetCustomerID(i)
	return buo
}

// AddCustomerID adds i to the "customer_id" field.
func (buo *BookingUpdateOne) AddCustomerID(i int) *BookingUpdateOne {
	buo.mutation.AddCustomerID(i)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BookingUpdateOne) SetStatus(b booking.Status) *BookingUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableStatus(b *booking.Status) *BookingUpdateOne {
	if b != nil {
		buo.SetStatus(*b)
	}
	return buo
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (buo *BookingUpdateOne) Where(ps ...predicate.Booking) *BookingUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if v, ok := buo.mutation.BookingCode(); ok {
		if err := booking.BookingCodeValidator(v); err != nil {
			return &ValidationError{Name: "booking_code", err: fmt.Errorf(`ent: validator failed for field "Booking.booking_code": %w`, err)}
		}
	}
	if v, ok := buo.mutation.FlightID(); ok {
		if err := booking.FlightIDValidator(v); err != nil {
			return &ValidationError{Name: "flight_id", err: fmt.Errorf(`ent: validator failed for field "Booking.flight_id": %w`, err)}
		}
	}
	if v, ok := buo.mutation.CustomerID(); ok {
		if err := booking.CustomerIDValidator(v); err != nil {
			return &ValidationError{Name: "customer_id", err: fmt.Errorf(`ent: validator failed for field "Booking.customer_id": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	return nil
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.BookingCode(); ok {
		_spec.SetField(booking.FieldBookingCode, field.TypeString, value)
	}
	if value, ok := buo.mutation.BookingDate(); ok {
		_spec.SetField(booking.FieldBookingDate, field.TypeTime, value)
	}
	if value, ok := buo.mutation.CancelDate(); ok {
		_spec.SetField(booking.FieldCancelDate, field.TypeTime, value)
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.FlightID(); ok {
		_spec.SetField(booking.FieldFlightID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedFlightID(); ok {
		_spec.AddField(booking.FieldFlightID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.CustomerID(); ok {
		_spec.SetField(booking.FieldCustomerID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedCustomerID(); ok {
		_spec.AddField(booking.FieldCustomerID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
