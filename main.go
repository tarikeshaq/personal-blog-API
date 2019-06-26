package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
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

func hasher(s string) []byte {
	val := sha256.Sum256([]byte(s))
	return val[:]
}

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
		log.Fatalf(err.Error())
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return client.Database("blogs")
}

func setupRoutes() *mux.Router {
	password := hasher(os.Getenv("PASSWORD"))
	username := hasher(os.Getenv("USERNAME"))
	router := mux.NewRouter()
	router.HandleFunc("/blogs", GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/blogs/{blogId}", GetOneBlogHandler).Methods("GET")
	router.HandleFunc("/blogs", BasicAuth(AddNewBlogHandler,
		username, password,
		"Please input your username and password")).Methods("POST")

	router.HandleFunc("/blogs/{blogId}", BasicAuth(RemoveBlogHandler,
		username, password,
		"Please input your username and password")).Methods("DELETE")
	return router
}

func main() {
	ctx, cancel := setupContext()
	database = setupDB(ctx, cancel)
	router := setupRoutes()

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
