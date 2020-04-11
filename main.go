package main

import (
	"os"

	"github.com/joho/godotenv"
	
	"github.com/victor-nach/user-management-go/api/routes"
)

func main() {
	_ = godotenv.Load()
	router := routes.Router()
	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = "8080"
	}
	router.Run(":" + PORT)
}
