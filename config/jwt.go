package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("sakura endo")

type JwtClaim struct {
	Username string
	jwt.RegisteredClaims
}
