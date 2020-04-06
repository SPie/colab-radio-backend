package router

import (
    "github.com/gin-gonic/gin"    

    "colab-radio/auth"
    "colab-radio/db"
    "colab-radio/user"
)

func SetUp(connectionFactory *db.ConnectionFactory, authControllerFactory auth.AuthControllerFactory) *gin.Engine {
    engine := gin.New()

    api := engine.Group("/api")
    {
	api.GET("/auth", func(c *gin.Context) {
	    authControllerFactory.NewAuthController().InitAuth(c)
	})
	api.POST("/auth-finish", func(c *gin.Context) {
	    defer connectionFactory.Close()

	    authControllerFactory.NewAuthController().FinishAuth(c, user.NewUserRepository(connectionFactory))
	})
    }

    return engine
}
