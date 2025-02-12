package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/db"
	"github.com/fridrock/avito-shop/handlers"
	"github.com/fridrock/avito-shop/storage"
	"github.com/fridrock/avito-shop/utils"
	"github.com/gorilla/mux"
)

func main() {
	conn := db.CreateConnection()
	defer conn.Close()
	userStorage := storage.NewUserStorage(conn)
	tokenService := auth.NewTokenService()
	authHandler := handlers.NewAuthHandler(userStorage, tokenService, utils.NewPasswordHasher())
	authManager := auth.NewAuthManager(tokenService)
	coinStorage := storage.NewCoinStorage(conn)
	sendCoinHandler := handlers.NewSendCoinHandler(coinStorage, userStorage)
	productStorage := storage.NewProductStorage(conn)
	infoStorage := storage.NewInfoStorage(conn)
	infoHandler := handlers.NewInfoHandler(infoStorage, userStorage)
	buyHandler := handlers.NewBuyHandler(productStorage, userStorage)
	router := mux.NewRouter()
	router.Handle("/api/auth", utils.HandleErrorMiddleware(authHandler.Auth)).Methods("POST")
	router.Handle("/api/sendCoin", utils.HandleErrorMiddleware(authManager.AuthMiddleware(sendCoinHandler.SendCoin))).Methods("POST")
	router.Handle("/api/buy/{item}", utils.HandleErrorMiddleware(authManager.AuthMiddleware(buyHandler.Buy))).Methods("GET")
	router.Handle("/api/info", utils.HandleErrorMiddleware(authManager.AuthMiddleware(infoHandler.GetInfo))).Methods("GET")
	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      router,
	}
	slog.Info("Starting server on port 8000")
	server.ListenAndServe()

}
