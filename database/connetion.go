package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	dbtype := os.Getenv("DBTYPE")

	log.Println("Connecting to database " + dbname + " in " + dbtype + "...")

	stringConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Mexico_City", host, user, password, dbname, port)
	database, err := sql.Open(dbtype, stringConnection)

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	log.Println("Connected Successfully!")

	return database
}
