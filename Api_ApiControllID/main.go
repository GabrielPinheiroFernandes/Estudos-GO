package main

import (
	"APIControlID/controllers"
	crudapi "APIControlID/device_handler/crud_api"
	"log"
)

func main() {
	crudApi, err := crudapi.NewControlIdCrudApi()
	if err != nil {
		log.Fatalf("Erro ao criar controller")
	}
	controller := controllers.NewController(crudApi)
	controller.Inicialize()
}
