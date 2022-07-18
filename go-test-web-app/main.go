package main

import (
	"go-test-web-app/api"
	"go-test-web-app/repository"
)

func main() {

	repository.Init()
	defer repository.Close()

	router := api.Init()
	router.Run(":8080")
}
