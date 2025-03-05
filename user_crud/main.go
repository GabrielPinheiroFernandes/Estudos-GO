package main

import (
	"userCrud/controllers"
	"userCrud/repository"
)

func main() {
	

	repo:=&repository.LocalUserRepository{}
	// repo:=&repository.SqliteUserRepository{}
	controller := controllers.NewController(repo)
	controller.Run()
}
