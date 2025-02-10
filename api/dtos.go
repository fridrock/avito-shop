package api

import "github.com/google/uuid"

type UserInfo struct {
	Id       uuid.UUID
	Username string
}

type MerchDto struct {
	Type      string `db:"type"`
	Quanitity int    `db:"quantity"`
}

type CoinHistoryDto struct {
	Username string
	Amount   int
}

type InfoResponse struct {
	Coins       int              `json:"coins"`
	Inventory   []MerchDto       `json:"inventory"`
	CoinHistory []CoinHistoryDto `json:"coinHistory"`
	Sent        []CoinHistoryDto `json:"sent"`
}

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type SendCoinRequest struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
}
