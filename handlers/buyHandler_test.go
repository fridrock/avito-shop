package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/mocks"
	"github.com/fridrock/avito-shop/storage"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestBuyEmptyItemName(t *testing.T) {
	_, _, _, buyHandler, req := prepare(t)

	status, err := buyHandler.Buy(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, err.Error(), "empty item name")
}

func TestBuyNonExistingItem(t *testing.T) {
	_, _, productStorage, buyHandler, req := prepare(t)
	productStorage.On("FindProductByName", "name").Return(storage.Product{}, errors.Join(sql.ErrNoRows, fmt.Errorf("no such product")))
	req = mux.SetURLVars(req, map[string]string{"item": "name"})

	status, err := buyHandler.Buy(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, errors.Join(sql.ErrNoRows, fmt.Errorf("no such product")).Error(), err.Error())
}

func TestBuyNotEnoughMoney(t *testing.T) {
	userId, userStorage, productStorage, buyHandler, req := prepare(t)
	req = mux.SetURLVars(req, map[string]string{"item": "name"})
	userStorage.On("CheckEnoughCoins", 100, userId).Return(false)
	productStorage.On("FindProductByName", "name").Return(storage.Product{
		Price: 100,
	}, nil)

	status, err := buyHandler.Buy(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, fmt.Errorf("not enough money").Error(), err.Error())
}

func TestBuyErrorInTransaction(t *testing.T) {
	userId, userStorage, productStorage, buyHandler, req := prepare(t)
	req = mux.SetURLVars(req, map[string]string{"item": "name"})
	userStorage.On("CheckEnoughCoins", 100, userId).Return(true)
	productStorage.On("FindProductByName", "name").Return(storage.Product{
		Price: 100,
	}, nil)
	productStorage.On("Buy", userId, storage.Product{Price: 100}).Return(fmt.Errorf("error in transaction"))
	status, err := buyHandler.Buy(httptest.NewRecorder(), req)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, fmt.Errorf("error in transaction").Error(), err.Error())
}

func TestBuySuccess(t *testing.T) {
	userId, userStorage, productStorage, buyHandler, req := prepare(t)
	req = mux.SetURLVars(req, map[string]string{"item": "name"})
	userStorage.On("CheckEnoughCoins", 100, userId).Return(true)
	productStorage.On("FindProductByName", "name").Return(storage.Product{
		Price: 100,
	}, nil)
	productStorage.On("Buy", userId, storage.Product{Price: 100}).Return(nil)

	status, err := buyHandler.Buy(httptest.NewRecorder(), req)

	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}

func prepare(t *testing.T) (uuid.UUID, *mocks.UserStorage, *mocks.ProductStorage, BuyHandler, *http.Request) {
	userStorage := mocks.NewUserStorage(t)
	productStorage := mocks.NewProductStorage(t)
	buyHandler := NewBuyHandler(productStorage, userStorage)
	userId := uuid.New()
	req := prepareRequest(userId)
	return userId, userStorage, productStorage, buyHandler, req
}

func prepareRequest(userId uuid.UUID) *http.Request {
	req := httptest.NewRequest("GET", "/api/buy/name", nil)
	return req.WithContext(context.WithValue(req.Context(), auth.Key, userId))
}
