package auth

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/gin-gonic/gin"    
    "github.com/zmb3/spotify"

    "colab-radio/user"
)

type AuthControllerFactory struct {
    authCallbackUrl string
    spotifyClientId string
    spotifySecret string
}

func NewAuthControllerFactory(authCallbackUrl string, spotifyClientId string, spotifySecret string) AuthControllerFactory {
    return AuthControllerFactory{authCallbackUrl: authCallbackUrl, spotifyClientId: spotifyClientId, spotifySecret: spotifySecret}
}

func (authControllerFactory AuthControllerFactory) NewAuthController() AuthController {
     return AuthController{
	authenticator: NewAuthenticator(
	    authControllerFactory.authCallbackUrl,
	    authControllerFactory.spotifyClientId,
	    authControllerFactory.spotifySecret,
	),
	stateCreator: func () string {
	    alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	    state := make([]rune, 32)
	    for i := range state {
		state[i] = alphabet[rand.Intn(len(alphabet))]
	    }

	    return string(state)
	},
    }   
}

type StateCreator func() string

type AuthController struct {
    authenticator spotify.Authenticator
    stateCreator StateCreator
}

func (authController AuthController) InitAuth(c *gin.Context) *gin.Context {
    state := authController.stateCreator()
    c.Header("X-Authentication-State", state)
    c.JSON(303, map[string]string{"authUrl": authController.authenticator.AuthURL(state)})

    return c
}

func (authController AuthController) FinishAuth(c *gin.Context, userRepository *user.UserRepository) *gin.Context {
    token, err := authController.authenticator.Token(c.GetHeader("X-Authentication-State"), c.Request)
    if err != nil {
	return notAuthenticated(c, err)
    }

    client := authController.authenticator.NewClient(token)
    spotifyUser, err := (&client).CurrentUser()
    if err != nil {
	return notAuthenticated(c, err)
    }

    if !userRepository.Exists(spotifyUser.ID) {
	userRepository.CreateUser(spotifyUser.ID, spotifyUser.Email)
    }

    c.SetCookie("AccessToken", token.AccessToken, int(token.Expiry.Sub(time.Now()).Seconds()), "", "", false, true)
    c.SetCookie("RefreshToken", token.RefreshToken, 0, "", "", false, true)
    c.JSON(200, map[string]string{})

    return c    
}

func notAuthenticated(c *gin.Context, err error) *gin.Context {
    fmt.Println(err)
    c.JSON(401, map[string]string{})
    return c   
}

