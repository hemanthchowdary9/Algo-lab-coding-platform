package database

import (
	"coding-platform/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var psqlDb *sql.DB = nil

func GetPsqlConnection() *sql.DB {
	if psqlDb == nil {
		InitializeDB()
	}
	return psqlDb
}

func InitializeDB() *sql.DB {
	psql := openPsqlConnection()
	setupDb(psql)
	checkDbConnection(psql)
	psqlDb = psql
	fmt.Println("Successfully initialized DB")
	return psql
}

func setupDb(sqlDb *sql.DB) {
	sqlDb.SetMaxIdleConns(100)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
}

func checkDbConnection(sqlDb *sql.DB) {
	if err := sqlDb.Ping(); err != nil {
		log.Fatal("Error pinging database", err)
	}
	fmt.Println("DB ping successful...!")
}

func openPsqlConnection() *sql.DB {
	database := config.FetchDatabaseConfigs()
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database["host"],
		database["port"],
		database["user"],
		database["password"],
		database["name"])
	fmt.Println("DB connection string: ", psqlConn)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal("error in opening db connection", err)
	}
	return db
}
