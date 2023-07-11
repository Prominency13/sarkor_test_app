package main

import (
	"log"
	"sarkor/test"
	"sarkor/test/pkg/handler"
	"sarkor/test/pkg/repository"
	"sarkor/test/pkg/service"

	"github.com/spf13/viper"
)

func main(){
	if err := initConfig(); err != nil {
		log.Fatalf("No configuration file found")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(test.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
		log.Fatalf("Error ocured while running server: %s", err.Error())
	}

}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	
	return viper.ReadInConfig()
}
