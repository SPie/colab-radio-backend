package main

import (
    "os"

    "github.com/joho/godotenv"

    "colab-radio/auth"
    "colab-radio/db"
    "colab-radio/router"
)

func main() {
    err := godotenv.Load()
    if err != nil {
	panic(err)
    }

    router := router.SetUp(setUpConnectionFactory(), setUpAuthenticatorFactory(), setUpAuthControllerFactory())
    
    err = router.Run()
    if err != nil {
	panic(err)
    }
}

func setUpAuthenticatorFactory() auth.AuthenticatorFactory {
    return auth.NewAuthenticatorFactory(
    	os.Getenv("AUTH_CALLBACK_URL"),
	os.Getenv("SPOTIFY_CLIENT_ID"),
	os.Getenv("SPOTIFY_SECRET"),
    )
}

func setUpAuthControllerFactory() auth.AuthControllerFactory {
    return auth.NewAuthControllerFactory()
}

func setUpConnectionFactory() *db.ConnectionFactory {
    return db.NewConnectionFactory(
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_DATABASE"),
    )
}
