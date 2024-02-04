package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	"github.com/oprimogus/cardapiogo/internal/services/oauth2"
)

const TimeExpireInHour = 1

func GenerateJWTForValidation() (string, error) {
	key := os.Getenv("JWT_SECRET")
	expireIn := time.Now().Add(time.Hour * time.Duration(TimeExpireInHour)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": os.Getenv("JWT_EMISSOR"),
		"exp": expireIn,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", errors.InternalServerError(err.Error())
	}
	return s, err
}

func GenerateJWTWithClaims(user *user.User) (string, error) {
	key := os.Getenv("JWT_SECRET")
	expireIn := time.Now().Add(time.Hour * time.Duration(TimeExpireInHour)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      os.Getenv("JWT_EMISSOR"),
		"sub":      user.ID,
		"role":     user.Role,
		"provider": user.AccountProvider,
		"exp":      expireIn,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", errors.InternalServerError(err.Error())
	}
	return s, err
}

func ValidateStateToken(stateToken string) (bool, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(stateToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, errors.BadRequest(fmt.Sprintf("signature method not expected: %v", token.Header["alg"]))
		}
		if !token.Valid {
			return false, errors.BadRequest("Token not valid.")
		}
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
			return "", errors.InternalServerError(err.Error())
		}
		return "", errors.New(dbErr.Status, dbErr.ErrorMessage)
	}
	isSamePassword := service.IsValidPassword(loginParams.Password, existUser.Password)
	if isSamePassword {
		jwt, err := GenerateJWTWithClaims(existUser)
		if err != nil {
			return "", errors.InternalServerError(err.Error())
		}
		return jwt, nil
	}
	return "", errors.New(http.StatusBadRequest, "Invalid Password.")
}

func createUserInOauth(ctx context.Context, s *user.Service, u *user.CreateUserWithOAuthParams) (*user.User, error) {
	err := s.CreateUserWithOAuth(ctx, *u)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if !ok {
			return nil, errors.InternalServerError(err.Error())
		}
		return nil, dbErr
	}
	createdUser, err := s.GetUserByEmail(ctx, u.Email)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if !ok {
			return nil, errors.InternalServerError(err.Error())
		}
		return nil, dbErr
	}
	return createdUser, nil
}

func LoginWithOauth(ctx context.Context, s *user.Service, userData *oauth2.GoogleUserInfo) (string, error) {
	existUser, err := s.GetUserByEmail(ctx, userData.Email)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if !ok {
			return "", errors.InternalServerError(err.Error())
		}
		if dbErr.ErrorMessage == errors.NOT_FOUND_RECORD {
			u := user.CreateUserWithOAuthParams{
				Email:           userData.Email,
				Role:            types.UserRoleConsumer,
				AccountProvider: types.AccountProviderGoogle,
			}

			createdUser, err := createUserInOauth(ctx, s, &u)
			if err != nil {
				return "", err
			}

			jwt, err := GenerateJWTWithClaims(createdUser)
			if err != nil {
				return "", err
			}
			return jwt, nil
		}
	}
	jwt, err := GenerateJWTWithClaims(existUser)
	if err != nil {
		return "", err
	}
	return jwt, nil

}
