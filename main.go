package main

import (
	"fmt"
	"log"
	"os"
	"verihubs-cryptocurrencies/internal/pkg/database"
	"verihubs-cryptocurrencies/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error read env file with err: %s", err)
	}

	// if err := database.ConnectDB(); err != nil {
	// 	log.Fatalf("Can't connect to database with err: %s", err)
	// }

	db := database.ConnectDB()

	router := router.NewRoutes(db)
	router.LoadHandlers()
	router.Echo.Logger.Fatal(router.Echo.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}