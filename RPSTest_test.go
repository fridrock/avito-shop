package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

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
			_, authResponse, _, _ := MakeAuthRequest(testServer.URL, `{"username":"user1", "password":"user1"}`)
			_, _, _, _ = MakeGetInfoRequest(testServer.URL, authResponse.Token)
			_, _, _ = MakeSendCoinRequest(testServer.URL, authResponse.Token, `{"toUser":"user2", "amount":10}`)
			_, _, _ = MakeBuyRequest(testServer.URL, authResponse.Token, "wallet")

			rpsChannel <- 1

		}()
		time.Sleep(time.Second / time.Duration(requests))
	}

	wg.Wait()
	close(rpsChannel)
}

func TestRPS(t *testing.T) {
	requestRate := 1000
	testDuration := 10 * time.Second

	measureRPS(requestRate, testDuration)
}
