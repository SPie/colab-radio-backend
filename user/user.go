package user

import (
    "github.com/jinzhu/gorm" 
    "github.com/google/uuid"
)

type User struct {
    gorm.Model
    Uuid string
    SpotifyId string
    Email string
}

func NewUser(spotifyId string, email string) User {
    return User{SpotifyId: spotifyId, Email: email}
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
    u, err := uuid.NewUUID()
    if err != nil {
	return err
    }

    user.Uuid = u.String()

    return nil
}
