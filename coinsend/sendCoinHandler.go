package coinsend

import (
	"errors"
	"net/http"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/utils"
)

type SendCoinHandler interface {
	SendCoin(w http.ResponseWriter, r *http.Request) (int, error)
}

type SendCoinHandlerImpl struct {
	storage SendCoinStorage
}

func (sc *SendCoinHandlerImpl) SendCoin(w http.ResponseWriter, r *http.Request) (int, error) {
	sendCoinRequest, err := utils.Parse[api.SendCoinRequest](r)
	if err != nil {
		return http.StatusBadRequest, err
	}
	err = sc.storage.SendCoin(sendCoinRequest, utils.UserIdFromContext(r.Context()))
	if err != nil {
		if errors.Is(err, WrongData) {
			return http.StatusBadRequest, err
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func NewSendCoinHandler(storage SendCoinStorage) SendCoinHandler {
	return &SendCoinHandlerImpl{
		storage: storage,
	}
}
