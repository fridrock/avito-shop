package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/mocks"
	"github.com/fridrock/avito-shop/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetInfoError(t *testing.T) {
	userId, infoStorage, userStorage, infoHandler, req := prepareGetInfo(t)
	infoStorage.On("GetInfoResponse", userId).Return(api.InfoResponse{}, fmt.Errorf("some error"))
	userStorage.On("GetUserById", mock.Anything).Return(storage.User{
		Coins: 100,
	}, nil)
	status, err := infoHandler.GetInfo(httptest.NewRecorder(), req)
	assert.Equal(t, status, http.StatusInternalServerError)
	assert.Equal(t, fmt.Errorf("some error").Error(), err.Error())
}

func TestGetInfo(t *testing.T) {
	userId, infoStorage, userStorage, infoHandler, req := prepareGetInfo(t)
	infoStorage.On("GetInfoResponse", userId).Return(api.InfoResponse{}, nil)
	userStorage.On("GetUserById", mock.Anything).Return(storage.User{
		Coins: 100,
	}, nil)
	status, err := infoHandler.GetInfo(httptest.NewRecorder(), req)
	assert.Equal(t, status, http.StatusOK)
	assert.Nil(t, err)
}

func prepareGetInfo(t *testing.T) (uuid.UUID, *mocks.InfoStorage, *mocks.UserStorage, InfoHandler, *http.Request) {
	userId := uuid.New()
	infoStorage := mocks.NewInfoStorage(t)
	userStorage := mocks.NewUserStorage(t)
	infoHandler := NewInfoHandler(infoStorage, userStorage)
	return userId, infoStorage, userStorage, infoHandler, prepareGetInfoRequest(userId)
}

func prepareGetInfoRequest(userId uuid.UUID) *http.Request {
	req := httptest.NewRequest("GET", "/api/info", nil)
	return req.WithContext(context.WithValue(req.Context(), auth.Key, userId))
}
