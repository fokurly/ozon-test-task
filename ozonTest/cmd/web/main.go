package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	interfacee "ozonTest/pkg/storage/interface"
	"ozonTest/pkg/storage/local"
	"ozonTest/pkg/storage/postgresql"
)

var Storage interfacee.Storage

func main() {
	var postgres bool
	flag.BoolVar(&postgres, "postgres", true, "Choose db. True-postgres, false - local")
	flag.Parse()

	if !postgres {
		Storage = local.NewLocalStorage()
	} else {
		Storage = postgresql.NewPostgresStorage()
	}

	route := mux.NewRouter()

	route.HandleFunc("/create", createShortLink).Methods("POST")
	route.HandleFunc("/getLongLink", getLongLink).Methods("GET")

	log.Println("Запуск сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", route)
	fmt.Println(http.LocalAddrContextKey)
	log.Fatal(err)
}
