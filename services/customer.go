package services

import (
	"fmt"
	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	PRODUCT_NOT_FOUND          = "The product %s does not exist"
	PRODUCT_HAS_EXISTS_TO_USER = "The product %s is already linked to that customer"
)

type (
	CustomerService struct {
		session *mgo.Session
	}
)

func NewCustomerService(s *mgo.Session) *CustomerService {
	return &CustomerService{s}
}

func (uc CustomerService) CreateCustomer(user models.Customer) error {
	err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).Insert(&user)
	return err
}

func (uc CustomerService) GetCustomer(oid bson.ObjectId) (models.Customer, error) {

	u := models.Customer{}
	err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).FindId(oid).One(&u)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (uc CustomerService) GetAll() ([]models.Customer, error) {

	var results []models.Customer
	err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).Find(nil).All(&results)

	if err != nil {
		return results, err
	}

	return results, nil
}

func (uc CustomerService) RemoveCustomer(oid bson.ObjectId) bool {

	if err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).RemoveId(oid); err != nil {
		return false
	}

	return true
}

func (uc CustomerService) UpdateCustomer(oid bson.ObjectId, user models.Customer) (models.Customer, error) {

	if err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).UpdateId(oid, &user); err != nil {
		return user, err

	}

	return user, nil
}

func (uc CustomerService) AddFavorites(customerId bson.ObjectId, productIds []bson.ObjectId) (bool, error) {
	var products []models.Product

	productService := NewProductService(uc.session)

	customer, errCustomer := uc.GetCustomer(customerId)

	exitsProductsIds := uc.getProductIdsByCustomer(customer)

	for _, productId := range productIds {
		product, err := productService.GetProduct(productId)

		if err != nil {
			return false, errors.New(fmt.Sprintf(PRODUCT_NOT_FOUND, productId.Hex()))
		}

		if !uc.inObjectIdsSlice(exitsProductsIds, productId) {
			products = append(products, product)
		} else {
			return false, errors.New(fmt.Sprintf(PRODUCT_HAS_EXISTS_TO_USER, productId.Hex()))
		}
	}

	if errCustomer != nil {
		return false, errCustomer
	}

	customer.Products = append(customer.Products, products...)

	if err := uc.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).UpdateId(customerId, &customer); err != nil {
		return false, err
	}

	return true, nil
}

func (uc CustomerService) getProductIdsByCustomer(customer models.Customer) []bson.ObjectId {
	var exitsProductsIds []bson.ObjectId

	for _, product := range customer.Products {
		exitsProductsIds = append(exitsProductsIds, product.Id)
	}

	return exitsProductsIds
}

func (uc CustomerService) inObjectIdsSlice(slice []bson.ObjectId, oid bson.ObjectId) bool {
	for _, item := range slice {
		if item == oid {
			return true
		}
	}

	return false
}
