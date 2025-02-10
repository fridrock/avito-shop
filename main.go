package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/coinsend"
	"github.com/fridrock/avito-shop/db"
	"github.com/fridrock/avito-shop/utils"
	"github.com/gorilla/mux"
)

func main() {
	conn := db.CreateConnection()
	defer conn.Close()
	authStorage := auth.NewAuthStorage(conn)
	tokenService := auth.NewTokenService()
	authHandler := auth.NewAuthHandler(authStorage, tokenService)
	authManager := auth.NewAuthManager(tokenService)
	sendCoinStorage := coinsend.NewSendCoinStorage(conn)

	sendCoinHandler := coinsend.NewSendCoinHandler(sendCoinStorage)
	router := mux.NewRouter()
	router.Handle("/api/auth", utils.HandleErrorMiddleware(authHandler.Auth)).Methods("POST")
	router.Handle("/api/sendCoin", utils.HandleErrorMiddleware(authManager.AuthMiddleware(sendCoinHandler.SendCoin))).Methods("POST")
	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      router,
	}
	slog.Info("Starting server on port 8000")
	server.ListenAndServe()

}
