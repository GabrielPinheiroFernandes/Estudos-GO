package main

import (
	"userCrud/controllers"
	"userCrud/repository"
)

func main() {
	repo:=repository.LocalUserRepository{}
	controller := controllers.NewController(repo)
	controller.Run()
}
