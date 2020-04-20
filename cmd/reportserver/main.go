package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"dendrix.io/nayalabs/reportserver/api/controllers"
	"dendrix.io/nayalabs/reportserver/database"
	"dendrix.io/nayalabs/reportserver/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found. Error:%s", err.Error())
	}
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		port = 3001
	}
	db := connectToDatabase()
	defer db.Close()
	handler := controllers.StartUp(repository.NewPaymentRepository())
	go startHTTPServer(port, handler)

	log.Println(fmt.Sprintf("Started http server on port %d", port))
	//Block and wait for interrupt signal
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	//Block until we receive our signal.
	<-c
	log.Println("Shutting down the server. Goodbye!")
}

//Start function starts up the server
func startHTTPServer(port int, handler http.Handler) {
	addr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase() *sql.DB {
	db := database.NewDB().Open()
	repository.SetDatabase(db)
	return db
}
