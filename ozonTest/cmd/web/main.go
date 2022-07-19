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

const (
	port = ":4000"
)

func main() {

	//if err := initConfig(); err != nil {
	//	log.Fatalf("error initializing config: %s", err.Error())
	//}

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
	err := http.ListenAndServe(port, route)
	//	err := http.ListenAndServe(":"+viper.GetString("port"), route)
	fmt.Println(http.LocalAddrContextKey)
	log.Fatal(err)
}

//func initConfig() error {
//	viper.AddConfigPath("configs")
//	viper.SetConfigName("config")
//	return viper.ReadInConfig()
//}
