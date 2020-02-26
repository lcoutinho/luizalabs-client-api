package models

import "gopkg.in/mgo.v2/bson"

type Customer struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Products []Product     `json:"products,omitempty" bson:"products,omitempty"`
}
