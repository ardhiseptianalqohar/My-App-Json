package config

// import "github.com/tkanos/gonfig"

// type Configuration struct {
// 	DB_USERNAME string
// 	DB_PASSWORD string
// 	DB_PORT     string
// 	DB_HOST     string
// 	DB_NAME     string
// }

// func GetConfig() Configuration {
// 	conf := Configuration{}
// 	gonfig.GetConf("config/config.jason", &conf)
// 	return conf
// }

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", dbConnection)

	if err != nil {
		log.Fatal("gagal membuka sql", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("Gagal menghubungkan DB", errPing.Error())
	} else {
		fmt.Println("=========================================")
		fmt.Println("\tSUCCES CONNECT TO DATABASE")
		fmt.Println("=========================================")
	}

	return db
}
