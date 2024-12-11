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

	authRoute := api.Group("/auth")
	{
		authRoute.POST("/register", compHandler.RegisterUserCredential)
		authRoute.POST("/login", compHandler.LoginUserCredentials)

		googleRoute := authRoute.Group("/google")
		{
			googleRoute.POST("/login", compHandler.LoginUserGoogle)
		}
	}

	authRoute.Use(middleware.AuthMiddleware())
	{
		authRoute.GET("/auth-test", compHandler.AuthTest)
	}

	predictRoute := api.Group("/predict")
	predictRoute.Use(middleware.AuthMiddleware())
	{
		sibiRoute := predictRoute.Group("/sibi")
		{
			sibiRoute.POST("/image", compHandler.SIBIImagePredict)
			sibiRoute.POST("/video", compHandler.SIBIVideoPredict)
		}

		bisindoRoute := predictRoute.Group("/bisindo")
		{
			bisindoRoute.POST("/image", compHandler.BISINDOImagePredict)
			bisindoRoute.POST("/video", compHandler.BISINDOVideoPredict)
		}
	}
}
