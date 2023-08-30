package db

import (
	"fmt"
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var once sync.Once
var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db == nil {
		once.Do(ConnectDB)
	}
	return db
}

func ConnectDB() {
	cfg := config.GetConfig()
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", cfg.DataBase.Host, cfg.DataBase.Port, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.DBName)

	var err error
	db, err = sqlx.Connect("postgres", dbInfo)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected!")
}
