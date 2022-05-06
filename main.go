package main

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

func validateAuthorizationToken(w http.ResponseWriter, r *http.Request) {
	AuthorizationHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(AuthorizationHeader, "Bearer ") {
		return
	}

	tokenString := strings.Split(AuthorizationHeader, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecretKey), nil
	})

	var statusCode = http.StatusUnauthorized
	var msg = ""

	if token.Valid {
		statusCode = http.StatusOK
		msg = "Good token"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			msg = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			msg = "Token expired"
		} else {
			msg = "Couldn't handle this token:"
		}
	} else {
		msg = "Couldn't handle this token:"
	}

	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(msg))

	if err != nil {
		log.Fatal("Write response error", err)
	}
}

func main() {
	InitArgs()
	log.Println("JWT-FILTER SERVER START SUCCESSFULLY")

	// Устанавливаем роутер
	http.HandleFunc("/", validateAuthorizationToken)

	// устанавливаем порт веб-сервера
	err := http.ListenAndServe(":"+JwtFilterPort, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
