package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerRequest() {
	rutas := mux.NewRouter().StrictSlash(true)

	rutas.Handle("/getdata", http.HandlerFunc(getRandomizer))

	log.Fatal(http.ListenAndServe(":8080", rutas))

}
