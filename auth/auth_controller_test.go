package auth

import (
    "fmt"
    "net/http"
    "testing"

    "github.com/gin-gonic/gin"    
    "github.com/stretchr/testify/assert"
)

type TestAuthenticator struct {
    AuthUrl string
}

func (auth TestAuthenticator) AuthUrl(state string) string {
    return auth.AuthUrl
}

func (auth Authenticator) Token(state string, r *http.Request) (oauth2.Token, error) {
    // TODO
}

func TestFinishAuth(t *testing.T) {
    code := "Code"
    request := http.Request{URL: http.URL{RawQuery: fmt.Sprintf("http://example.localhost?code=%s", code)}}
    authenticator := TestAuthenticator{}
    state := "State"

    
} 
