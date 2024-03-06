package usecases

import (
	"time"

	authError "segmentation/auth/errors"

	"github.com/dgrijalva/jwt-go"
)

type TokenUsecase interface {
	GenerateToken(id *uint, username *string) (*string, error)
	ParseToken(token *string) (*uint, error)
}

type tokenUsecaseImpl struct {
	secretKey string
}

func NewTokenUsecaseImpl(secretKey string) TokenUsecase {
	return &tokenUsecaseImpl{
		secretKey: secretKey,
	}
}

func (u *tokenUsecaseImpl) GenerateToken(id *uint, username *string) (*string, error) {

	key := []byte(u.secretKey)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	clams := token.Claims.(jwt.MapClaims)
	clams["authorized"] = true
	clams["dentist_id"] = *id
	clams["username"] = *username
	clams["exp"] = time.Now().Add(time.Hour * 2).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (u *tokenUsecaseImpl) ParseToken(token *string) (*uint, error) {
	key := []byte(u.secretKey)
	// Parse the token
	t, err := jwt.Parse(*token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &authError.TokenNotAuthorizedError{}
		}
		return key, nil
	})
	if err != nil {
		return nil, &authError.ServerInternalError{Err: err}
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, &authError.TokenNotAuthorizedError{}
	}
	dentistID := uint(claims["dentist_id"].(float64))
	return &dentistID, nil
}
