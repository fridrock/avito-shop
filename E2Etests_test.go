package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthE2E(t *testing.T) {
	//Регистрируем нового пользователя
	status, authResponse, errorResponse, err := MakeAuthRequest(testServer.URL, `{"username":"user4", "password":"user4"}`)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotEmpty(t, authResponse)
	assert.Empty(t, errorResponse)
	// Авторизовываем существующего пользователя
	status, authResponse, errorResponse, err = MakeAuthRequest(testServer.URL, `{"username":"user1", "password":"user1"}`)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotEmpty(t, authResponse)
	assert.Empty(t, errorResponse)
	// Неправильный пароль
	status, authResponse, errorResponse, err = MakeAuthRequest(testServer.URL, `{"username":"user1", "password":"user"}`)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Empty(t, authResponse)
	assert.NotEmpty(t, errorResponse)
}

func TestSendCoinE2E(t *testing.T) {
	_, authResponse, _, err := MakeAuthRequest(testServer.URL, `{"username":"user1", "password":"user1"}`)
	if err != nil {
		t.Fatalf(fmt.Sprintf("error authorizing %v", err.Error()))
	}
	status, _, err := MakeSendCoinRequest(testServer.URL, authResponse.Token, `{"toUser":"user2", "amount": 10}`)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	// Нельзя отправлять самому себе
	status, errorResponse, err := MakeSendCoinRequest(testServer.URL, authResponse.Token, `{"toUser":"user1", "amount":10}`)
	assert.Nil(t, err)
	assert.NotEmpty(t, errorResponse)
	assert.Equal(t, http.StatusBadRequest, status)
	// Нельзя отправлять больше чем есть
	status, errorResponse, err = MakeSendCoinRequest(testServer.URL, authResponse.Token, `{"toUser":"user2", "amount":1000000}`)
	assert.Nil(t, err)
	assert.NotEmpty(t, errorResponse)
	assert.Equal(t, http.StatusBadRequest, status)
	// Нельзя отправлять отрицательное количество
	status, errorResponse, err = MakeSendCoinRequest(testServer.URL, authResponse.Token, `{"toUser":"user2", "amount":-10}`)
	assert.Nil(t, err)
	assert.NotEmpty(t, errorResponse)
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestBuyE2E(t *testing.T) {
	_, authResponse, _, err := MakeAuthRequest(testServer.URL, `{"username":"user2", "password":"user2"}`)
	if err != nil {
		t.Fatalf(fmt.Sprintf("error authorizing %v", err.Error()))
	}
	//Нельзя купить товар, которого нет
	status, errorResponse, err := MakeBuyRequest(testServer.URL, authResponse.Token, "nonexisting")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.NotEmpty(t, errorResponse)
	//Нельзя купить товар, который стоит больше чем есть
	status, errorResponse, err = MakeBuyRequest(testServer.URL, authResponse.Token, "pink-hoody")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.NotEmpty(t, errorResponse)
	//Успешная покупка товара
	status, errorResponse, err = MakeBuyRequest(testServer.URL, authResponse.Token, "socks")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Empty(t, errorResponse)
}

func TestGetInfoE2E(t *testing.T) {
	_, authResponse, _, err := MakeAuthRequest(testServer.URL, `{"username":"user2", "password":"user2"}`)
	if err != nil {
		t.Fatalf(fmt.Sprintf("error authorizing %v", err.Error()))
	}
	status, infoResponse, errorResponse, err := MakeGetInfoRequest(testServer.URL, authResponse.Token)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Empty(t, errorResponse)
	assert.NotEmpty(t, infoResponse.CoinHistory)
	assert.NotEmpty(t, infoResponse.Coins)
	assert.NotEmpty(t, infoResponse.Inventory)
	assert.NotEmpty(t, infoResponse.Sent)
}
