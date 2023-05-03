package customer

import (
	"customer/ent"
	customer2 "customer/ent/customer"
	"github.com/golang/protobuf/ptypes"
)

func ConvertRegisterInput(input *ent.RegisterInput) (*RegisterInput, error) {
	var number int64
	if input.MembershipNumber != nil {
		number = int64(*input.MembershipNumber)
	}

	return &RegisterInput{
		Email:            input.Email,
		Address:          input.Address,
		LicenseId:        input.LicenseID,
		Name:             input.Name,
		Password:         input.Password,
		PhoneNumber:      input.PhoneNumber,
		MembershipNumber: number,
	}, nil
}

func ConvertLoginInput(input *ent.LoginInput) (*LoginInput, error) {
	return &LoginInput{
		Email:    input.Email,
		Password: input.Password,
	}, nil
}

func ConvertUpdateInput(input *ent.UpdateCustomerInput, id int64) *UpdateCustomerInput {
	return &UpdateCustomerInput{
		Id:               id,
		Email:            *input.Email,
		Address:          *input.Address,
		LicenseId:        *input.LicenseID,
		Name:             *input.Name,
		PhoneNumber:      *input.PhoneNumber,
		MembershipNumber: int64(*input.MembershipNumber),
	}
}

func ConvertChangePasswordInput(input *ent.ChangePasswordInput, id int64) *ChangePasswordInput {
	return &ChangePasswordInput{
		Id:              id,
		OldPassword:     input.OldPassword,
		NewPassword:     input.NewPassword,
		ConfirmPassword: input.ConfirmPassword,
	}
}

func FromProtoCustomer(p *Customer) (*ent.Customer, error) {
	id := int(p.Id)
	membershipNumber := int(p.MembershipNumber)
	createAt, err := ptypes.Timestamp(p.CreateAt)
	if err != nil {
		return nil, err
	}
	updateAt, err := ptypes.Timestamp(p.UpdateAt)
	if err != nil {
		return nil, err
	}
	var role customer2.Role
	if p.Role == "ADMIN" {
		role = customer2.RoleADMIN
	}
	if p.Role == "SUBSCRIBER" {
		role = customer2.RoleSUBSCRIBER
	}

	e := &ent.Customer{
		ID:               id,
		Name:             p.Name,
		Email:            p.Email,
		Password:         p.Password,
		PhoneNumber:      p.PhoneNumber,
		Address:          p.Address,
		Role:             role,
		IsActive:         p.IsActive,
		LicenseID:        p.LicenseId,
		MembershipNumber: &membershipNumber,
		CreateAt:         createAt,
		UpdateAt:         updateAt,
	}

	return e, nil
}
