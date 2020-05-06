package user

import (
    "encoding/json"

    "github.com/gin-gonic/gin"
)

const USER_CONTROLLER = "UserController"

type UserControllerFactory struct {}

func NewUserControllerFactory() UserControllerFactory {
    return UserControllerFactory{}
}

func (userControllerFactory UserControllerFactory) NewUserController() UserController {
    return UserController{}
}

type UserController struct {}

func (userController UserController) GetAuthenticatedUser(c *gin.Context) *gin.Context {
    authenticatedUser, _ := c.Get("AuthenticatedUser")

    userData, err := json.Marshal(authenticatedUser)
    if err != nil {
	c.JSON(500, map[string]string{"error": err.Error()})
	return c
    }

    c.JSON(200, map[string]string{"user": string(userData)})

    return c
}
