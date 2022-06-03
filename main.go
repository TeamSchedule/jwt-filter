package main

import (
	"jwt-filter/server"
	"jwt-filter/server/conf"
	"log"
)

func main() {
	r := server.SetupServer()

	err := r.Run(":" + conf.JwtFilterPort)
	if err != nil {
		log.Fatalln("JWT_FILTER service crashed!")
		return
	}
}
