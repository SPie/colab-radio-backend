package user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"colab-radio/db"
)

// User Model
type User struct {
	gorm.Model
	UUID      string
	SpotifyID string
	Email     string
}

// NewUser initializes a new user
func NewUser(spotifyID string, email string) User {
	return User{SpotifyID: spotifyID, Email: email}
}

// BeforeCreate intercepts the user creation to insert a UUID
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	u, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	user.UUID = u.String()

	return nil
}

// UserRepository Interface for User database handling
type UserRepository interface {
	GetUserBySpotifyID(spotifyID string) *User
	Exists(spotifyID string) bool
	CreateUser(spotifyID string, email string) *User
}

// GormUserRepository the Gorm implementation of UserRepository
type GormUserRepository struct {
	connectionHandler *db.ConnectionHandler
}

// NewUserRepository created new UserRepository
func NewUserRepository(connectionHandler *db.ConnectionHandler) UserRepository {
	return &GormUserRepository{connectionHandler: connectionHandler}
}

// GetUserBySpotifyID get the id provided by spotify
func (userRepository *GormUserRepository) GetUserBySpotifyID(spotifyID string) *User {
	user := User{}
	userRepository.connectionHandler.GetConnection().First(&user, &User{SpotifyID: spotifyID})

	return &user
}

// Exists checks if a spotify user already exiss
func (userRepository *GormUserRepository) Exists(spotifyID string) bool {
	return userRepository.GetUserBySpotifyID(spotifyID).ID != 0
}

// CreateUser creates new user with spotify id and email
func (userRepository *GormUserRepository) CreateUser(spotifyID string, email string) *User {
	user := NewUser(spotifyID, email)
	userRepository.connectionHandler.GetConnection().Create(&user)

	return &user
}
