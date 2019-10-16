package user

import (
    "github.com/spie/colab-radio-backend/db"
)

type User struct {
    db.Model
    Uuid string
    SpotifyId string
    Email string
}
