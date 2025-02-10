package buy

import (
	"fmt"
	"net/http"

	"github.com/fridrock/avito-shop/utils"
	"github.com/gorilla/mux"
)

type BuyHandler interface {
	Buy(w http.ResponseWriter, r *http.Request) (int, error)
}

type BuyHandlerImpl struct {
	storage BuyStorage
}

func (bh *BuyHandlerImpl) Buy(w http.ResponseWriter, r *http.Request) (int, error) {
	userId := utils.UserIdFromContext(r.Context())
	itemName := mux.Vars(r)["item"]
	if itemName == "" {
		return http.StatusBadRequest, fmt.Errorf("empty item name")
	}
	err := bh.storage.Buy(userId, itemName)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func NewBuyHandler(storage BuyStorage) BuyHandler {
	return &BuyHandlerImpl{
		storage: storage,
	}
}
