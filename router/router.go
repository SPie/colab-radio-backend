package router

import (
    "github.com/gin-gonic/gin"    

    "colab-radio/auth"
    "colab-radio/db"
    "colab-radio/user"
)

func SetUp(connectionFactory *db.ConnectionFactory, authenticatorFactory auth.AuthenticatorFactory, authControllerFactory auth.AuthControllerFactory) *gin.Engine {
    engine := gin.New()

    engine.Use(SetUpRepositories(connectionFactory))

    engine.Use(SetUpAuthenticator(authenticatorFactory))

    engine.Use(SetUpControllers(authControllerFactory))

    api := engine.Group("/api")
    {
	api.GET("/auth", func(c *gin.Context) {
	    authControllerFactory.NewAuthController().InitAuth(c)
	})
	api.POST("/auth-finish", func(c *gin.Context) {
	    authControllerFactory.NewAuthController().FinishAuth(c)
	})
	
	// TODO authorization middleware

	api.GET("/users", func (c *gin.Context) {
	    // TODO get authenticated user
	})
    }

    return engine
}

func SetUpRepositories(connectionFactory *db.ConnectionFactory) gin.HandlerFunc {
    return func(c *gin.Context) {
	defer connectionFactory.Close()

	c.Set(user.USER_REPOSITORY, user.NewUserRepository(connectionFactory))

	c.Next()
    }
}

func SetUpAuthenticator(authenticatorFactory auth.AuthenticatorFactory) gin.HandlerFunc {
    return func(c *gin.Context) {
	c.Set(auth.AUTHENTICATOR, authenticatorFactory.NewAuthenticator())

	c.Next()
    }
}

func SetUpControllers(authControllerFactory auth.AuthControllerFactory) gin.HandlerFunc {
    return func(c *gin.Context) {
	c.Set(auth.AUTH_CONTROLLER, authControllerFactory.NewAuthController())

	c.Next()
    }
}
