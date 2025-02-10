package utils

import (
	"encoding/json"
	"net/http"
)

func WriteEncoded(w http.ResponseWriter, data any) (int, error) {
	responseText, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	w.Write([]byte(responseText))
	return http.StatusOK, err
}
