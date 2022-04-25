package authenticator

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Name     string `json:"name"`
}
