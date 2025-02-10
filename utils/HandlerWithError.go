package utils

import (
	"log/slog"
	"net/http"

	"github.com/fridrock/avito-shop/api"
)

// TODO get from users repo or even from itself repo
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type HandlerWithError func(w http.ResponseWriter, r *http.Request) (status int, err error)

func HandleErrorMiddleware(h HandlerWithError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		status, err := h(w, r)
		if err != nil {
			w.WriteHeader(status)
			slog.Error(err.Error())
			errorResponse := api.ErrorResponse{
				Errors: err.Error(),
			}
			WriteEncoded(w, errorResponse)
		}
	})
}
