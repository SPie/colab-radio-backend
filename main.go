package main

import (
	"os"

	"github.com/joho/godotenv"

	"colab-radio/auth"
	"colab-radio/context"
	"colab-radio/db"
	"colab-radio/router"
	"colab-radio/track"
	"colab-radio/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	connectionHandler, err := setUpConnectionHandler()
	if err != nil {
		panic(err)
	}
	defer connectionHandler.Close()

	appContext := context.NewAppContext(
		user.NewRepository(connectionHandler),
		setUpAuthController(),
		setUpUserController(),
		track.NewController(track.NewService()),
	)

	router := router.SetUp(appContext)

	err = router.Run(os.Getenv("HOST_ADDRESS"))
	if err != nil {
		panic(err)
	}
}

func setUpConnectionHandler() (*db.ConnectionHandler, error) {
	return db.New(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
}

func setUpAuthController() auth.Controller {
	return auth.NewController(
		os.Getenv("AUTH_CALLBACK_URL"),
		os.Getenv("SPOTIFY_CLIENT_ID"),
		os.Getenv("SPOTIFY_SECRET"),
	)
}

func setUpUserController() user.Controller {
	return user.NewController()
}
