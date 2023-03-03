package main

import (
	"log"

	"github.com/monirz/goserve/http"

	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//get the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := http.NewServer()

	s.Run()
}
