// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer-service/ent/customer"
	"customer-service/ent/predicate"
	"errors"
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
	TypeCustomer = "Customer"
)

// CustomerMutation represents an operation that mutates the Customer nodes in the graph.
type CustomerMutation struct {
	config
	op                   Op
	typ                  string
	id                   *int
	name                 *string
	phone_number         *string
	email                *string
	license_id           *string
	address              *string
	membership_number    *int
	addmembership_number *int
	is_active            *bool
	password             *string
	create_at            *time.Time
	update_at            *time.Time
	role                 *string
	clearedFields        map[string]struct{}
	done                 bool
	oldValue             func(context.Context) (*Customer, error)
	predicates           []predicate.Customer
}

var _ ent.Mutation = (*CustomerMutation)(nil)

// customerOption allows management of the mutation configuration using functional options.
type customerOption func(*CustomerMutation)

// newCustomerMutation creates new mutation for the Customer entity.
func newCustomerMutation(c config, op Op, opts ...customerOption) *CustomerMutation {
	m := &CustomerMutation{
		config:        c,
		op:            op,
		typ:           TypeCustomer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCustomerID sets the ID field of the mutation.
func withCustomerID(id int) customerOption {
	return func(m *CustomerMutation) {
		var (
			err   error
			once  sync.Once
			value *Customer
		)
		m.oldValue = func(ctx context.Context) (*Customer, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Customer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCustomer sets the old Customer of the mutation.
func withCustomer(node *Customer) customerOption {
	return func(m *CustomerMutation) {
		m.oldValue = func(context.Context) (*Customer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CustomerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CustomerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CustomerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CustomerMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Customer.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *CustomerMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CustomerMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *CustomerMutation) ResetName() {
	m.name = nil
}

// SetPhoneNumber sets the "phone_number" field.
func (m *CustomerMutation) SetPhoneNumber(s string) {
	m.phone_number = &s
}

// PhoneNumber returns the value of the "phone_number" field in the mutation.
func (m *CustomerMutation) PhoneNumber() (r string, exists bool) {
	v := m.phone_number
	if v == nil {
		return
	}
	return *v, true
}

// OldPhoneNumber returns the old "phone_number" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldPhoneNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhoneNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhoneNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhoneNumber: %w", err)
	}
	return oldValue.PhoneNumber, nil
}

// ResetPhoneNumber resets all changes to the "phone_number" field.
func (m *CustomerMutation) ResetPhoneNumber() {
	m.phone_number = nil
}

// SetEmail sets the "email" field.
func (m *CustomerMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *CustomerMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *CustomerMutation) ResetEmail() {
	m.email = nil
}

// SetLicenseID sets the "license_id" field.
func (m *CustomerMutation) SetLicenseID(s string) {
	m.license_id = &s
}

// LicenseID returns the value of the "license_id" field in the mutation.
func (m *CustomerMutation) LicenseID() (r string, exists bool) {
	v := m.license_id
	if v == nil {
		return
	}
	return *v, true
}

// OldLicenseID returns the old "license_id" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldLicenseID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLicenseID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLicenseID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLicenseID: %w", err)
	}
	return oldValue.LicenseID, nil
}

// ResetLicenseID resets all changes to the "license_id" field.
func (m *CustomerMutation) ResetLicenseID() {
	m.license_id = nil
}

// SetAddress sets the "address" field.
func (m *CustomerMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *CustomerMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ResetAddress resets all changes to the "address" field.
func (m *CustomerMutation) ResetAddress() {
	m.address = nil
}

// SetMembershipNumber sets the "membership_number" field.
func (m *CustomerMutation) SetMembershipNumber(i int) {
	m.membership_number = &i
	m.addmembership_number = nil
}

// MembershipNumber returns the value of the "membership_number" field in the mutation.
func (m *CustomerMutation) MembershipNumber() (r int, exists bool) {
	v := m.membership_number
	if v == nil {
		return
	}
	return *v, true
}

// OldMembershipNumber returns the old "membership_number" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldMembershipNumber(ctx context.Context) (v *int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMembershipNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMembershipNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMembershipNumber: %w", err)
	}
	return oldValue.MembershipNumber, nil
}

// AddMembershipNumber adds i to the "membership_number" field.
func (m *CustomerMutation) AddMembershipNumber(i int) {
	if m.addmembership_number != nil {
		*m.addmembership_number += i
	} else {
		m.addmembership_number = &i
	}
}

// AddedMembershipNumber returns the value that was added to the "membership_number" field in this mutation.
func (m *CustomerMutation) AddedMembershipNumber() (r int, exists bool) {
	v := m.addmembership_number
	if v == nil {
		return
	}
	return *v, true
}

// ClearMembershipNumber clears the value of the "membership_number" field.
func (m *CustomerMutation) ClearMembershipNumber() {
	m.membership_number = nil
	m.addmembership_number = nil
	m.clearedFields[customer.FieldMembershipNumber] = struct{}{}
}

// MembershipNumberCleared returns if the "membership_number" field was cleared in this mutation.
func (m *CustomerMutation) MembershipNumberCleared() bool {
	_, ok := m.clearedFields[customer.FieldMembershipNumber]
	return ok
}

// ResetMembershipNumber resets all changes to the "membership_number" field.
func (m *CustomerMutation) ResetMembershipNumber() {
	m.membership_number = nil
	m.addmembership_number = nil
	delete(m.clearedFields, customer.FieldMembershipNumber)
}

// SetIsActive sets the "is_active" field.
func (m *CustomerMutation) SetIsActive(b bool) {
	m.is_active = &b
}

// IsActive returns the value of the "is_active" field in the mutation.
func (m *CustomerMutation) IsActive() (r bool, exists bool) {
	v := m.is_active
	if v == nil {
		return
	}
	return *v, true
}

// OldIsActive returns the old "is_active" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldIsActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsActive: %w", err)
	}
	return oldValue.IsActive, nil
}

// ResetIsActive resets all changes to the "is_active" field.
func (m *CustomerMutation) ResetIsActive() {
	m.is_active = nil
}

// SetPassword sets the "password" field.
func (m *CustomerMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *CustomerMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *CustomerMutation) ResetPassword() {
	m.password = nil
}

// SetCreateAt sets the "create_at" field.
func (m *CustomerMutation) SetCreateAt(t time.Time) {
	m.create_at = &t
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *CustomerMutation) CreateAt() (r time.Time, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldCreateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *CustomerMutation) ResetCreateAt() {
	m.create_at = nil
}

// SetUpdateAt sets the "update_at" field.
func (m *CustomerMutation) SetUpdateAt(t time.Time) {
	m.update_at = &t
}

// UpdateAt returns the value of the "update_at" field in the mutation.
func (m *CustomerMutation) UpdateAt() (r time.Time, exists bool) {
	v := m.update_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateAt returns the old "update_at" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldUpdateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateAt: %w", err)
	}
	return oldValue.UpdateAt, nil
}

// ResetUpdateAt resets all changes to the "update_at" field.
func (m *CustomerMutation) ResetUpdateAt() {
	m.update_at = nil
}

// SetRole sets the "role" field.
func (m *CustomerMutation) SetRole(s string) {
	m.role = &s
}

// Role returns the value of the "role" field in the mutation.
func (m *CustomerMutation) Role() (r string, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldRole(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// ResetRole resets all changes to the "role" field.
func (m *CustomerMutation) ResetRole() {
	m.role = nil
}

// Where appends a list predicates to the CustomerMutation builder.
func (m *CustomerMutation) Where(ps ...predicate.Customer) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CustomerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CustomerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Customer, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CustomerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CustomerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Customer).
func (m *CustomerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CustomerMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.name != nil {
		fields = append(fields, customer.FieldName)
	}
	if m.phone_number != nil {
		fields = append(fields, customer.FieldPhoneNumber)
	}
	if m.email != nil {
		fields = append(fields, customer.FieldEmail)
	}
	if m.license_id != nil {
		fields = append(fields, customer.FieldLicenseID)
	}
	if m.address != nil {
		fields = append(fields, customer.FieldAddress)
	}
	if m.membership_number != nil {
		fields = append(fields, customer.FieldMembershipNumber)
	}
	if m.is_active != nil {
		fields = append(fields, customer.FieldIsActive)
	}
	if m.password != nil {
		fields = append(fields, customer.FieldPassword)
	}
	if m.create_at != nil {
		fields = append(fields, customer.FieldCreateAt)
	}
	if m.update_at != nil {
		fields = append(fields, customer.FieldUpdateAt)
	}
	if m.role != nil {
		fields = append(fields, customer.FieldRole)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CustomerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case customer.FieldName:
		return m.Name()
	case customer.FieldPhoneNumber:
		return m.PhoneNumber()
	case customer.FieldEmail:
		return m.Email()
	case customer.FieldLicenseID:
		return m.LicenseID()
	case customer.FieldAddress:
		return m.Address()
	case customer.FieldMembershipNumber:
		return m.MembershipNumber()
	case customer.FieldIsActive:
		return m.IsActive()
	case customer.FieldPassword:
		return m.Password()
	case customer.FieldCreateAt:
		return m.CreateAt()
	case customer.FieldUpdateAt:
		return m.UpdateAt()
	case customer.FieldRole:
		return m.Role()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CustomerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case customer.FieldName:
		return m.OldName(ctx)
	case customer.FieldPhoneNumber:
		return m.OldPhoneNumber(ctx)
	case customer.FieldEmail:
		return m.OldEmail(ctx)
	case customer.FieldLicenseID:
		return m.OldLicenseID(ctx)
	case customer.FieldAddress:
		return m.OldAddress(ctx)
	case customer.FieldMembershipNumber:
		return m.OldMembershipNumber(ctx)
	case customer.FieldIsActive:
		return m.OldIsActive(ctx)
	case customer.FieldPassword:
		return m.OldPassword(ctx)
	case customer.FieldCreateAt:
		return m.OldCreateAt(ctx)
	case customer.FieldUpdateAt:
		return m.OldUpdateAt(ctx)
	case customer.FieldRole:
		return m.OldRole(ctx)
	}
	return nil, fmt.Errorf("unknown Customer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CustomerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case customer.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case customer.FieldPhoneNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhoneNumber(v)
		return nil
	case customer.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case customer.FieldLicenseID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLicenseID(v)
		return nil
	case customer.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case customer.FieldMembershipNumber:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMembershipNumber(v)
		return nil
	case customer.FieldIsActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsActive(v)
		return nil
	case customer.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case customer.FieldCreateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	case customer.FieldUpdateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateAt(v)
		return nil
	case customer.FieldRole:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	}
	return fmt.Errorf("unknown Customer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CustomerMutation) AddedFields() []string {
	var fields []string
	if m.addmembership_number != nil {
		fields = append(fields, customer.FieldMembershipNumber)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CustomerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case customer.FieldMembershipNumber:
		return m.AddedMembershipNumber()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CustomerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case customer.FieldMembershipNumber:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMembershipNumber(v)
		return nil
	}
	return fmt.Errorf("unknown Customer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CustomerMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(customer.FieldMembershipNumber) {
		fields = append(fields, customer.FieldMembershipNumber)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CustomerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CustomerMutation) ClearField(name string) error {
	switch name {
	case customer.FieldMembershipNumber:
		m.ClearMembershipNumber()
		return nil
	}
	return fmt.Errorf("unknown Customer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CustomerMutation) ResetField(name string) error {
	switch name {
	case customer.FieldName:
		m.ResetName()
		return nil
	case customer.FieldPhoneNumber:
		m.ResetPhoneNumber()
		return nil
	case customer.FieldEmail:
		m.ResetEmail()
		return nil
	case customer.FieldLicenseID:
		m.ResetLicenseID()
		return nil
	case customer.FieldAddress:
		m.ResetAddress()
		return nil
	case customer.FieldMembershipNumber:
		m.ResetMembershipNumber()
		return nil
	case customer.FieldIsActive:
		m.ResetIsActive()
		return nil
	case customer.FieldPassword:
		m.ResetPassword()
		return nil
	case customer.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	case customer.FieldUpdateAt:
		m.ResetUpdateAt()
		return nil
	case customer.FieldRole:
		m.ResetRole()
		return nil
	}
	return fmt.Errorf("unknown Customer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CustomerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CustomerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CustomerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CustomerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CustomerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CustomerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CustomerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Customer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CustomerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Customer edge %s", name)
}
