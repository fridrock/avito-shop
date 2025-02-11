package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/mocks"
	"github.com/fridrock/avito-shop/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSendCoinWrongInputFormat(t *testing.T) {
	userId, _, _, sendCoinHandler, _ := prepareSendCoin(t)
	req := httptest.NewRequest("POST", "/api/sendCoin", strings.NewReader(`{"toUser":"someUser"}`))
	req = req.WithContext(context.WithValue(req.Context(), auth.Key, userId))

	status, err := sendCoinHandler.SendCoin(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.NotNil(t, err)
}

func TestSendCoinWrongInput(t *testing.T) {
	userId, userStorage, _, sendCoinHandler, req := prepareSendCoin(t)
	userStorage.On("CheckEnoughCoins", 10, userId).Return(false)
	userStorage.On("FindUserByUsername", "someUser").Return(storage.User{
		Id: userId,
	}, nil)

	status, err := sendCoinHandler.SendCoin(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.NotNil(t, err)
}

func TestSendCoinErrorInTransaction(t *testing.T) {
	userId, userStorage, coinStorage, sendCoinHandler, req := prepareSendCoin(t)
	userStorage.On("CheckEnoughCoins", 10, userId).Return(true)
	toUserId := uuid.New()
	userStorage.On("FindUserByUsername", "someUser").Return(storage.User{
		Id: toUserId,
	}, nil)
	coinStorage.On("SendCoin", 10, userId, toUserId).Return(fmt.Errorf("error in transaction"))

	status, err := sendCoinHandler.SendCoin(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.NotNil(t, err)
}

func TestSendCoinSuccess(t *testing.T) {
	userId, userStorage, coinStorage, sendCoinHandler, req := prepareSendCoin(t)
	userStorage.On("CheckEnoughCoins", 10, userId).Return(true)
	toUserId := uuid.New()
	userStorage.On("FindUserByUsername", "someUser").Return(storage.User{
		Id: toUserId,
	}, nil)
	coinStorage.On("SendCoin", 10, userId, toUserId).Return(nil)

	status, err := sendCoinHandler.SendCoin(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}

func prepareSendCoin(t *testing.T) (uuid.UUID, *mocks.UserStorage, *mocks.CoinStorage, SendCoinHandler, *http.Request) {
	userId := uuid.New()
	userStorage := mocks.NewUserStorage(t)
	coinStorage := mocks.NewCoinStorage(t)
	sendCoinHandler := NewSendCoinHandler(coinStorage, userStorage)
	return userId, userStorage, coinStorage, sendCoinHandler, prepareRequestSendCoin(userId)
}
func prepareRequestSendCoin(userId uuid.UUID) *http.Request {
	req := httptest.NewRequest("POST", "/api/sendCoin", strings.NewReader(`{"toUser":"someUser", "amount":10}`))
	return req.WithContext(context.WithValue(req.Context(), auth.Key, userId))
}
