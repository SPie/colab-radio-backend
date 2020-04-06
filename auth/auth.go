package auth

import (
    "github.com/zmb3/spotify"
)

func NewAuthenticator(callbackUrl string, clientId string, secret string) spotify.Authenticator {
    authenticator := spotify.NewAuthenticator(
	callbackUrl,
	spotify.ScopeUserReadPrivate,
	spotify.ScopeUserReadEmail,
    )
    authenticator.SetAuthInfo(clientId, secret)

    return authenticator
}
