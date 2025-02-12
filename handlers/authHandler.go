package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/auth"
	"github.com/fridrock/avito-shop/storage"
	"github.com/fridrock/avito-shop/utils"
	"github.com/google/uuid"
)

type AuthHandler interface {
	Auth(w http.ResponseWriter, r *http.Request) (int, error)
}

type authHandler struct {
	storage        storage.UserStorage
	tokenService   auth.TokenService
	passwordHasher utils.PasswordHasher
}

func (ah *authHandler) Auth(w http.ResponseWriter, r *http.Request) (int, error) {
	authRequest, err := utils.Parse[api.AuthRequest](r)
	if err != nil {
		return http.StatusBadRequest, err
	}
	user, err := ah.storage.FindUserByUsername(authRequest.Username)
	// Авторизация
	if err == nil {
		if ah.passwordHasher.CheckPassword(authRequest.Password, user.HashedPassword) {
			return ah.sendToken(w, authRequest, user.Id)
		} else {
			return http.StatusUnauthorized, fmt.Errorf("wrong password")
		}
	} else {
		if errors.Is(err, sql.ErrNoRows) {
			hashedPassword, err := ah.passwordHasher.HashPassword(authRequest.Password)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			user = storage.User{
				Username:       authRequest.Username,
				HashedPassword: hashedPassword,
			}
			user.Id, err = ah.storage.SaveUser(user)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			return ah.sendToken(w, authRequest, user.Id)
		}
		return http.StatusInternalServerError, err
	}
}

func (ah *authHandler) sendToken(w http.ResponseWriter, authRequest api.AuthRequest, userId uuid.UUID) (int, error) {
	authResponse, err := ah.tokenService.GenerateToken(authRequest, userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return utils.WriteEncoded(w, authResponse)
}
func NewAuthHandler(st storage.UserStorage, tokenService auth.TokenService, ph utils.PasswordHasher) AuthHandler {
	return &authHandler{
		storage:        st,
		tokenService:   tokenService,
		passwordHasher: ph,
	}
}
