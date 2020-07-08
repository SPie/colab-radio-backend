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

// Repository Interface for User database handling
type Repository interface {
	GetUserBySpotifyID(spotifyID string) *User
	Exists(spotifyID string) bool
	CreateUser(spotifyID string, email string) *User
}

// repository the Gorm implementation of UserRepository
type repository struct {
	connectionHandler *db.ConnectionHandler
}

// NewRepository created new UserRepository
func NewRepository(connectionHandler *db.ConnectionHandler) Repository {
	return &repository{connectionHandler: connectionHandler}
}

// GetUserBySpotifyID get the id provided by spotify
func (repository *repository) GetUserBySpotifyID(spotifyID string) *User {
	user := User{}
	repository.connectionHandler.GetConnection().First(&user, &User{SpotifyID: spotifyID})

	return &user
}

// Exists checks if a spotify user already exiss
func (repository *repository) Exists(spotifyID string) bool {
	return repository.GetUserBySpotifyID(spotifyID).ID != 0
}

// CreateUser creates new user with spotify id and email
func (repository *repository) CreateUser(spotifyID string, email string) *User {
	user := NewUser(spotifyID, email)
	repository.connectionHandler.GetConnection().Create(&user)

	return &user
}
