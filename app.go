package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/tarikeshaq/personal-blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key string

const (
	hostKey     = key("hostKey")
	usernameKey = key("usernameKey")
	passwordKey = key("passwordKey")
	databaseKey = key("databaseKey")
)

var client *mongo.Client
var database *mongo.Database

type Post = models.Post

func GetAllBlogsHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var posts []Post
	collection := database.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// ctx, err:= context.WithTimeout(context.Background(), 10*time.Second)
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
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	blogId, _ := primitive.ObjectIDFromHex(params["blogId"])
	var post Post
	collection := database.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	response.Header().Set("content-type", "application/json")

	var post Post
	_ = json.NewDecoder(request.Body).Decode(&post)
	collection := database.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func RemoveBlogHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	vars := mux.Vars(request)
	blogId, _ := primitive.ObjectIDFromHex(vars["blogId"])
	collection := database.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"_id", blogId}}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func main() {
	// setup mongodb connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, hostKey, os.Getenv("MONGO_HOST"))
	ctx = context.WithValue(ctx, usernameKey, os.Getenv("MONGO_USERNAME"))
	ctx = context.WithValue(ctx, passwordKey, os.Getenv("MONGO_PASSWORD"))
	ctx = context.WithValue(ctx, databaseKey, os.Getenv("MONGO_DATABASE"))

	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(usernameKey).(string),
		ctx.Value(passwordKey).(string),
		ctx.Value(hostKey).(string),
		ctx.Value(databaseKey).(string),
	)

	fmt.Println(uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Errorf(err.Error())
	}
	err = client.Connect(ctx)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	database = client.Database("blogs")

	router := mux.NewRouter()
	router.HandleFunc("/blogs", GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/blogs/{blogId}", GetOneBlogHandler).Methods("GET")
	router.HandleFunc("/blogs", AddNewBlogHandler).Methods("POST")
	router.HandleFunc("/blogs/{blogId}", RemoveBlogHandler).Methods("DELETE")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
