package grpc

import (
	"customer-service/ent"
	"customer-service/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// toProtoCustomer transforms the ent type to the pb type
func ToProtoCustomer(e *ent.Customer) (*pb.Customer, error) {
	v := &pb.Customer{}
	address := e.Address
	v.Address = address
	create_at := timestamppb.New(e.CreateAt)
	v.CreateAt = create_at
	email := e.Email
	v.Email = email
	id := int64(e.ID)
	v.Id = id
	is_active := e.IsActive
	v.IsActive = is_active
	license_id := e.LicenseID
	v.LicenseId = license_id
	if e.MembershipNumber != nil {
		membership_number := int64(*e.MembershipNumber)
		v.MembershipNumber = &membership_number
	}
	name := e.Name
	v.Name = name
	password := e.Password
	v.Password = password
	phone_number := e.PhoneNumber
	v.PhoneNumber = phone_number
	role := e.Role
	v.Role = role
	update_at := timestamppb.New(e.UpdateAt)
	v.UpdateAt = update_at
	return v, nil
}
