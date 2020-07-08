package context

import (
	"colab-radio/auth"
	"colab-radio/track"
	"colab-radio/user"
)

// AppContext holds repositories and controllers
type AppContext struct {
	UserRepository  user.Repository
	AuthController  auth.Controller
	UserController  user.Controller
	TrackController track.Controller
}

// NewAppContext Initializes the AppContext
func NewAppContext(userRepository user.Repository, authController auth.Controller, userController user.Controller, trackController track.Controller) *AppContext {
	return &AppContext{UserRepository: userRepository, AuthController: authController, UserController: userController, TrackController: trackController}
}
