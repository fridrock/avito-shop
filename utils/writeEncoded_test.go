package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fridrock/avito-shop/api"
)

func Test_writeEncodedTypeThatCannotSerialize(t *testing.T) {
	status, err := WriteEncoded(nil, make(chan int))
	if err == nil || status != http.StatusInternalServerError {
		t.Fatalf("encoded wrong type")
	}
}
func Test_writeEncoded(t *testing.T) {
	w := httptest.NewRecorder()
	status, err := WriteEncoded(w, api.InfoResponse{Coins: 100})
	if err != nil || status != http.StatusOK {
		t.Fatalf("encoding response error")
	}
}
