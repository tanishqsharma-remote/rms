package middleware_dir

import (
	"RMS/database_dir"
	"RMS/model_dir"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

//todo what is jwt
// todo why over session use jwt

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var validSession model_dir.Session
		sessionToken := r.Header.Get("sessionToken")
		validSession, err := database_dir.GetSession(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if validSession.Expiry.Before(time.Now()) {
			_, execErr := database_dir.DelSession(sessionToken)
			if execErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(execErr)
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userToken := r.Header.Get("jwtToken")
		checkToken, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token. ")
			}
			return model_dir.JwtKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		claims, ok := checkToken.Claims.(jwt.MapClaims)
		if ok && checkToken.Valid {
			ctx := context.WithValue(r.Context(), "email", claims)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		//todo user could have multiple role
		claims := r.Context().Value("email").(jwt.MapClaims)
		if claims["role"] != "admin" {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(writer, r)

	})
}
func SubAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("email").(jwt.MapClaims)
		if claims["role"] != "sub-admin" {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(writer, r)

	})
}
