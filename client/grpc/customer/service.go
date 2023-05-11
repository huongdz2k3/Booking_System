package customer

import (
	"context"
	"customer/ent"
	"customer/ent/customer"
	grpc_client "customer/internal/grpc"
	"customer/internal/logger"
	"customer/pb"
	"go.uber.org/zap"
	"sync"
)

type grpcClient struct {
	client pb.CustomerServiceClient
	err    error
}

var (
	once     sync.Once
	instance *grpcClient
)

func GetCustomerGRPCClient() *grpcClient {
	once.Do(func() {
		client, err := grpc_client.NewGrpcClient(grpc_client.CustomerServiceClient, grpc_client.CustomerServiceHost)
		if err != nil {
			logger.NewLogger().Fatal("failed connecting to server", zap.Error(err))
		}
		instance = &grpcClient{
			client: client.(pb.CustomerServiceClient),
			err:    nil,
		}

	})
	return instance
}

func GetCustomerByEmail(input *pb.GetCustomerByEmailInput) (*pb.Customer, error) {
	// create a gRPC client instance
	client := GetCustomerGRPCClient()

	// connect gRPC server
	resp, err := client.client.GetCustomerByEmail(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Register(input *ent.RegisterInput) (*pb.Customer, error) {

	registerInput, _ := ConvertRegisterInput(input)
	// connect gRPC server
	resp, err := GetCustomerGRPCClient().client.Register(context.Background(), registerInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Login(input *ent.LoginInput) (*pb.Customer, error) {

	// connect gRPC server
	loginInput, _ := ConvertLoginInput(input)
	resp, err := GetCustomerGRPCClient().client.Login(context.Background(), loginInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Update(input *ent.UpdateCustomerInput, id int) (*pb.Customer, error) {

	updateInput := ConvertUpdateInput(input, int64(id))
	// connect gRPC server
	resp, err := GetCustomerGRPCClient().client.Update(context.Background(), updateInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ChangePassword(input *ent.ChangePasswordInput, id int) (string, error) {

	changePasswordInput := ConvertChangePasswordInput(input, int64(id))
	// connect gRPC server
	_, err := GetCustomerGRPCClient().client.ChangePassword(context.Background(), changePasswordInput)
	if err != nil {
		return "", err
	}
	return "Password changes successfully", nil
}

func UpdateRole(input customer.Role, id int) (*pb.Customer, error) {
	updateRoleInput := &pb.UpdateRoleInput{
		Id:   int64(id),
		Role: input.String(),
	}
	// connect gRPC server
	resp, err := GetCustomerGRPCClient().client.UpdateRole(context.Background(), updateRoleInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
