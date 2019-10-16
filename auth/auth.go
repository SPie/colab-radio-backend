package auth

import (
    "net/http"
    "oauth2"

    "github.com/zmb3/spotify"
)

type Authenticator struct {
    auth spotify.Authenticator
}

func (auth Authenticator) AuthUrl(state string) string {
    return auth.auth.AuthURL(state)
}

func (auth Authenticator) Token(state string, r *http.Request) (*oauth2.Token, error) {
    // TODO
}

func (auth Authenticator) NewClient(token *oauth2.Token) Client {
    // TODO
}

func New(redirectUrl string, clientId string, secret string) Authenticator {
    auth := spotify.NewAuthenticator(
	redirectUrl,
	spotify.ScopeUserReadPrivate,
	spotify.ScopeUserReadEmail,
    )
    auth.SetAuthInfo(clientId, secret)

    return Authenticator{auth: auth}
}
