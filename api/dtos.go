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

type FromUserDto struct {
	FromUser string `db:"fromUser"`
	Amount   int    `db:"amount"`
}
type ToUserDto struct {
	ToUser string `db:"toUser"`
	Amount int    `db:"amount"`
}

type InfoResponse struct {
	Coins       int           `json:"coins"`
	Inventory   []MerchDto    `json:"inventory"`
	CoinHistory []FromUserDto `json:"coinHistory"`
	Sent        []ToUserDto   `json:"sent"`
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
