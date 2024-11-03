package routers

import (
	"genggam-makna-api/config"
	"genggam-makna-api/handlers"
	"genggam-makna-api/middleware"
	"genggam-makna-api/repositories"
	"genggam-makna-api/services"

	"github.com/gin-gonic/gin"
)

func CompRouter(api *gin.RouterGroup) {
	api.Use(middleware.ClientTracker(config.InitDB()))

	compRepository := repositories.NewComponentRepository(config.InitDB())
	compService := services.NewService(compRepository)
	compHandler := handlers.NewCompHandlers(compService)

	api.GET("/ping", compHandler.Ping)

	authRoute := api.Group("/user")
	{
		authRoute.POST("/register", compHandler.RegisterUserCredential)
	}

	authRoute.Use(middleware.AuthMiddleware())
	{
		authRoute.GET("/auth-test", compHandler.AuthTest)
	}

}
