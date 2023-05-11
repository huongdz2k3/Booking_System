package main

import (
	"flight-service/cmd/api"
	"flight-service/internal/logger"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

// run server with CLI
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server CLI",
	Long:  "run server with CLI",
}

// init initializes the env and logger.
func init() {
	logger := initLogger()

	apiCmd := api.NewServerCmd(logger)
	rootCmd.AddCommand(apiCmd)
}

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("run command has failed with error: %v\n", err)
		os.Exit(1)
	}
}
