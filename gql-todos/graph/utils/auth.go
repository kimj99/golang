package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("test-secret-key")

func DecodeAPIKeys(key string) bool {
	return key == "blah"
}
func CreateJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	signedToken, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedToken
}
