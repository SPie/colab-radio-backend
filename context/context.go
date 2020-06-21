package context

import (
	"colab-radio/auth"
	"colab-radio/user"
)

// AppContext holds repositories and controllers
type AppContext struct {
	UserRepository user.UserRepository
	AuthController auth.AuthController
	UserController user.UserController
}

// NewAppContext Initializes the AppContext
func NewAppContext(userRepository user.UserRepository, authController auth.AuthController, userController user.UserController) *AppContext {
	return &AppContext{UserRepository: userRepository, AuthController: authController, UserController: userController}
}