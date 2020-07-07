package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request){
		fmt.Fprintln(resp, "Up and Running")
	})
	router.HandleFunc("/post", getPost).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")


	log.Println("server listening port", port)
	http.ListenAndServe(port,router)
}