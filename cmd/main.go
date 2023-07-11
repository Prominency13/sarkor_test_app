package main

import (
	"log"
	"sarkor/test"
)

func main(){
	server := new(test.Server)
	if err := server.Run("8080"); err != nil{
		log.Fatalf("Error ocured while running server: %s", err.Error())
	}

}
