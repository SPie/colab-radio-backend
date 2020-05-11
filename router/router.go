package router

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"colab-radio/auth"
	"colab-radio/db"
	"colab-radio/user"
)

func SetUp(connectionFactory *db.ConnectionFactory, authenticatorFactory auth.AuthenticatorFactory, authControllerFactory auth.AuthControllerFactory, userControllerFactory user.UserControllerFactory) *gin.Engine {
	engine := gin.New()

	engine.Use(setUpCors())

	engine.Use(setUpRepositories(connectionFactory))

	engine.Use(setUpAuthenticator(authenticatorFactory))

	engine.Use(setUpControllers(authControllerFactory, userControllerFactory))

	api := engine.Group("/api")
	{
		api.GET("/auth", func(c *gin.Context) {
			getAuthController(c).InitAuth(c)
		})
		api.POST("/auth-finish", func(c *gin.Context) {
			getAuthController(c).FinishAuth(c)
		})

		api.Use(auth.Authorization())
		{
			api.GET("/users", func(c *gin.Context) {
				getUserController(c).GetAuthenticatedUser(c)
			})
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

func setUpRepositories(connectionFactory *db.ConnectionFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer connectionFactory.Close()

		c.Set(user.USER_REPOSITORY, user.NewUserRepository(connectionFactory))

		c.Next()
	}
}

func setUpAuthenticator(authenticatorFactory auth.AuthenticatorFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(auth.AUTHENTICATOR, authenticatorFactory.NewAuthenticator())

		c.Next()
	}
}

func setUpControllers(authControllerFactory auth.AuthControllerFactory, userControllerFactory user.UserControllerFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(auth.AUTH_CONTROLLER, authControllerFactory.NewAuthController())
		c.Set(user.USER_CONTROLLER, userControllerFactory.NewUserController())

		c.Next()
	}
}

func getAuthController(c *gin.Context) auth.AuthController {
	authController, _ := c.Get(auth.AUTH_CONTROLLER)

	return authController.(auth.AuthController)
}

func getUserController(c *gin.Context) user.UserController {
	userController, _ := c.Get(user.USER_CONTROLLER)

	return userController.(user.UserController)
}
