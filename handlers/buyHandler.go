package handlers

import (
	"fmt"
	"net/http"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/storage"
	"github.com/gorilla/mux"
)

type BuyHandler interface {
	Buy(w http.ResponseWriter, r *http.Request) (int, error)
}

type BuyHandlerImpl struct {
	productStorage storage.ProductStorage
	userStorage    storage.UserStorage
}

func (bh *BuyHandlerImpl) Buy(w http.ResponseWriter, r *http.Request) (int, error) {
	userId := auth.UserFromContext(r.Context())
	itemName := mux.Vars(r)["item"]
	if itemName == "" {
		return http.StatusBadRequest, fmt.Errorf("empty item name")
	}

	product, err := bh.productStorage.FindProductByName(itemName)
	if err != nil {
		return http.StatusBadRequest, err
	}

	hasEnoughMoney := bh.userStorage.CheckEnoughCoins(product.Price, userId)
	if !hasEnoughMoney {
		return http.StatusBadRequest, fmt.Errorf("not enough money")
	}
	err = bh.productStorage.Buy(userId, product)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func NewBuyHandler(ps storage.ProductStorage, us storage.UserStorage) BuyHandler {
	return &BuyHandlerImpl{
		productStorage: ps,
		userStorage:    us,
	}
}
