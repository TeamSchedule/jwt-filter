package main

import (
	"log"
	"os"
)

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
var JwtFilterPort = os.Getenv("JWT_FILTER_PORT")

func InitArgs() {
	if JwtSecretKey == "" {
		log.Fatal("environment variable `JWT_SECRET_KEY` must be specified and must be non empty value")
	}

	if JwtFilterPort == "" {
		log.Fatal("environment variable `JWT_FILTER_PORT` must be specified and must be non empty value")
	}
}
