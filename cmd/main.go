package main

import (
	"sarkor/test"
	"sarkor/test/pkg/handler"
	"sarkor/test/pkg/repository"
	"sarkor/test/pkg/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main(){
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("No configuration file found")
	}

	db, err := repository.NewSqliteDB(repository.Config{Path: viper.GetString("db")})
	if err != nil{
		logrus.Fatalf("Failed to initialise DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(test.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
		logrus.Fatalf("Error occurred while running server: %s", err.Error())
	}

	// rows, err := db.Query("SELECT id, name FROM user")
	// if err != nil {
    //     panic(err)
    // }
	// for rows.Next() {
    //     var id int
    //     var name string
    //     err = rows.Scan(&id, &name)
    //     if err != nil {
    //         panic(err)
    //     }
    //     fmt.Println(id, name)
    // }

	// if err := rows.Err(); err != nil {
	// 	logrus.Fatal(err)
	//   }
	// defer rows.Close()
	
	defer db.Close()
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	
	return viper.ReadInConfig()
}
