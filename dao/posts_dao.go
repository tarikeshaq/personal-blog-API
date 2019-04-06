package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "posts"
)

func (m *PostsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *PostsDAO) FindAll() ([]Post, error) {
	var posts []Post
	err := db.C(COLLECTION).Find(bson.M{}).All(&posts)
	return posts, err
}
