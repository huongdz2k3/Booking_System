package api

import (
	"customer/ent"
	"customer/internal/utils"
	"customer/middleware"
	"customer/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServerCmd(logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {
			client := &ent.Client{}
			utils.LoadEnv()
			// Create a Gin router instance
			r := gin.Default()

			r.Use(middleware.AuthMiddleware(), middleware.CorsMiddleware(), middleware.RequestCtxMiddleware())

			get := utils.GetEnvWithKey
			PORT := get("PORT")

			logger.Info("Listening on port: " + PORT)

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
				logger.Error("Get error from run server", zap.Error(err))
			}

		},
	}
}
