package main

import (
	"database/sql"
	"ecom/cmd/api"
	"ecom/config"
	"ecom/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := config.Load()
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBHost + ":" + cfg.DBPort,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}


func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Successfully Connected ! ")
}