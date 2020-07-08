package user

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Controller user routes
type Controller interface {
	GetAuthenticatedUser() gin.HandlerFunc
}

type controller struct{}

// NewController initializes a new user controller
func NewController() Controller {
	return controller{}
}

// GetAuthenticatedUser returns the authenticated user set by authenticate middleware
func (controller controller) GetAuthenticatedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticatedUser, _ := c.Get("authenticated-user")

		userData, err := json.Marshal(authenticatedUser)
		if err != nil {
			c.JSON(500, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(200, map[string]string{"user": string(userData)})
	}
}
