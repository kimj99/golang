package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("test-secret-key")

func DecodeAPIKeys(key string) bool {
	return key == "blah"
}
func CreateJWT(name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = name
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	signedToken, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return signedToken, nil
}

func ParseToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		name := claims["username"].(string)
		return name, nil
	} else {
		return "", err
	}
}
