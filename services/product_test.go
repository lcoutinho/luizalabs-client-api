package services

import (
	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/lcoutinho/luizalabs-client-api/db"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type (
	ProductServiceTestSuite struct {
		suite.Suite
		assert         *assert.Assertions
		productService *ProductService
		session        *mgo.Session
		oid            bson.ObjectId
		oidFavorites   bson.ObjectId
		productIds     []bson.ObjectId
	}
)

func (s *ProductServiceTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.session = db.GetSession()
	s.productService = NewProductService(s.session)
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}

func (s *ProductServiceTestSuite) TearDownSuite() {
	s.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).RemoveAll(bson.M{"title": "Product Teste"})
}

func (s *ProductServiceTestSuite) TestCreateCustomer() {

	product := models.Product{
		Title:       "Product Teste",
		Price:       11.1,
		Image:       "http://teste.com/teste.png",
		Brand:       "Teste",
		ReviewScore: 10,
	}

	err := s.productService.CreateProduct(product)
	s.assert.NoError(err)
}

func (s *ProductServiceTestSuite) TestGetProduct() {

	var productResult models.Product

	product := models.Product{
		Title:       "Product Teste",
		Price:       11.1,
		Image:       "http://teste.com/teste.png",
		Brand:       "Teste",
		ReviewScore: 10,
	}

	err := s.productService.CreateProduct(product)
	s.assert.NoError(err)
	s.session.DB(config.DB_NAME).C(config.DB_COLLECTION_PRODUCTS).Find(bson.M{"title": product.Title}).One(&productResult)

	s.oid = productResult.Id
	productGet, err := s.productService.GetProduct(s.oid)
	s.assert.NoError(err)

	s.assert.Equal(product.Title, productGet.Title)
	s.assert.Equal(product.Price, productGet.Price)
	s.assert.Equal(product.Image, productGet.Image)
	s.assert.Equal(product.Brand, productGet.Brand)
	s.assert.Equal(product.ReviewScore, productGet.ReviewScore)

}

func (s *ProductServiceTestSuite) TestGetAll() {

	_, err := s.productService.GetAll(1)
	s.assert.NoError(err)

}

func (s *ProductServiceTestSuite) TestRemoveUpdateProduct() {
	productEdit := models.Product{
		Title:       "Product Teste",
		Price:       11.1,
		Image:       "http://teste.com/teste.png",
		Brand:       "Nova Brand",
		ReviewScore: 10,
	}

	product, err := s.productService.UpdateProduct(s.oid, productEdit)
	s.assert.NoError(err)
	s.assert.Equal(productEdit.Brand, product.Brand)
}

func (s *ProductServiceTestSuite) TestUpdateRemoveProduct() {

	isSuccess := s.productService.RemoveProduct(s.oid)
	s.assert.True(isSuccess)
}
