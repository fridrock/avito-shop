package utils

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"testing"

	"github.com/fridrock/avito-shop/api"
)

// TODO Refactor
func TestAuthRequest(t *testing.T) {
	authRequest := `{"username":"someUsername","password":"somePassword"}`
	req, err := requestFromJson(authRequest)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	parsed, err := Parse[api.AuthRequest](req)
	if err != nil {
		t.Fatalf("Parsing error : %v", err)
	} else {
		slog.Info(fmt.Sprintf("%v", parsed))
	}
	if parsed.Password != "somePassword" || parsed.Username != "someUsername" {
		t.Fatalf("Value of parsed dto is incorrect: %v", parsed)
	}

}

func TestAuthRequestInvalid(t *testing.T) {
	authRequest := `{"username":"someUsername"}`
	req, err := requestFromJson(authRequest)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	parsed, err := Parse[api.AuthRequest](req)
	if err == nil {
		t.Fatalf("Parsing error : %v, %v", parsed, err)
	} else {
		slog.Error(fmt.Sprintf("%v", err))
	}

}

func TestInvalidEncoding(t *testing.T) {
	authRequest := `{"username:true"}`
	req, err := requestFromJson(authRequest)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	_, err = Parse[api.AuthRequest](req)
	if err == nil {
		t.Fatalf("Read wrong value type")
	}
}

func TestSendCoin(t *testing.T) {
	sendCoin := `{"toUser":"someUsername","amount":12}`
	req, err := requestFromJson(sendCoin)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	parsed, err := Parse[api.SendCoinRequest](req)
	if err != nil {
		t.Fatalf("Parsing error : %v", err)
	} else {
		slog.Info(fmt.Sprintf("%v", parsed))
	}
}

func requestFromJson(jsonString string) (*http.Request, error) {
	req, err := http.NewRequest("POST", "http://example.com", bytes.NewBuffer([]byte(jsonString)))
	if err != nil {
		return nil, err
	}
	return req, nil
}
