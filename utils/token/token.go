package token

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Token struct {
	SecretKey     string
	TokenLifespan int
}

type tokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (t *Token) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(t.TokenLifespan))),
			Issuer:    "sharpic",
		},
	})

	return token.SignedString([]byte(t.SecretKey))
}

func (t *Token) TokenValid(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SecretKey), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) ExtractTokenUsername(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Username, nil
	}
	return "", nil
}
