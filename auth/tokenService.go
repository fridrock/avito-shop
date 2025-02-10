package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fridrock/avito-shop/api"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenService interface {
	GenerateToken(api.AuthRequest, uuid.UUID) (api.AuthResponse, error)
	ValidateToken(string) (api.UserInfo, error)
}

type TokenServiceImpl struct {
	SECRET_KEY []byte
}

func (ts *TokenServiceImpl) GenerateToken(authRequest api.AuthRequest, userId uuid.UUID) (api.AuthResponse, error) {
	var dto api.AuthResponse
	accessTokenString, err := ts.generateAccess(authRequest, userId)
	if err != nil {
		return dto, err
	}
	dto.Token = accessTokenString
	return dto, nil
}

func (ts *TokenServiceImpl) generateAccess(user api.AuthRequest, userId uuid.UUID) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userId,
		"username": user.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	return accessToken.SignedString(ts.SECRET_KEY)
}

func (ts *TokenServiceImpl) ValidateToken(tokenString string) (api.UserInfo, error) {
	var dto api.UserInfo
	parsed, err := ts.parseToken(tokenString)
	if err != nil {
		return dto, err
	}
	if time.Now().After(parsed.exp.Time) {
		return dto, fmt.Errorf("expired token %v", parsed.exp.Time)
	}
	dto.Id = parsed.id
	dto.Username = parsed.username
	return dto, nil
}

type tokenParsed struct {
	id       uuid.UUID
	username string
	exp      *jwt.NumericDate
}

func (ts *TokenServiceImpl) parseToken(tokenString string) (tokenParsed, error) {
	var dto tokenParsed
	tokenObj, err := ts.checkSigning(tokenString)
	if err != nil {
		return dto, err
	}
	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok {
		id, err := uuid.Parse(fmt.Sprintf("%v", claims["id"]))
		if err != nil {
			return dto, err
		}
		dto.id = id
		username := fmt.Sprintf("%v", claims["username"])
		if username == "" {
			return dto, fmt.Errorf("no username in token")
		}
		dto.username = username
		exp, err := claims.GetExpirationTime()
		if err != nil {
			return dto, err
		}
		dto.exp = exp
	}
	return dto, nil
}

func (ts *TokenServiceImpl) checkSigning(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ts.SECRET_KEY, nil
	})
}

func NewTokenService() TokenService {
	varName := "SECRET_KEY"
	secret, exists := os.LookupEnv(varName)
	if !exists {
		log.Fatalf("Can't load env variable: %v", varName)
	}

	return &TokenServiceImpl{
		SECRET_KEY: []byte(secret),
	}
}
