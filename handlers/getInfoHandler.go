package handlers

import (
	"net/http"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/storage"
	"github.com/fridrock/avito-shop/utils"
)

type InfoHandler interface {
	GetInfo(w http.ResponseWriter, r *http.Request) (int, error)
}

type getInfoHandler struct {
	infoStorage storage.InfoStorage
	userStorage storage.UserStorage
}

func (ih *getInfoHandler) GetInfo(w http.ResponseWriter, r *http.Request) (int, error) {
	userId := auth.UserFromContext(r.Context())
	user, err := ih.userStorage.GetUserById(userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	res, err := ih.infoStorage.GetInfoResponse(userId)
	res.Coins = user.Coins
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return utils.WriteEncoded(w, res)
}

func NewInfoHandler(is storage.InfoStorage, us storage.UserStorage) InfoHandler {
	return &getInfoHandler{
		infoStorage: is,
		userStorage: us,
	}
}
