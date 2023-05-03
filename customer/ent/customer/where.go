// Code generated by ent, DO NOT EDIT.

package customer

import (
	"customer-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPhoneNumber, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldEmail, v))
}

// LicenseID applies equality check predicate on the "license_id" field. It's identical to LicenseIDEQ.
func LicenseID(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldLicenseID, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddress, v))
}

// MembershipNumber applies equality check predicate on the "membership_number" field. It's identical to MembershipNumberEQ.
func MembershipNumber(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldMembershipNumber, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIsActive, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPassword, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdateAt, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldRole, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldName, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldEmail, v))
}

// LicenseIDEQ applies the EQ predicate on the "license_id" field.
func LicenseIDEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldLicenseID, v))
}

// LicenseIDNEQ applies the NEQ predicate on the "license_id" field.
func LicenseIDNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldLicenseID, v))
}

// LicenseIDIn applies the In predicate on the "license_id" field.
func LicenseIDIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldLicenseID, vs...))
}

// LicenseIDNotIn applies the NotIn predicate on the "license_id" field.
func LicenseIDNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldLicenseID, vs...))
}

// LicenseIDGT applies the GT predicate on the "license_id" field.
func LicenseIDGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldLicenseID, v))
}

// LicenseIDGTE applies the GTE predicate on the "license_id" field.
func LicenseIDGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldLicenseID, v))
}

// LicenseIDLT applies the LT predicate on the "license_id" field.
func LicenseIDLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldLicenseID, v))
}

// LicenseIDLTE applies the LTE predicate on the "license_id" field.
func LicenseIDLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldLicenseID, v))
}

// LicenseIDContains applies the Contains predicate on the "license_id" field.
func LicenseIDContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldLicenseID, v))
}

// LicenseIDHasPrefix applies the HasPrefix predicate on the "license_id" field.
func LicenseIDHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldLicenseID, v))
}

// LicenseIDHasSuffix applies the HasSuffix predicate on the "license_id" field.
func LicenseIDHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldLicenseID, v))
}

// LicenseIDEqualFold applies the EqualFold predicate on the "license_id" field.
func LicenseIDEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldLicenseID, v))
}

// LicenseIDContainsFold applies the ContainsFold predicate on the "license_id" field.
func LicenseIDContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldLicenseID, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAddress, v))
}

// MembershipNumberEQ applies the EQ predicate on the "membership_number" field.
func MembershipNumberEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldMembershipNumber, v))
}

// MembershipNumberNEQ applies the NEQ predicate on the "membership_number" field.
func MembershipNumberNEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldMembershipNumber, v))
}

// MembershipNumberIn applies the In predicate on the "membership_number" field.
func MembershipNumberIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldMembershipNumber, vs...))
}

// MembershipNumberNotIn applies the NotIn predicate on the "membership_number" field.
func MembershipNumberNotIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldMembershipNumber, vs...))
}

// MembershipNumberGT applies the GT predicate on the "membership_number" field.
func MembershipNumberGT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldMembershipNumber, v))
}

// MembershipNumberGTE applies the GTE predicate on the "membership_number" field.
func MembershipNumberGTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldMembershipNumber, v))
}

// MembershipNumberLT applies the LT predicate on the "membership_number" field.
func MembershipNumberLT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldMembershipNumber, v))
}

// MembershipNumberLTE applies the LTE predicate on the "membership_number" field.
func MembershipNumberLTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldMembershipNumber, v))
}

// MembershipNumberIsNil applies the IsNil predicate on the "membership_number" field.
func MembershipNumberIsNil() predicate.Customer {
	return predicate.Customer(sql.FieldIsNull(FieldMembershipNumber))
}

// MembershipNumberNotNil applies the NotNil predicate on the "membership_number" field.
func MembershipNumberNotNil() predicate.Customer {
	return predicate.Customer(sql.FieldNotNull(FieldMembershipNumber))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldIsActive, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldPassword, v))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldUpdateAt, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldRole, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
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
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		p(s.Not())
	})
}
