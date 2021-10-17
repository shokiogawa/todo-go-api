package presentation

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"os"
	"time"
)

type Auth struct {
}

func (auth *Auth) GetToken(publicId string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = publicId
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err = token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return
	}
	return
}

func (auth *Auth) Parse(tokenString string) (publicUserId string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.Errorf("invalid token")
		}
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	publicUserId, ok = claims["sub"].(string)
	if !ok {
		err = fmt.Errorf("not found %s in %s", "sub", tokenString)
	}

	return
}
