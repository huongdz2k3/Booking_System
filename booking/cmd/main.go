package main

import (
	"booking/ent"
	"booking/internal/logger"
	"booking/internal/utils"
	"context"
	"go.uber.org/zap"
)

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}
func main() {
	initLogger()
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	client, err := ent.Open("postgres", get("POSTGRES_CONNECTION_STRING"))

	if err != nil {
		logger.NewLogger().Fatal("failed opening connection to postgres", zap.Error(err))

	}

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.NewLogger().Fatal("failed creating schema resources", zap.Error(err))
	}
}
