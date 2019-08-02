package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello Athens!")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Println("starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
