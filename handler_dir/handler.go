package handler_dir

import (
	"RMS/database_dir"
	"RMS/model_dir"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model_dir.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var authorized model_dir.Credentials

	authorized, er := database_dir.GetPassword(credentials)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	if credentials.Password != authorized.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ExpiryTime := time.Now().Add(time.Minute * 30).Unix()
	Expires := time.Now().Add(time.Minute * 30)
	sessionToken := uuid.NewString()

	_, exErr := database_dir.InsertSession(sessionToken, authorized, Expires)
	if exErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(exErr)
		return
	}
	w.Header().Add("sessionToken", sessionToken)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var auth model_dir.Authentication

	auth, er = database_dir.GetPersonRole(authorized.Email)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	claims["email"] = authorized.Email
	claims["exp"] = ExpiryTime
	claims["id"] = auth.ID
	claims["role"] = auth.Role

	userTokenString, SignErr := token.SignedString(model_dir.JwtKey)
	if SignErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(SignErr)
		return
	}
	var userToken model_dir.Token
	userToken.Email = authorized.Email
	userToken.TokenString = userTokenString
	EncodeErr := json.NewEncoder(w).Encode(userToken)
	if EncodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(EncodeErr)
		return
	}

}
func Logout(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("sessionToken")
	_, execErr := database_dir.DelSession(sessionToken)
	if execErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(execErr)
		return
	}
	_, er := io.WriteString(w, "Successfully Logged out")
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
