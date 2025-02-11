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

type InfoHandlerImpl struct {
	infoStorage storage.InfoStorage
}

func (ih *InfoHandlerImpl) GetInfo(w http.ResponseWriter, r *http.Request) (int, error) {
	userId := auth.UserFromContext(r.Context())
	res, err := ih.infoStorage.GetInfoResponse(userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return utils.WriteEncoded(w, res)
}

func NewInfoHandler(is storage.InfoStorage) InfoHandler {
	return &InfoHandlerImpl{
		infoStorage: is,
	}
}
