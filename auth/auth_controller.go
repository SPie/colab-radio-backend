package auth

import (
    "github.com/gin-gonic/gin"    
)

func InitAuth(c *gin.Context, auth Authenticator, state string) *gin.Context {
    redirectUrl := c.GetString("redirectUrl")
    if redirectUrl == "" {
	
    }

    c.JSON(303, gin.H{"authUrl": auth.AuthUrl(state)})

    return c
}

func FinishAuth(c *gin.Context, auth Authenticator, state string) *gin.Context {
    
} 
