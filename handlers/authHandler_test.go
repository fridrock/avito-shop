package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/mocks"
	"github.com/fridrock/avito-shop/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthBadInputFormat(t *testing.T) {
	_, _, _, authHandler, _ := prepareAuth(t)
	req := httptest.NewRequest("POST", "/api/auth", strings.NewReader(`{"username":"someusername"`))

	status, err := authHandler.Auth(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.NotNil(t, err)
}

func TestAuthCreateUserErrorWhileHashingPassword(t *testing.T) {
	userStorage, _, passwordHasher, authHandler, req := prepareAuth(t)
	userStorage.On("FindUserByUsername", "someusername").Return(storage.User{}, errors.Join(sql.ErrNoRows, fmt.Errorf("no such user")))
	passwordHasher.On("HashPassword", mock.Anything).Return("", fmt.Errorf("error in hashing password"))

	status, err := authHandler.Auth(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, fmt.Errorf("error in hashing password").Error(), err.Error())
}

func TestAuthCreateUserErrorWhileSaving(t *testing.T) {
	userStorage, _, passwordHasher, authHandler, req := prepareAuth(t)
	userStorage.On("FindUserByUsername", "someusername").Return(storage.User{}, errors.Join(sql.ErrNoRows, fmt.Errorf("no such user")))
	userStorage.On("SaveUser", mock.Anything).Return(uuid.New(), fmt.Errorf("error saving in database"))
	passwordHasher.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	status, err := authHandler.Auth(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, fmt.Errorf("error saving in database").Error(), err.Error())

}

func TestAuthCreateUserSuccess(t *testing.T) {
	userStorage, tokenService, passwordHasher, authHandler, req := prepareAuth(t)
	userStorage.On("FindUserByUsername", "someusername").Return(storage.User{}, errors.Join(sql.ErrNoRows, fmt.Errorf("no such user")))
	userStorage.On("SaveUser", mock.Anything).Return(uuid.New(), nil)
	tokenService.On("GenerateToken", mock.Anything, mock.Anything).Return(api.AuthResponse{
		Token: "Valid token",
	}, nil)
	passwordHasher.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	status, err := authHandler.Auth(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}

func TestAuthUserExistCorrectPassword(t *testing.T) {
	userStorage, tokenService, passwordHasher, authHandler, req := prepareAuth(t)
	userStorage.On("FindUserByUsername", "someusername").Return(storage.User{}, nil)
	tokenService.On("GenerateToken", mock.Anything, mock.Anything).Return(api.AuthResponse{
		Token: "Valid token",
	}, nil)
	passwordHasher.On("CheckPassword", mock.Anything, mock.Anything).Return(true)

	status, err := authHandler.Auth(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}
func TestAuthUserExistWrongPassword(t *testing.T) {
	userStorage, _, passwordHasher, authHandler, req := prepareAuth(t)
	userStorage.On("FindUserByUsername", "someusername").Return(storage.User{}, nil)
	passwordHasher.On("CheckPassword", mock.Anything, mock.Anything).Return(false)

	status, err := authHandler.Auth(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Equal(t, fmt.Errorf("wrong password").Error(), err.Error())
}

func prepareAuth(t *testing.T) (*mocks.UserStorage, *mocks.TokenService, *mocks.PasswordHasher, AuthHandler, *http.Request) {
	userStorage := mocks.NewUserStorage(t)
	tokenService := mocks.NewTokenService(t)
	passwordHasher := mocks.NewPasswordHasher(t)
	authHandler := NewAuthHandler(userStorage, tokenService, passwordHasher)
	req := httptest.NewRequest("POST", "/api/auth", strings.NewReader(`{"username":"someusername", "password":"password"}`))
	return userStorage, tokenService, passwordHasher, authHandler, req
}
