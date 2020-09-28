package main

import (
	"log"

	"github.com/andersondalmina/golang-sockets/persist"

	"github.com/andersondalmina/golang-sockets/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	persist.CreateDatabase()

	cmd.Execute()
}
