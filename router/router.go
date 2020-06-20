package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"colab-radio/auth"
	"colab-radio/router/context"
)

// SetUp initializes all routes and middlewares
func SetUp(appContext *context.AppContext) *gin.Engine {
	engine := gin.Default()

	engine.Use(setUpCors())

	engine.Group("/api")
	{
		engine.GET("/auth", appContext.AuthController.InitAuth(auth.CreateState))
		engine.POST("/auth-finish", appContext.AuthController.FinishAuth(appContext.UserRepository))

		engine.Use(appContext.AuthController.Authentication(appContext.UserRepository))
		{
			engine.GET("/users", appContext.UserController.GetAuthenticatedUser())
		}
	}

	return engine
}

func setUpCors() gin.HandlerFunc {
    config := cors.Config{
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "X-Authentication-State"},
        AllowCredentials: false,
        ExposeHeaders: []string{"X-Authentication-State"},
        MaxAge: 12 * time.Hour,
        AllowAllOrigins: true,
    }
    
    return cors.New(config)
}