package user

import (
   "github.com/jinzhu/gorm" 
)

type User struct {
    gorm.Model
    Uuid string
    SpotifyId string
    Email string
}
