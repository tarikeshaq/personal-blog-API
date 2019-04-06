package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `bson:"title" json:"title"`
	Content
	Date string `bson:"date" json:"date"`
}

type Content struct {
	Overall    string `bson:"overall" json:"overall"`
	Like       string `bson:"like" json:"like"`
	NotLike    string `bson:"not_like" json:"not_like"`
	Rating     int32  `bson:"rating" json:"rating"`
	RatingDesc string `bson:"rating_desc" json:"rating_desc"`
}
