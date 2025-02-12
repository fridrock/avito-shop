package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/db"
	"github.com/fridrock/avito-shop/handlers"
	"github.com/fridrock/avito-shop/storage"
	"github.com/fridrock/avito-shop/utils"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	conn := db.CreateConnection()
	defer conn.Close()
	router := setupRouter(conn)
	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      router,
	}
	slog.Info("Starting server on port 8000")
	go func() {
		TestRPS()
	}()
	defer server.Close()
	server.ListenAndServe()
}

func setupRouter(conn *sqlx.DB) *mux.Router {
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
	return router
}

func measureRPS(requests int, duration time.Duration) {
	var wg sync.WaitGroup
	rpsChannel := make(chan int)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		var total int
		for {
			select {
			case <-rpsChannel:
				total += 4
			case <-ticker.C:
				fmt.Printf("Current RPS: %d\n", total)
				total = 0
			}
		}
	}()

	startTime := time.Now()
	for time.Since(startTime) < duration {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, authResponse, _, _ := MakeAuthRequest("http://localhost:8000", `{"username":"user1", "password":"user1"}`)
			_, _, _, _ = MakeGetInfoRequest("http://localhost:8000", authResponse.Token)
			_, _, _ = MakeSendCoinRequest("http://localhost:8000", authResponse.Token, `{"toUser":"user2", "amount":10}`)
			_, _, _ = MakeBuyRequest("http://localhost:8000", authResponse.Token, "wallet")

			rpsChannel <- 1

		}()
		time.Sleep(time.Second / time.Duration(requests))
	}

	wg.Wait()
	close(rpsChannel)
}

func TestRPS() {
	requestRate := 1000
	testDuration := 20 * time.Second

	measureRPS(requestRate, testDuration)
}
