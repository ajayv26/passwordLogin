package main

import (
	"fmt"
	"jwt/routes"
	"jwt/settings"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "root"
	PASSWORD = "secret"
	DBNAME   = "jwt"
)

func main() {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	dbClient, err := settings.ConnectPostgres(connString)
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()

	if dbClient == nil {
		log.Fatal("db connection failed")
	}
	RunServer()
}

func RunServer() {
	serverAddress := ":8080"
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes.Routes(router)

	log.Printf("server running on port %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatal(err)
	}
}
