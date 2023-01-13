package database

import (
	"context"
	"log"
	"time"

	"github.com/anandMohanan/taa_backend_golang/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = ""

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}

}

func (db *DB) GetPost(id string) *model.Post {
	var post model.Post
	return &post
}

func (db *DB) GetPosts() []*model.Post {
	var posts []*model.Post
	return posts
}

func (db *DB) LikePost(id string) *model.Post {
	var post model.Post
	return &post
}

func (db *DB) DeletePost(id string) string {
	return ""
}

func (db *DB) CreatePost(title string, body string) *model.Post {
	var post model.Post
	return &post
}

func (db *DB) Login(username string, password string) *model.User {
	var user model.User
	return &user
}

func (db *DB) Register(input *model.RegisterInput) *model.User {
	var user model.User
	return &user
}
