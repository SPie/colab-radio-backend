package auth

import (
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

func TestFinishAuth(t *testing.T) {
    code := "Code"
    context := gin.Context{Keys: map[string]interface{}{"code": code}}
    state := "State"
    authenticator := TestAuthenticator{}
} 
