package main

import (
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/route"
	"github.com/gorilla/mux"
)

func init() {
    functions.HTTP("WebHook", route.URL)
}

func main() {
    // Connect to the database and set up the index
    config.ConnectDB()

    r := mux.NewRouter()
    r.HandleFunc("/register", controller.Register).Methods("POST")
    r.HandleFunc("/login", controller.Login).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", r))
}