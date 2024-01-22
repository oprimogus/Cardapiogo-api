package auth

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

const TimeExpireInHour = 1

func GenerateJWTOAuthWithClaims(expireInHours int, provider string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	expireIn := time.Now().Add(time.Hour * time.Duration(expireInHours)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"provider": provider,
		"exp":      expireIn,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", errors.NewErrorResponse(500, err.Error())
	}
	return s, err
}

func GenerateJWTWithClaims(user *user.User) (string, error) {
	key := os.Getenv("JWT_SECRET")
	expireIn := time.Now().Add(time.Hour * time.Duration(TimeExpireInHour)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role": user.Role,
		"exp": expireIn,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", errors.NewErrorResponse(500, err.Error())
	}
	return s, err
}

func ValidateStateToken(stateToken string) (bool, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(stateToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func Login(ctx context.Context, service *user.Service, loginParams *user.Login) (string, error) {
	existUser, err := service.GetUserByEmail(ctx, loginParams.Email)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if !ok {
			return "", errors.NewErrorResponse(http.StatusInternalServerError, err.Error())
		}
		return "", errors.NewErrorResponse(dbErr.Status, dbErr.ErrorMessage)
	}
	isSamePassword := service.IsValidPassword(loginParams.Password, existUser.Password)
	if isSamePassword {
		jwt, err := GenerateJWTWithClaims(existUser)
		if err != nil {
			return "", errors.NewErrorResponse(http.StatusInternalServerError, err.Error())
		}
		return jwt, nil
	}
	return "", errors.NewErrorResponse(http.StatusBadRequest, "Invalid Password.")
}