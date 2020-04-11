package auth

import (
    "strconv"
    "time"

    "golang.org/x/oauth2"
    "github.com/zmb3/spotify"
)

const (
    AUTHENTICATOR   = "Authenticator"
    CLIENT          = "Client"
    AUTH_CONTROLLER = "AuthController"

    ACCESS_TOKEN  	= "AccessToken"
    ACCESS_TOKEN_EXPIRY = "AccessTokenExpiry"
    REFRESH_TOKEN 	= "RefreshToken"
)

type AuthenticatorFactory struct {
    callbackUrl string
    clientId string
    secret string
}

func NewAuthenticatorFactory(callbackUrl string, clientId string, secret string) AuthenticatorFactory {
    return AuthenticatorFactory{
	callbackUrl: callbackUrl,
	clientId: clientId,
	secret: secret,
    }
}

func (authenticatorFactory AuthenticatorFactory) NewAuthenticator() spotify.Authenticator {
    authenticator := spotify.NewAuthenticator(
	authenticatorFactory.callbackUrl,
	spotify.ScopeUserReadPrivate,
	spotify.ScopeUserReadEmail,
    )
    authenticator.SetAuthInfo(authenticatorFactory.clientId, authenticatorFactory.secret)

    return authenticator
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
