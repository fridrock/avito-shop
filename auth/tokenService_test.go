package auth

import (
	"os"
	"testing"

	"github.com/fridrock/avito-shop/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var tknService TokenService

func TestMain(m *testing.M) {
	//setup
	os.Setenv("SECRET_KEY", "SECRET_FOR_TEST")
	tknService = NewTokenService()
	//running test
	m.Run()
}

func TestParseTokenSuccess(t *testing.T) {
	id := uuid.New()
	authResponse, _ := tknService.GenerateToken(api.AuthRequest{
		Username: "someusername",
		Password: "somePassword",
	}, id)
	parsedUser, err := tknService.ValidateToken(authResponse.Token)
	assert.Nil(t, err)
	assert.Equal(t, id.String(), parsedUser.Id.String())
	assert.Equal(t, "someusername", parsedUser.Username)
}

func TestParseExpiredToken(t *testing.T) {
	expiredTokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyOTY3NjUsImlkIjoiYTk1MTdkMGItODY1Zi00ODhmLTk0NzItNTE0ZGEyNGQ1MjczIiwidXNlcm5hbWUiOiJzb21ldXNlcm5hbWUifQ.RixFXoRMyCqbnh4YdEyrPgDRF-_x-FA14TG7_SS5sGg"
	_, err := tknService.ValidateToken(expiredTokenString)
	assert.NotNil(t, err)
	assert.Equal(t, "token has invalid claims: token is expired", err.Error())
}
