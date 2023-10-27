package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func Initialize() {

	host, ok := os.LookupEnv("APOD_DB_HOST")
	if !ok {
		fmt.Println("APOD_DB_HOST environment variable not set")
		os.Exit(1)
	}
	port, ok := os.LookupEnv("APOD_DB_PORT")
	if !ok {
		fmt.Println("APOD_DB_PORT environment variable not set")
		os.Exit(1)
	}
	user, ok := os.LookupEnv("APOD_DB_USER")
	if !ok {
		fmt.Println("APOD_DB_USER environment variable not set")
		os.Exit(1)
	}
	password, ok := os.LookupEnv("APOD_DB_PASSWORD")
	if !ok {
		fmt.Println("APOD_DB_PASSWORD environment variable not set")
		os.Exit(1)
	}
	dbname, ok := os.LookupEnv("APOD_DB_NAME")
	if !ok {
		fmt.Println("APOD_DB_NAME environment variable not set")
		os.Exit(1)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database connections pool initialized")
}

func Close() {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database connection closed")
}

func GetConnection() *sql.DB {
	return db
}
