package main

import (
	"github.com/victor-nach/go-testing/api"
)

func main() {
	router := api.Router()
	router.Run(":8080")

}
