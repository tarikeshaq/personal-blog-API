package dao

import (
	"log"

	. "personal-blog-api/models"

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

func (p *PostsDAO) Connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

func (p *PostsDAO) FindAll() ([]Post, error) {
	var posts []Post
	err := db.C(COLLECTION).Find(bson.M{}).All(&posts)
	return posts, err
}

func (p *PostsDAO) FindById(id string) (Post, error) {
	var post Post
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&post)
	return post, err
}

func (p *PostsDAO) Insert(post Post) error {
	err := db.C(COLLECTION).Insert(post)
	return err
}

func (p *PostsDAO) Delete(post Post) error {
	err := db.C(COLLECTION).Remove(&post)
	return err
}
