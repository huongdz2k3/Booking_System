package main

import (
	"customer/config"
	"customer/ent"
	"customer/internal/logger"
	"customer/internal/utils"
	"customer/resolver"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

const configFile = "config.yaml"

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}

// initConfig initializes the config.
func initConfig() *config.Configurations {
	viper.SetConfigType("yaml")

	// Expand environment variables inside the config file
	b, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	expand := os.ExpandEnv(string(b))
	configReader := strings.NewReader(expand)

	viper.AutomaticEnv()

	if err := viper.ReadConfig(configReader); err != nil {
		fmt.Printf("read config has failed with error: %v\n", err)
		os.Exit(1)
	}

	configs := config.Configurations{}
	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}
	return &configs
}

func main() {
	client := &ent.Client{}
	utils.LoadEnv()
	initLogger()
	// Create a Gin router instance
	r := gin.Default()

	get := utils.GetEnvWithKey
	PORT := get("PORT")

	logger.NewLogger().Info("Listening on port: " + PORT)

	srv := handler.NewDefaultServer(resolver.NewSchema(client))

	// Define the GraphQL endpoint
	r.POST("/query", gin.WrapH(srv))

	// Define the GraphQL Playground endpoint
	r.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":" + PORT); err != nil {
		logger.NewLogger().Error("Get error from run server", zap.Error(err))
	}

}
