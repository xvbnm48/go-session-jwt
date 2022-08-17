package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xvbnm48/go-session-jwt/config"
	"github.com/xvbnm48/go-session-jwt/helper"
)

func JWTMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				response := map[string]string{
					"message": "Unauthorized",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		// take a token value from the cookie
		tokenString := c.Value
		// claim the token
		claims := &config.JwtClaim{}

		// parse the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{
					"message": "Unauthorized",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{
					"message": "Token Expired",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{
					"message": "Unauthorized",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		if !token.Valid {
			response := map[string]string{
				"message": "Unauthorized",
			}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
