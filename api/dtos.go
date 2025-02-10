package api

import "github.com/google/uuid"

type UserInfo struct {
	Id       uuid.UUID
	Username string
}

// TODO make validation rules
type MerchDto struct {
	Type      string
	Quanitity int
}

type CoinHistoryDto struct {
	Username string
	Amount   int
}

type InfoResponse struct {
	Coins       int              `json:"coins" validate:"required"`
	Inventory   []MerchDto       `json:"inventory" validate: "required"`
	CoinHistory []CoinHistoryDto `json:"coinHistory" validate: "required"`
	Sent        []CoinHistoryDto `json:"sent" validate:"required"`
}

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SendCoinRequest struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
}
