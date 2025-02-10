package coinsend

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/storage"
	"github.com/fridrock/avito-shop/utils"
)

type SendCoinHandler interface {
	SendCoin(w http.ResponseWriter, r *http.Request) (int, error)
}

type SendCoinHandlerImpl struct {
	coinStorage storage.CoinStorage
	userStorage storage.UserStorage
}

func (sc *SendCoinHandlerImpl) SendCoin(w http.ResponseWriter, r *http.Request) (int, error) {
	sendCoinRequest, err := utils.Parse[api.SendCoinRequest](r)
	if err != nil {
		return http.StatusBadRequest, err
	}
	//get user balance
	curUserId := auth.UserFromContext(r.Context())
	hasEnoughCoins := sc.userStorage.CheckEnoughCoins(sendCoinRequest.Amount, curUserId)
	//user with toUser username exist
	toUser, err := sc.userStorage.FindUserByUsername(sendCoinRequest.ToUser)
	resError := err
	if !hasEnoughCoins {
		resError = errors.Join(resError, fmt.Errorf("doesn't have such amount of coins of account"))
	}
	if curUserId.String() == toUser.Id.String() {
		resError = errors.Join(resError, fmt.Errorf("can't send coin to yourself"))
	}
	if resError != nil {
		return http.StatusBadRequest, resError
	}

	err = sc.coinStorage.SendCoin(sendCoinRequest.Amount, auth.UserFromContext(r.Context()), toUser.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func NewSendCoinHandler(cs storage.CoinStorage, us storage.UserStorage) SendCoinHandler {
	return &SendCoinHandlerImpl{
		coinStorage: cs,
		userStorage: us,
	}
}
