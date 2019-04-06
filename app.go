package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllBlogsEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "SHOULD RETURN ALL BLOGS")

}

func GetPostEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "SHOULD RETURN A SPECIFIC POST")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", AllBlogsEndPoint).Methods("GET")
	r.HandleFunc("/posts/{id}", GetPostEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
