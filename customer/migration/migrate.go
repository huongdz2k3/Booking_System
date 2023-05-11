package migration

import (
	"context"
	"customer-service/ent"
	"customer-service/ent/customer"
	"customer-service/internal/logger"
	"customer-service/tool"
	"go.uber.org/zap"
)

func CreateAdmin(client *ent.Client) {
	//create admin
	admin := &ent.Customer{
		Email:       "admin@localhost",
		Password:    "admin123",
		Role:        "ADMIN",
		Name:        "admin",
		PhoneNumber: "0985830569",
		Address:     "HN",
		LicenseID:   "1231231231",
	}
	user, _ := client.Customer.Query().Where(customer.Email(admin.Email)).First(context.Background())
	if user != nil {
		return
	}
	admin.Password = tool.HashPassword(admin.Password)
	_, err := client.Customer.Create().SetEmail(admin.Email).SetPassword(admin.Password).SetRole(admin.Role).SetName(admin.Name).SetPhoneNumber(admin.PhoneNumber).SetAddress(admin.Address).SetLicenseID(admin.LicenseID).Save(context.Background())
	if err != nil {
		logger.NewLogger().Fatal("failed creating admin", zap.Error(err))
	}
}
