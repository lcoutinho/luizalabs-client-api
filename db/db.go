package db

import (
	"gopkg.in/mgo.v2"
	"os"
)

func GetSession() *mgo.Session {

	s, err := mgo.Dial(os.Getenv("MONGODB_DSN"))

	if err != nil {
		panic(err)
	}

	return s
}
