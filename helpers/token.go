package helpers

import (
	"errors"
	"go-gin-sosmed/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("mysecretkey")

type JWTClaims struct {
  ID int `json:"id"`
  jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
  claims := JWTClaims{
    user.ID,
    jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
      IssuedAt: jwt.NewNumericDate(time.Now()),
      NotBefore: jwt.NewNumericDate(time.Now()),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  ss, err := token.SignedString(mySigningKey)

  return ss, err
}

func ValidateToken(tokenString string) (*int, error) {
  token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
    return mySigningKey, nil
  })

  if err != nil {
    if err == jwt.ErrSignatureInvalid {
      return nil, errors.New("signature is invalid")
    }

    return nil, errors.New("your token was expired")
  }

  claims, ok := token.Claims.(*JWTClaims)
  if !ok || !token.Valid {
    return nil, errors.New("invalid token")
  }

  return &claims.ID, nil
}