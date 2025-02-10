package auth

import (
	"fmt"
	"net/http"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/utils"
)

type AuthHandler interface {
	Auth(w http.ResponseWriter, r *http.Request) (int, error)
}

type AuthHandlerImpl struct {
	storage        AuthStorage
	tokenService   TokenService
	passwordHasher utils.PasswordHasher
}

func (ah *AuthHandlerImpl) Auth(w http.ResponseWriter, r *http.Request) (int, error) {
	authRequest, err := utils.Parse[api.AuthRequest](r)
	if err != nil {
		return 400, err
	}
	fmt.Println(authRequest)
	return 0, nil
}

func NewAuthHandler(storage AuthStorage, tokenService TokenService) AuthHandler {
	return &AuthHandlerImpl{
		storage:        storage,
		tokenService:   tokenService,
		passwordHasher: utils.NewPasswordHasher(),
	}
}
