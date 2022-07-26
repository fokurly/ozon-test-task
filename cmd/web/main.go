package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	interfacee "ozonTest/pkg/storage/interface"
	"ozonTest/pkg/storage/local"
	"ozonTest/pkg/storage/postgresql"
)

var Storage interfacee.Storage

const (
	port = ":4000"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var postgres bool
	flag.BoolVar(&postgres, "postgres", true, "Choose db. True-postgres, false - local")
	flag.Parse()

	if !postgres {
		Storage = local.NewLocalStorage()
	} else {
		user, _ := os.LookupEnv("POSTGRES_USER")
		pass, _ := os.LookupEnv("POSTGRES_PASSWORD")
		host, _ := os.LookupEnv("POSTGRES_HOST")
		dbPort, _ := os.LookupEnv("POSTGRES_PORT")
		dbName, _ := os.LookupEnv("POSTGRES_NAME")

		Storage = postgresql.NewPostgresStorage(user, pass, host, dbPort, dbName)
	}

	route := mux.NewRouter()

	route.HandleFunc("/create", createShortLink).Methods("POST")
	route.HandleFunc("/getLongLink", getLongLink).Methods("GET")

	log.Println("Запуск сервера на http://localhost:4000")
	err := http.ListenAndServe(port, route)

	fmt.Println(http.LocalAddrContextKey)
	log.Fatal(err)
}
