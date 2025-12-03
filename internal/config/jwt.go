package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", id),
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	})
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRETKEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRandomString() (string, error) {
	str := make([]byte, 32)
	_, err := rand.Read(str)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(str), nil
}

func JwtMiddleware() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRETKEY")), nil)

	return tokenAuth
}
