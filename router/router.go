package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"colab-radio/auth"
	"colab-radio/context"
)

// SetUp initializes all routes and middlewares
func SetUp(appContext *context.AppContext) *gin.Engine {
	engine := gin.Default()

	engine.Use(setUpCors())

	api := engine.Group("/api")
	{
		api.GET("/auth", appContext.AuthController.InitAuth(auth.CreateState))
		api.POST("/auth-finish", appContext.AuthController.FinishAuth(appContext.UserRepository))

		api.Use(appContext.AuthController.Authentication(appContext.UserRepository))
		{
			api.GET("/users", appContext.UserController.GetAuthenticatedUser())

			api.GET("/tracks", appContext.TrackController.Search())
		}
	}

	return engine
}

func setUpCors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Authentication-State"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Authentication-State"},
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{"http://localhost:8080"}, //TODO
	}

	return cors.New(config)
}
