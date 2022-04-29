package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strings"
)

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
var JwtFilterPort = os.Getenv("JWT_FILTER_PORT")

func sayHello(w http.ResponseWriter, r *http.Request) {
	AuthorizationHeader := r.Header.Get("Authorization")
	fmt.Println(AuthorizationHeader)
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

	fmt.Println(statusCode)
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, msg)
}

func main() {
	fmt.Println("JWT FILTER - START SERVER")
	fmt.Println("JWT_SECRET_KEY = " + JwtSecretKey)
	fmt.Println("JWT_FILTER_PORT = " + JwtFilterPort)

	// Устанавливаем роутер
	http.HandleFunc("/", sayHello)

	// устанавливаем порт веб-сервера
	err := http.ListenAndServe(":"+JwtFilterPort, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
