package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type IDB interface {
	Open() *sql.DB
}

type postgres struct {
	driver           string
	connectionString string
}

func (p *postgres) Open() *sql.DB {
	db, err := sql.Open(p.driver, p.connectionString)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to connect to the database. Error %v", err))
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("Pinging database failed")
	}
	log.Println("Successfully obtained the database instance")
	return db
}

func NewDB() IDB {
	host_port := os.Getenv("host_port")
	hostname := os.Getenv("hostname")
	username := os.Getenv("username")
	password := os.Getenv("password")
	databasename := os.Getenv("database_name")
	pgdb := new(postgres)
	pgdb.driver = "postgres"
	pgdb.connectionString = fmt.Sprintf("port=%s host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, databasename)
	log.Println("Connection string: " + pgdb.connectionString)
	return pgdb
}
