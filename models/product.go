package models

import "gopkg.in/mgo.v2/bson"

type Product struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Price       float64       `json:"price" bson:"price"`
	Image       string        `json:"image" bson:"image"`
	Brand       string        `json:"brand" bson:"brand"`
	Title       string        `json:"title" bson:"title"`
	ReviewScore int           `json:"reviewScore" bson:"reviewScore"`
}
