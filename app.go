package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllBlogsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Not yet Implemented, gets all posts"))
}

func GetOneBlogHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	blogId := vars["blogId"]
	writer.Write([]byte("Not yet implemented, gets one post with Id" + blogId))
}

func AddNewBlogHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Not yet implemented, creates a new post using the payload from the request"))
}

func RemoveBlogHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	blogId := vars["blogId"]
	writer.Write([]byte("Not yet implemented, removes a post with ID " + blogId))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/blogs", GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/blogs/{blogId}", GetOneBlogHandler).Methods("GET")
	router.HandleFunc("/blogs", AddNewBlogHandler).Methods("POST")
	router.HandleFunc("/blogs/{blogId}", RemoveBlogHandler).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}
