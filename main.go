package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/tarikeshaq/personal-blog-api/models"
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

func setupContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = context.WithValue(ctx, hostKey, os.Getenv("MONGO_HOST"))
	ctx = context.WithValue(ctx, usernameKey, os.Getenv("MONGO_USERNAME"))
	ctx = context.WithValue(ctx, passwordKey, os.Getenv("MONGO_PASSWORD"))
	ctx = context.WithValue(ctx, databaseKey, os.Getenv("MONGO_DATABASE"))
	return ctx, cancel
}

func setupDB(ctx context.Context, cancel context.CancelFunc) *mongo.Database {
	defer cancel()
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(usernameKey).(string),
		ctx.Value(passwordKey).(string),
		ctx.Value(hostKey).(string),
		ctx.Value(databaseKey).(string),
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Errorf(err.Error())
	}
	err = client.Connect(ctx)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return client.Database("blogs")
}

func setupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/blogs", GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/blogs/{blogId}", GetOneBlogHandler).Methods("GET")
	router.HandleFunc("/blogs", BasicAuth(AddNewBlogHandler,
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"),
		"Please input your username and password")).Methods("POST")

	router.HandleFunc("/blogs/{blogId}", BasicAuth(RemoveBlogHandler,
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"),
		"Please input your username and password")).Methods("DELETE")
	return router
}

func main() {
	ctx, cancel := setupContext()
	database = setupDB(ctx, cancel)
	router := setupRoutes()

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
