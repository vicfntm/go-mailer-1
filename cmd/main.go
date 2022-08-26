package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gomail "github.com/vcfntm/go-mailer-1"
	"github.com/vcfntm/go-mailer-1/src/handlers"
	"github.com/vcfntm/go-mailer-1/src/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env file loading failed: %s", err.Error())
	}

	services := services.NewService()
	srv := new(gomail.Server)
	handlers := handlers.NewHandler(services)

	if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while server running %s", err.Error())
	}

	fmt.Printf("run on port: %s", os.Getenv("SERVER_PORT"))
}
