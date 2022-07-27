package server

import (
	"context"
	"finalproject/database"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func SecureMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/login") ||
			strings.Contains(r.URL.Path, "/register") {
			next.ServeHTTP(w, r)
			return
		}

		// Check to see if this request can go thru
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			WriteJsonResp(w, ErrorForbidden, "FORBIDDEN")
			return
		}

		splitToken := strings.Split(auth, "Bearer ")
		if len(splitToken) != 2 {
			WriteJsonResp(w, ErrorForbidden, "FORBIDDEN")
			return
		}

		accessToken := splitToken[1]
		if len(accessToken) == 0 {
			WriteJsonResp(w, ErrorForbidden, "FORBIDDEN")
			return
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(Config.SecretKey), nil
		})
		if err != nil {
			e, ok := err.(*jwt.ValidationError)
			if !ok || ok && e.Errors&jwt.ValidationErrorIssuedAt == 0 { // Don't report error that token used before issued.
				WriteJsonResp(w, ErrorBadRequest, "BAD_REQUEST")
				return
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok { //|| !token.Valid {
			WriteJsonResp(w, ErrorBadRequest, "BAD_REQUEST")
			return
		}

		uid := claims["uid"].(float64)

		userID := int64(uid)

		l, err := database.SqlDatabase.GetUserByID(context.Background(), userID)
		if err != nil {
			WriteJsonResp(w, ErrorDataHandleError, err)
			return
		}
		//Set logonuser
		LogonUser = l
		fmt.Println(uid)

		next.ServeHTTP(w, r)
	})
}
