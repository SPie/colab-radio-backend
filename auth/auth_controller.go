package auth

import (
    "errors"
    "fmt"
    "math/rand"
    "time"

    "github.com/gin-gonic/gin"    
    "github.com/zmb3/spotify"

    "colab-radio/user"
)

const (
    AUTH_CONTROLLER = "AuthController"
    AUTHENTICATED_USER = "AuthenticatedUser"
)

type AuthControllerFactory struct {}

func NewAuthControllerFactory() AuthControllerFactory {
    return AuthControllerFactory{}
}

func (authControllerFactory AuthControllerFactory) NewAuthController() AuthController {
     return AuthController{
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
    stateCreator StateCreator
}

func (authController AuthController) InitAuth(c *gin.Context) *gin.Context {
    state := authController.stateCreator()
    c.Header("X-Authentication-State", state)
    c.JSON(303, map[string]string{"authUrl": getAuthenticator(c).AuthURL(state)})

    return c
}

func (authController AuthController) FinishAuth(c *gin.Context) *gin.Context {
    authenticator := getAuthenticator(c)

    token, err := authenticator.Token(c.GetHeader("X-Authentication-State"), c.Request)
    if err != nil {
	return notAuthenticated(c, err)
    }

    client := authenticator.NewClient(token)
    spotifyUser, err := (&client).CurrentUser()
    if err != nil {
	return notAuthenticated(c, err)
    }

    userRepository := getUserRepository(c)
    if !userRepository.Exists(spotifyUser.ID) {
	userRepository.CreateUser(spotifyUser.ID, spotifyUser.Email)
    }

    attachAccessToken(c, token.AccessToken, token.Expiry)
    c.SetCookie(REFRESH_TOKEN, token.RefreshToken, 0, "", "", false, true)
    c.JSON(204, map[string]string{})

    return c    
}

func getAuthenticator(c *gin.Context) spotify.Authenticator {
    authenticator, _ := c.Get(AUTHENTICATOR)
    return authenticator.(spotify.Authenticator)
}

func getUserRepository(c *gin.Context) *user.UserRepository {
    userRepository, _ := c.Get(user.USER_REPOSITORY)
    return userRepository.(*user.UserRepository)
}

func notAuthenticated(c *gin.Context, err error) *gin.Context {
    fmt.Println(err)
    c.JSON(401, map[string]string{})
    return c   
}

func attachAccessToken(c *gin.Context, accessToken string, expiry time.Time) {
    expiryTimestamp := int(expiry.Sub(time.Now()).Seconds())
    c.SetCookie(ACCESS_TOKEN, accessToken, expiryTimestamp, "", "", false, true)
    c.SetCookie(ACCESS_TOKEN_EXPIRY, string(expiryTimestamp), expiryTimestamp, "", "", false, true)
}

func Authorization() gin.HandlerFunc {
    return func(c *gin.Context) {
	refreshToken, err := c.Cookie(REFRESH_TOKEN)
	if err != nil {
	    c.AbortWithError(401, errors.New("No refresh token"))
	    return
	}

	accessToken, _ := c.Cookie(ACCESS_TOKEN)

	token := createOAuthToken(accessToken, refreshToken)

	accessTokenExpiry, err := c.Cookie(ACCESS_TOKEN_EXPIRY)
	if err == nil {
	    token.Expiry = parseAccessTokenExpiry(accessTokenExpiry)
	}

	authenticator := getAuthenticator(c)	

	client := authenticator.NewClient(token)
	
	spotifyUser, err := client.CurrentUser()
	if err != nil {
	    c.AbortWithError(401, err)
	    return
	}

	userRepository := getUserRepository(c)
	user := userRepository.GetUserBySpotifyId(spotifyUser.ID)
	if user.ID == 0 {
	    c.AbortWithError(401, errors.New("User doesn't exist"))
	    return
	}

	c.Set(AUTHENTICATED_USER, user)

	c.Next()

	attachAccessToken(c, token.AccessToken, token.Expiry)
    }    
}
