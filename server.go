package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramdani10/mux-example/routes"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request){
		fmt.Fprintln(resp, "Up and Running")
	})
	router.HandleFunc("/post", routes.GetPosts).Methods("GET")
	router.HandleFunc("/post", routes.AddPost).Methods("POST")


	log.Println("server listening port", port)
	http.ListenAndServe(port,router)
}