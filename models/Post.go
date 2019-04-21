package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	Summary   string             `json:"summary,omitempty" bson:"summary,omitempty"`
	Post_type string             `json:"post_type,omitempty" bson:"post_type,omitempty"`
	Content   Content            `json:"content,omitempty" bson:"content,omitempty"`
	Date      string             `json:"date,omitempty" bson:"date,omitempty"`
}

type Content struct {
	Overall     string `json:"overall,omitempty" bson:"overall,omitempty"`
	Like        string `json:"like,omitempty" bson:"like,omitempty"`
	Not_like    string `json:"not_like,omitempty" bson:"not_like,omitempty"`
	Rating      int    `json:"rating,omitempty" bson:"rating,omitempty"`
	Rating_desc string `json:"rating_desc,omitempty" bson:"rating_desc,omitempty"`
}
