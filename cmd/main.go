package main

import (
	"log"
	"sarkor/test/pkg/handler"
	"sarkor/test"
)

func main(){
	handlers := new(handler.UserHandler)
	server := new(test.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil{
		log.Fatalf("Error ocured while running server: %s", err.Error())
	}

}
