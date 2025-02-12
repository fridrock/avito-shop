package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fridrock/avito-shop/api"
)

func MakeAuthRequest(serverUrl string, body string) (int, api.AuthResponse, api.ErrorResponse, error) {
	var authResponse api.AuthResponse
	var errorResponse api.ErrorResponse
	resp, err := http.Post(serverUrl+"/api/auth", "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return -1, authResponse, errorResponse, err
	}
	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, authResponse, errorResponse, fmt.Errorf("error getting body")
		}
		err = json.Unmarshal(responseBody, &errorResponse)
		if err != nil {
			return resp.StatusCode, authResponse, errorResponse, fmt.Errorf("error unmarshalling error response")
		}
		return resp.StatusCode, authResponse, errorResponse, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, authResponse, errorResponse, fmt.Errorf("error getting body")
	}
	err = json.Unmarshal(responseBody, &authResponse)
	return resp.StatusCode, authResponse, errorResponse, err
}

func MakeSendCoinRequest(serverUrl string, token string, body string) (int, api.ErrorResponse, error) {
	var errorResponse api.ErrorResponse
	req, err := http.NewRequest("POST", serverUrl+"/api/sendCoin", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return -1, errorResponse, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, errorResponse, fmt.Errorf("error getting body")
		}
		err = json.Unmarshal(responseBody, &errorResponse)
		if err != nil {
			return resp.StatusCode, errorResponse, fmt.Errorf("error unmarshalling error response")
		}
	}
	return resp.StatusCode, errorResponse, err
}

func MakeBuyRequest(serverUrl string, token string, itemName string) (int, api.ErrorResponse, error) {
	var errorResponse api.ErrorResponse
	req, err := http.NewRequest("GET", serverUrl+"/api/buy/"+itemName, nil)
	if err != nil {
		return -1, errorResponse, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, errorResponse, fmt.Errorf("error getting body")
		}
		err = json.Unmarshal(responseBody, &errorResponse)
		if err != nil {
			return resp.StatusCode, errorResponse, fmt.Errorf("error unmarshalling error response")
		}
	}
	return resp.StatusCode, errorResponse, err
}

func MakeGetInfoRequest(serverUrl string, token string) (int, api.InfoResponse, api.ErrorResponse, error) {
	var infoResponse api.InfoResponse
	var errorResponse api.ErrorResponse
	req, err := http.NewRequest("GET", serverUrl+"/api/info", nil)
	if err != nil {
		return -1, infoResponse, errorResponse, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp.StatusCode, infoResponse, errorResponse, err
	}
	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, infoResponse, errorResponse, fmt.Errorf("error getting body")
		}
		err = json.Unmarshal(responseBody, &errorResponse)
		if err != nil {
			return resp.StatusCode, infoResponse, errorResponse, fmt.Errorf("error unmarshalling error response")
		}
		return resp.StatusCode, infoResponse, errorResponse, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, infoResponse, errorResponse, fmt.Errorf("error getting body")
	}
	err = json.Unmarshal(responseBody, &infoResponse)
	return resp.StatusCode, infoResponse, errorResponse, err
}
