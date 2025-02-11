package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/fridrock/avito-shop/api"
	"github.com/fridrock/avito-shop/utils"
	"github.com/google/uuid"
)

type UserContextKey string

const Key UserContextKey = "user"

func UserFromContext(ctx context.Context) uuid.UUID {
	return ctx.Value(Key).(uuid.UUID)
}

type AuthManager interface {
	AuthMiddleware(h utils.HandlerWithError) utils.HandlerWithError
}
type AuthManagerImpl struct {
	tokenService TokenService
}

func NewAuthManager(tokenService TokenService) AuthManager {
	return &AuthManagerImpl{
		tokenService: tokenService,
	}
}

func (am AuthManagerImpl) getUserFromToken(r *http.Request) (api.UserInfo, error) {
	var user api.UserInfo
	// Извлечение заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return user, fmt.Errorf("empty auth header")
	}

	// Проверка и разбиение заголовка
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return user, fmt.Errorf("wrong format of token")
	}

	user, err := am.tokenService.ValidateToken(parts[1])
	if err != nil {
		return user, fmt.Errorf("token invalidated")
	}
	return user, nil
}

func (am AuthManagerImpl) AuthMiddleware(h utils.HandlerWithError) utils.HandlerWithError {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		user, err := am.getUserFromToken(r)
		if err != nil {
			slog.Debug(err.Error())
			return http.StatusUnauthorized, err
		}
		ctx := context.WithValue(r.Context(), Key, user.Id)
		return h(w, r.WithContext(ctx))
	}
}
