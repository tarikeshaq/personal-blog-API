package main

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBlogsHandler(response http.ResponseWriter, request *http.Request) {
	writeHeaders(response)
	var posts []Post
	collection := database.Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}

func GetOneBlogHandler(response http.ResponseWriter, request *http.Request) {
	writeHeaders(response)

	params := mux.Vars(request)
	blogId, _ := primitive.ObjectIDFromHex(params["blogId"])
	var post Post
	collection := database.Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", blogId}}
	err := collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + params["blogId"] + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

func AddNewBlogHandler(response http.ResponseWriter, request *http.Request) {
	writeHeaders(response)

	var post Post
	_ = json.NewDecoder(request.Body).Decode(&post)
	collection := database.Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func RemoveBlogHandler(response http.ResponseWriter, request *http.Request) {
	writeHeaders(response)

	vars := mux.Vars(request)
	blogID, _ := primitive.ObjectIDFromHex(vars["blogId"])
	collection := database.Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", blogID}}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func writeHeaders(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET")
	response.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func BasicAuth(handler http.HandlerFunc, userHash, passHash []byte, realm string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, pass, ok := request.BasicAuth()

		if !ok || subtle.ConstantTimeCompare(hasher(user), userHash) != 1 ||
			subtle.ConstantTimeCompare(hasher(pass), passHash) != 1 {
			writer.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			writer.WriteHeader(401)
			writer.Write([]byte("Unauthorised.\n"))
			return
		}

		handler(writer, request)
	}
}
