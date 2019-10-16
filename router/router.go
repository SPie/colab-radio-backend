package router

import (
    "os"

    "github.com/gin-gonic/gin"    
    "github.com/spie/colab-radio-backend/auth"
)

type Router struct {
    engine *gin.Engine
}

func SetUp() *Router {
    engine := gin.New()

    api := engine.Group("/api")
    {
	api.GET("/auth", func(c *gin.Context) {
	    auth.InitAuth(c, createAuthenticator(), os.Getenv("AUTH_STATE"))
	})

	api.GET("/auth-finish", func(c *gin.Context) {
	    // TODO
	})
    }

    return &Router{engine: engine}
}

func (router *Router) Run() error {
    err := router.engine.Run()
    if err != nil {
	return err
    }

    return nil
}

func createAuthenticator() auth.Authenticator {
    return auth.New(
    	os.Getenv("AUTH_REDIRECT_URL"),
	os.Getenv("AUTH_CLIENT_ID"),
	os.Getenv("AUTH_SECRET"),
    )
}
