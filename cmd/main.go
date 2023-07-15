package main

import (
	"sarkor/test"
	"sarkor/test/pkg/handler"
	"sarkor/test/pkg/repository"
	"sarkor/test/pkg/service"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main(){
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("No configuration file found")
	}

	// db, err := repository.NewSqliteDB(repository.Config{Path: viper.GetString("db.path")})
	// if err != nil{
	// 	logrus.Fatalf("Failed to initialise DB: %s", err.Error())
	// }

	db, err := sqlx.Open("sqlite3", "../database.db")
    if err != nil {
		logrus.Fatalf("Error occurred while connecting to database:", err.Error())
    }

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(test.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
		logrus.Fatalf("Error occurred while running server: %s", err.Error())
	}


	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY, login TEXT, password TEXT, name TEXT, age TEXT);")
    if err != nil {
		logrus.Fatalf("Error occurred while processing SQL query", err.Error())
    }
	
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS phone(id INTEGER PRIMARY KEY, phone TEXT, description TEXT, is_fax TINYINT, user_id INTEGER, FOREIGN KEY(user_id) REFERENCES user(id));")
    if err != nil {
		logrus.Fatalf("Error occurred while processing SQL query", err.Error())
    }

	// _, err = db.Exec("INSERT INTO user(name) values('Alice');")
	// if err != nil {
    //     panic(err)
    // }
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
