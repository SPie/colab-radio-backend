package user

import (
    "github.com/jinzhu/gorm" 

    "colab-radio/db"
)

const USER_REPOSITORY = "UserRepository"

type UserRepository struct {
    connectionFactory *db.ConnectionFactory
}

func NewUserRepository(connectionFactory *db.ConnectionFactory) *UserRepository {
    return &UserRepository{connectionFactory: connectionFactory}
}

func (userRepository *UserRepository) getConnection() *gorm.DB {
    return userRepository.connectionFactory.GetConnection()
}

func (userRepository *UserRepository) GetUserBySpotifyId(spotifyId string) *User{
    user := User{}
    userRepository.getConnection().First(&user, &User{SpotifyId: spotifyId})

    return &user
}

func (userRepository *UserRepository) Exists(spotifyId string) bool {
    return userRepository.GetUserBySpotifyId(spotifyId).ID != 0
}

func (userRepository *UserRepository) CreateUser(spotifyId string, email string) *User {
    user := NewUser(spotifyId, email)
    userRepository.getConnection().Create(&user)

    return &user
}
