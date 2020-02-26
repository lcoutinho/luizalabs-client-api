package services

import (
	"fmt"
	"github.com/lcoutinho/luizalabs-client-api/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	AuthService struct {
		session *mgo.Session
	}
)

func NewAuthService(s *mgo.Session) *AuthService {
	return &AuthService{s}
}

func (uc AuthService) HasUser(username string, password string) bool {

	if err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_USERS).Find(bson.M{"username": username, "password": password}).One(nil); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
