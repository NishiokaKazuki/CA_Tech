package auth

import (
	"server/model/tables"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(user tables.AppUsers) string {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = user.Name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, _ := t.SignedString([]byte(user.Name))

	return strings.Replace(token, "_", "", -1)
}
