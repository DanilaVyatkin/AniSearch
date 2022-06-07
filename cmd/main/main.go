package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("hello world"))
}

func main() {
	log.Println("create router")
	router := httprouter.New()

	router.GET("/hello", Hello)
	log.Fatalln(http.ListenAndServe(":8080", router))
}
