package getinfo

import "net/http"

type GetInfoHandler interface {
	GetInfo(w http.ResponseWriter, r *http.Request) (int, error)
}

type GetInfoHandlerImpl struct {
	storage GetInfoStorage
}

func (gi *GetInfoHandlerImpl) GetInfo(w http.ResponseWriter, r *http.Request) (int, error) {
	return 0, nil
}

func NewGetInfoHandler(storage GetInfoStorage) GetInfoHandler {
	return &GetInfoHandlerImpl{
		storage: storage,
	}
}
