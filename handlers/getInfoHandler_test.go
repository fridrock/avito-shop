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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetInfoError(t *testing.T) {
	userId, infoStorage, infoHandler, req := prepareGetInfo(t)
	infoStorage.On("GetInfoResponse", userId).Return(api.InfoResponse{}, fmt.Errorf("some error"))
	status, err := infoHandler.GetInfo(httptest.NewRecorder(), req)
	assert.Equal(t, status, http.StatusInternalServerError)
	assert.Equal(t, fmt.Errorf("some error").Error(), err.Error())
}

func TestGetInfo(t *testing.T) {
	userId, infoStorage, infoHandler, req := prepareGetInfo(t)
	infoStorage.On("GetInfoResponse", userId).Return(api.InfoResponse{}, nil)
	status, err := infoHandler.GetInfo(httptest.NewRecorder(), req)
	assert.Equal(t, status, http.StatusOK)
	assert.Nil(t, err)
}

func prepareGetInfo(t *testing.T) (uuid.UUID, *mocks.InfoStorage, InfoHandler, *http.Request) {
	userId := uuid.New()
	infoStorage := mocks.NewInfoStorage(t)
	infoHandler := NewInfoHandler(infoStorage)
	return userId, infoStorage, infoHandler, prepareGetInfoRequest(userId)
}

func prepareGetInfoRequest(userId uuid.UUID) *http.Request {
	req := httptest.NewRequest("GET", "/api/info", nil)
	return req.WithContext(context.WithValue(req.Context(), auth.Key, userId))
}
