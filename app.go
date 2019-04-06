package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	dao "personal-blog-api/dao"

	. "personal-blog-api/models"

	"github.com/gorilla/mux"
)

func AllBlogsEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "SHOULD RETURN ALL BLOGS")

}

func GetPostEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "SHOULD RETURN A SPECIFIC POST")
}

func RemovePostEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "NOT IMPLEMENTED, REMOVE")
}

func InsertPostEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var post Post

	if err := json.NewDecoder(request.Body).Decode(&post); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	post.ID = bson.NewObjectId()
	if err := dao.Insert(post); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(writer, http.StatusCreated, post)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", AllBlogsEndPoint).Methods("GET")
	r.HandleFunc("/posts/{id}", GetPostEndPoint).Methods("GET")
	r.HandleFunc("/posts", InsertPostEndPoint).Methods("POST")
	r.HandleFunc("/post", RemovePostEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
