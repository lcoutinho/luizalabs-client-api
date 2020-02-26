package services

import (
	"fmt"
	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/lcoutinho/luizalabs-client-api/db"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strconv"
	"testing"
)

type (
	CustomerServiceTestSuite struct {
		suite.Suite
		assert          *assert.Assertions
		customerService *CustomerService
		session         *mgo.Session
		oid             bson.ObjectId
		oidFavorites    bson.ObjectId
		productIds      []bson.ObjectId
	}
)

func (s *CustomerServiceTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.session = db.GetSession()
	s.customerService = NewCustomerService(s.session)
}

func TestCustomerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerServiceTestSuite))
}

func (s *CustomerServiceTestSuite) TearDownSuite() {
	s.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).RemoveAll(bson.M{"name": "Teste"})
}

func (s *CustomerServiceTestSuite) TestCreateCustomer() {

	customer := models.Customer{
		Name:  "Teste",
		Email: "teste@teste.com",
	}

	err := s.customerService.CreateCustomer(customer)
	s.assert.NoError(err)
}

func (s *CustomerServiceTestSuite) TestCreateExistsCustomer() {

	customer := models.Customer{
		Name:  "Teste",
		Email: "teste1@teste.com",
	}

	err := s.customerService.CreateCustomer(customer)
	s.assert.NoError(err)
	err = s.customerService.CreateCustomer(customer)
	s.assert.Error(err)
	s.assert.Equal(`E11000 duplicate key error collection: customer_db.customers index: email_1 dup key: { email: "teste1@teste.com" }`, err.Error())
}

func (s *CustomerServiceTestSuite) TestGetCustomer() {

	var customerResult models.Customer

	customer := models.Customer{
		Name:  "Teste",
		Email: "teste2@teste.com",
	}

	err := s.customerService.CreateCustomer(customer)
	s.assert.NoError(err)
	s.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).Find(bson.M{"email": customer.Email}).One(&customerResult)

	s.oid = customerResult.Id

	customerGet, err := s.customerService.GetCustomer(s.oid)
	s.assert.NoError(err)

	s.assert.Equal(customer.Email, customerGet.Email)
	s.assert.Equal(customer.Name, customerGet.Name)

}

func (s *CustomerServiceTestSuite) TestGetAll() {

	_, err := s.customerService.GetAll()
	s.assert.NoError(err)

}

func (s *CustomerServiceTestSuite) TestUpdateCustomer() {
	customerEdit := models.Customer{
		Name:  "Teste",
		Email: "teste3@teste.com",
	}

	customer, err := s.customerService.UpdateCustomer(s.oid, customerEdit)
	s.assert.NoError(err)
	s.assert.Equal(customerEdit.Email, customer.Email)
}

func (s *CustomerServiceTestSuite) TestUpdateRemoveCustomer() {

	isSuccess := s.customerService.RemoveCustomer(s.oid)
	s.assert.True(isSuccess)
}

func (s *CustomerServiceTestSuite) TestUpdateRemoveCustomerError() {
	customerEdit := models.Customer{
		Name:  "Teste",
		Email: "teste3@teste.com",
	}

	_, err := s.customerService.UpdateCustomer(s.oid, customerEdit)
	s.assert.Error(err)

}
func (s *CustomerServiceTestSuite) TestUpdateRemoveCustomerErrorNotFound() {

	isSuccess := s.customerService.RemoveCustomer(s.oid)
	s.assert.False(isSuccess)
}

func (s *CustomerServiceTestSuite) TestAddFavorites() {

	productIds := s.createProducts(3)
	oid := s.createUser()
	s.oidFavorites = oid
	s.productIds = productIds
	isSuccess, err := s.customerService.AddFavorites(oid, productIds)
	s.assert.True(isSuccess)
	s.assert.NoError(err)
}
func (s *CustomerServiceTestSuite) TestAddFavoritesExists() {

	isSuccess, err := s.customerService.AddFavorites(s.oidFavorites, s.productIds)
	s.assert.False(isSuccess)
	s.assert.Error(err)
	s.assert.Contains(err.Error(), "linked to that customer")
}

func (s *CustomerServiceTestSuite) TestAddFavoritesNotFound() {
	invalidProductIds := []bson.ObjectId{s.oid, s.oidFavorites}

	isSuccess, err := s.customerService.AddFavorites(s.oidFavorites, invalidProductIds)
	s.assert.False(isSuccess)
	s.assert.Error(err)
	s.assert.Contains(err.Error(), "does not exist")

}

func (s *CustomerServiceTestSuite) createUser() bson.ObjectId {

	var customerResult models.Customer

	customer := models.Customer{
		Name:  "Teste",
		Email: "teste82@teste.com",
	}

	s.customerService.CreateCustomer(customer)
	s.session.DB(config.DB_NAME).C(config.DB_COLLECTION_CUSTOMERS).Find(bson.M{"email": customer.Email}).One(&customerResult)

	return customerResult.Id
}
func (s *CustomerServiceTestSuite) createProducts(quantity int) []bson.ObjectId {
	productService := NewProductService(s.session)
	var productIds []bson.ObjectId
	for i := 0; i <= quantity; i++ {
		var productResult models.Product

		t := strconv.Itoa(rand.Intn(1000))
		titleProduct := fmt.Sprintf("Produto teste %s", t)

		product := models.Product{
			Title:       titleProduct,
			Price:       11.1,
			Image:       "http://teste.com/teste.png",
			Brand:       "Teste",
			ReviewScore: 10,
		}
		err := productService.CreateProduct(product)

		if err == nil {
			s.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).Find(bson.M{"title": titleProduct}).One(&productResult)
			productIds = append(productIds, productResult.Id)
		}

	}

	return productIds
}
