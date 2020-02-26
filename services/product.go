package services

import (
	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"strconv"
)

const (
	LIMIT_PAGE = 10
)

type (
	ProductService struct {
		session *mgo.Session
	}
)

func NewProductService(s *mgo.Session) *ProductService {
	return &ProductService{s}
}

func (uc ProductService) CreateProduct(user models.Product) error {
	err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).Insert(&user)
	return err
}

func (uc ProductService) GetProduct(oid bson.ObjectId) (models.Product, error) {

	u := models.Product{}
	err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).FindId(oid).One(&u)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (uc ProductService) GetAll(page int) ([]models.Product, error) {

	var results []models.Product
	err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).Find(nil).Skip((page - 1) * LIMIT_PAGE).Limit(LIMIT_PAGE).All(&results)

	if err != nil {
		return results, err
	}

	return results, nil
}

func (uc ProductService) RemoveProduct(oid bson.ObjectId) bool {

	if err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).RemoveId(oid); err != nil {
		return false
	}

	return true
}

func (uc ProductService) UpdateProduct(oid bson.ObjectId, user models.Product) (models.Product, error) {

	if err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).UpdateId(oid, &user); err != nil {
		return user, err

	}

	return user, nil
}
