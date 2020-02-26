package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

func GetSession() *mgo.Session {

	s, err := mgo.Dial(os.Getenv("MONGODB_DSN"))

	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	return s
}
