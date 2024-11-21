package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	Id  string
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(id string, email string, secretKey string) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtToken{
		Id: id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	},
)
	return claim.SignedString([]byte(secretKey))
}


func VerifyToken(token string, secretKey string) (*JwtToken, error) {
	claims := &JwtToken{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return claims, nil
}
