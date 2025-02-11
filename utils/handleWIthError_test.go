package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handleErrorMiddlewareWithError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) (int, error) {
		return http.StatusInternalServerError, fmt.Errorf("some error")
	}
	w := httptest.NewRecorder()
	handlerWithMiddleware := HandleErrorMiddleware(handler)
	handlerWithMiddleware.ServeHTTP(w, nil)
	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		body, _ := io.ReadAll(res.Body)
		t.Fatalf("doesn't receive error message :%v", string(body))
	}
}

func Test_handleErrorMiddlewareWithoutError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) (int, error) {
		w.Write([]byte("success"))
		return http.StatusOK, nil
	}
	w := httptest.NewRecorder()
	handlerWithMiddleware := HandleErrorMiddleware(handler)
	handlerWithMiddleware.ServeHTTP(w, nil)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		t.Fatalf("Wrong status code :%v", string(body))
	}

}
