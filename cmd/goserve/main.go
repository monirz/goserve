package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/monirz/goserve/config"
	"github.com/monirz/goserve/handlers"

	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//get the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Println(dsn)
		panic(err)
	}

	s := handlers.NewServer(db)
	s.Config = cfg

	s.Run()
}
