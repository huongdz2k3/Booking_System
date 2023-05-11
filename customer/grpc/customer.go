package grpc

import (
	context "context"
	ent "customer-service/ent"
	"customer-service/ent/customer"
	"customer-service/internal/utils"
	"customer-service/pb"
	"customer-service/tool"
	"customer-service/validation"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// CustomerService implements CustomerServiceServer
type CustomerService struct {
	client *ent.Client
	pb.UnimplementedCustomerServiceServer
}

// NewCustomerService returns a new CustomerService
func NewCustomerService(client *ent.Client) *CustomerService {
	return &CustomerService{
		client: client,
	}
}

// toProtoCustomer transforms the ent type to the pb type
func toProtoCustomer(e *ent.Customer) (*pb.Customer, error) {
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
	if &e.MembershipNumber != nil {
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

// Register implements CustomerServiceServer.Register
func (svc *CustomerService) Register(ctx context.Context, req *pb.RegisterInput) (*pb.Customer, error) {
	if len(req.GetPassword()) < 8 {
		return nil, utils.WrapBadRequestError(ctx, "Password must be at least 8 characters")
	}
	req.Password = tool.HashPassword(req.GetPassword())

	// check email
	if validation.IsEmail(req.GetEmail()) == false {
		return nil, utils.WrapBadRequestError(ctx, "Invalid Email")
	}
	checkExist, _ := svc.GetCustomerByEmail(ctx, &pb.GetCustomerByEmailInput{Email: req.GetEmail()})

	if checkExist != nil {
		return nil, utils.WrapBadRequestError(ctx, "Email already exists")
	}

	// check license id
	checkLicenseExist, _ := svc.client.Customer.Query().Where(customer.LicenseID(req.GetLicenseId())).Only(ctx)
	if checkLicenseExist != nil {
		return nil, utils.WrapBadRequestError(ctx, "License ID already exists")
	}

	u, err := svc.client.Customer.Create().SetName(req.GetName()).SetEmail(req.GetEmail()).SetPhoneNumber(req.GetPhoneNumber()).SetLicenseID(req.GetLicenseId()).SetAddress(req.GetAddress()).SetMembershipNumber(int(req.GetMembershipNumber())).SetPassword(req.GetPassword()).SetRole("SUBSCRIBER").Save(ctx)
	if err != nil {
		return nil, utils.WrapBadRequestError(ctx, "Invalid input")
	}
	return toProtoCustomer(u)
}

// GetCustomerByEmail implements CustomerServiceServer.GetCustomerByEmail
func (svc *CustomerService) GetCustomerByEmail(ctx context.Context, req *pb.GetCustomerByEmailInput) (*pb.Customer, error) {
	u, err := svc.client.Customer.Query().Where(customer.Email(req.GetEmail())).Only(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}

	return toProtoCustomer(u)

}

// Login implements CustomerServiceServer.Login
func (svc *CustomerService) Login(ctx context.Context, req *pb.LoginInput) (*pb.Customer, error) {
	// check email and pass word
	cus, _ := svc.GetCustomerByEmail(ctx, &pb.GetCustomerByEmailInput{Email: req.GetEmail()})
	err := tool.ComparePassword(cus.GetPassword(), req.GetPassword())
	if (cus == nil) || ((err) != nil) {
		return nil, utils.WrapBadRequestError(ctx, "Wrong password or Email")
	}

	return cus, nil
}

// Update implements CustomerServiceServer.Update
func (svc *CustomerService) Update(ctx context.Context, req *pb.UpdateCustomerInput) (*pb.Customer, error) {
	u, err := svc.client.Customer.UpdateOneID(int(req.GetId())).SetName(req.GetName()).SetEmail(req.GetEmail()).SetPhoneNumber(req.GetPhoneNumber()).SetAddress(req.GetAddress()).SetUpdateAt(time.Now()).Save(ctx)
	if *req.MembershipNumber != 0 {
		membershipNumber := int(req.GetMembershipNumber())
		u.MembershipNumber = &membershipNumber
		u.Update().Save(ctx)
	}
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}

	return toProtoCustomer(u)
}

// ChangePassword implements CustomerServiceServer.ChangePassword
func (svc *CustomerService) ChangePassword(ctx context.Context, req *pb.ChangePasswordInput) (*pb.Customer, error) {
	u, err := svc.client.Customer.Query().Where(customer.ID(int(req.GetId()))).Only(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}
	err = tool.ComparePassword(u.Password, req.GetOldPassword())
	if err != nil {
		return nil, utils.WrapBadRequestError(ctx, "Wrong password")
	}
	if req.NewPassword != req.ConfirmPassword {
		return nil, utils.WrapBadRequestError(ctx, "Password and confirm password not match")
	}
	if len(req.NewPassword) < 8 {
		return nil, utils.WrapBadRequestError(ctx, "Password must be at least 8 characters")
	}
	u, err = svc.client.Customer.UpdateOneID(int(req.GetId())).SetPassword(tool.HashPassword(req.GetNewPassword())).SetUpdateAt(time.Now()).Save(ctx)

	return toProtoCustomer(u)
}

// UpdateRole implements CustomerServiceServer.UpdateRole
func (svc *CustomerService) UpdateRole(ctx context.Context, req *pb.UpdateRoleInput) (*pb.Customer, error) {
	u, err := svc.client.Customer.UpdateOneID(int(req.GetId())).SetRole(req.GetRole()).SetUpdateAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}

	return toProtoCustomer(u)
}
