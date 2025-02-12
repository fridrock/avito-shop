package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserFromContext(t *testing.T) {
	userId := uuid.New()
	ctx := context.Background()
	ctx = context.WithValue(ctx, Key, userId)
	assert.Equal(t, UserFromContext(ctx), userId)
}

func someHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusOK, nil
}

func TestGetUserFromToken(t *testing.T) {
	tokenService, _ := prepareAuthManager(t)
	// authorizedHandler := authManager.AuthMiddleware(someHandler)
	req := httptest.NewRequest("GET", "http://example.com", nil)
	authManager := authManager{
		tokenService: tokenService,
	}
	_, err := authManager.getUserFromToken(req)

	assert.Equal(t, "empty auth header", err.Error())

	req.Header.Set("Authorization", "Bearer")
	_, err = authManager.getUserFromToken(req)
	assert.Equal(t, "wrong format of token", err.Error())

	req.Header.Set("Authorization", "Bred token")
	_, err = authManager.getUserFromToken(req)
	assert.Equal(t, "wrong format of token", err.Error())

	tokenService.On("ValidateToken", "token").Return(api.UserInfo{}, fmt.Errorf("token invalidated")).Once()
	req.Header.Set("Authorization", "Bearer token")

	_, err = authManager.getUserFromToken(req)
	assert.Equal(t, "token invalidated", err.Error())

	tokenService.On("ValidateToken", "token").Return(api.UserInfo{
		Username: "Someusername",
		Id:       uuid.New(),
	}, nil).Once()
	req.Header.Set("Authorization", "Bearer token")
	_, err = authManager.getUserFromToken(req)
	assert.Nil(t, err)

}

func TestAuthMiddleware(t *testing.T) {
	tokenService, authManager := prepareAuthManager(t)
	req := httptest.NewRequest("GET", "http://example.com", nil)
	authorizedHandler := authManager.AuthMiddleware(someHandler)

	status, _ := authorizedHandler(httptest.NewRecorder(), req)
	assert.Equal(t, http.StatusUnauthorized, status)

	userId := uuid.New()
	tokenService.On("ValidateToken", "token").Return(api.UserInfo{
		Username: "Someusername",
		Id:       userId,
	}, nil)
	req.Header.Set("Authorization", "Bearer token")
	status, err := authorizedHandler(httptest.NewRecorder(), req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

}

func prepareAuthManager(t *testing.T) (*mocks.TokenService, AuthManager) {
	tokenService := mocks.NewTokenService(t)
	return tokenService, NewAuthManager(tokenService)
}
