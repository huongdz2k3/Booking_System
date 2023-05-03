package customer

import (
	"context"
	"customer/ent"
	"customer/ent/customer"
	"customer/internal/logger"
	"customer/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createGRPCClient() (CustomerServiceClient, error) {
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	// Open a connection to the server.
	conn, err := grpc.Dial(":"+get("CUSTOMER_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.NewLogger().Fatal("failed connecting to server", zap.Error(err))
	}
	return NewCustomerServiceClient(conn), nil
}

func GetCustomerByEmail(input *GetCustomerByEmailInput) (*Customer, error) {
	// create a gRPC client instance
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	// connect gRPC server
	resp, err := client.GetCustomerByEmail(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Register(input *ent.RegisterInput) (*Customer, error) {
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	registerInput, _ := ConvertRegisterInput(input)
	// connect gRPC server
	resp, err := client.Register(context.Background(), registerInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Login(input *ent.LoginInput) (*Customer, error) {
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	// connect gRPC server
	loginInput, _ := ConvertLoginInput(input)
	resp, err := client.Login(context.Background(), loginInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Update(input *ent.UpdateCustomerInput, id int) (*Customer, error) {
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	updateInput := ConvertUpdateInput(input, int64(id))
	// connect gRPC server
	resp, err := client.Update(context.Background(), updateInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ChangePassword(input *ent.ChangePasswordInput, id int) (string, error) {
	client, err := createGRPCClient()
	if err != nil {
		return "", err
	}
	changePasswordInput := ConvertChangePasswordInput(input, int64(id))
	// connect gRPC server
	_, err = client.ChangePassword(context.Background(), changePasswordInput)
	if err != nil {
		return "", err
	}
	return "Password changes successfully", nil
}

func UpdateRole(input customer.Role, id int) (*Customer, error) {
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	updateRoleInput := &UpdateRoleInput{
		Id:   int64(id),
		Role: input.String(),
	}
	// connect gRPC server
	resp, err := client.UpdateRole(context.Background(), updateRoleInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
