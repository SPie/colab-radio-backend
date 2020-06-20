package user

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// UserController user routes
type UserController struct{}

// NewUserController initializes a new user controller
func NewUserController() UserController {
    return UserController{}
}

// GetAuthenticatedUser returns the authenticated user set by authenticate middleware
func (userController UserController) GetAuthenticatedUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        authenticatedUser, _ := c.Get("AuthenticatedUser")

        userData, err := json.Marshal(authenticatedUser)
        if err != nil {
            c.JSON(500, map[string]string{"error": err.Error()})
            return
        }

        c.JSON(200, map[string]string{"user": string(userData)})
    }
}
