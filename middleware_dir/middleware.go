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

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database_dir.DBconnect()

		sessionToken := r.Header.Get("sessionToken")
		rows, err := database_dir.GetSession(db, sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var validSession model_dir.Session
		for rows.Next() {
			ScanErr := rows.Scan(&validSession.Email, &validSession.Expiry)
			if ScanErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(ScanErr)
				return
			}
		}
		if validSession.Expiry.Before(time.Now()) {
			_, execErr := database_dir.DelSession(db, sessionToken)
			if execErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(execErr)
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//var userToken model_dir.Token
		/*DecodeErr := json.NewDecoder(r.Body).Decode(&userToken)
		if DecodeErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(DecodeErr)

			return
		}*/
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
