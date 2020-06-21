package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"

	"colab-radio/user"
)

// AuthController all authentication routes
type AuthController struct {
	authenticator spotify.Authenticator
}

// NewAuthController creates a new AuthController
func NewAuthController(authCallbackURL string, clientID string, secret string) AuthController {
	authenticator := spotify.NewAuthenticator(
		authCallbackURL,
		spotify.ScopeUserReadPrivate,
		spotify.ScopeUserReadEmail,
	)
	authenticator.SetAuthInfo(clientID, secret)

	return AuthController{authenticator: authenticator}
}

// InitAuth initializes the authentication flow
// returns a authentication state string and the spotify auth url
func (authController AuthController) InitAuth(stateCreator func() string) gin.HandlerFunc {
	return func(c *gin.Context) {
		state := stateCreator()
		c.Header("X-Authentication-State", state)
		c.JSON(200, map[string]string{"authUrl": authController.authenticator.AuthURL(state)})
	}
}

// FinishAuth finalizes the authentication flow
func (authController AuthController) FinishAuth(userRepository user.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := authController.authenticator.Token(c.GetHeader("X-Authentication-State"), c.Request)
		if err != nil {
			notAuthenticated(c, err)
			return
		}

		client := authController.authenticator.NewClient(token)
		spotifyUser, err := (&client).CurrentUser()
		if err != nil {
			notAuthenticated(c, err)
			return
		}

		if !userRepository.Exists(spotifyUser.ID) {
			userRepository.CreateUser(spotifyUser.ID, spotifyUser.Email)
		}

		attachAccessToken(c, token.AccessToken, token.Expiry)
		c.SetCookie("refresh-token", token.RefreshToken, 0, "", "", false, true)
		c.JSON(204, map[string]string{})
	}
}

// Authentication Middleware
func (authController AuthController) Authentication(userRepository user.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh-token")
		if err != nil {
			c.AbortWithError(401, errors.New("No refresh token"))
			return
		}

		accessToken, _ := c.Cookie("access-token")

		token := createOAuthToken(accessToken, refreshToken)

		accessTokenExpiry, err := c.Cookie("access-token-expiry")
		if err == nil {
			token.Expiry = parseAccessTokenExpiry(accessTokenExpiry)
		}

		client := authController.authenticator.NewClient(token)

		spotifyUser, err := client.CurrentUser()
		if err != nil {
			c.AbortWithError(401, err)
			return
		}

		user := userRepository.GetUserBySpotifyID(spotifyUser.ID)
		if user.ID == 0 {
			c.AbortWithError(401, errors.New("User doesn't exist"))
			return
		}

		c.Set("authenticated-user", user)

		c.Next()

		attachAccessToken(c, token.AccessToken, token.Expiry)
	}
}

func notAuthenticated(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(401, map[string]string{})
}

func attachAccessToken(c *gin.Context, accessToken string, expiry time.Time) {
	expiryTimestamp := int(expiry.Sub(time.Now()).Seconds())
	c.SetCookie("access-token", accessToken, expiryTimestamp, "", "", false, true)
	c.SetCookie("access-token-expiry", string(expiryTimestamp), expiryTimestamp, "", "", false, true)
}

func createOAuthToken(accessToken string, refreshToken string) *oauth2.Token {
	return &oauth2.Token{AccessToken: accessToken, RefreshToken: refreshToken}
}

func parseAccessTokenExpiry(accessTokenExpiry string) time.Time {
	timestamp, err := strconv.ParseInt(accessTokenExpiry, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(timestamp, 0)
}

// CreateState function to return a random state
func CreateState() string {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	state := make([]rune, 32)
	for i := range state {
		state[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(state)
}