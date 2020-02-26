package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"github.com/lcoutinho/luizalabs-client-api/services"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	CustomerController struct {
		services *services.CustomerService
	}
)

func NewCustomerController(s *mgo.Session) *CustomerController {
	uc := services.NewCustomerService(s)
	return &CustomerController{uc}
}

func (uc CustomerController) CustomerList(c *gin.Context) {

	results, err := uc.services.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(500, Response(false, false, 0, err.Error()))
		return
	}

	c.JSON(200, Response(results, true, len(results), false))

}

func (uc CustomerController) GetCustomer(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}
	oid := bson.ObjectIdHex(id)

	u, err := uc.services.GetCustomer(oid)
	if err != nil {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Customer doesn't exist"))
		return
	}

	c.JSON(200, Response(u, true, 1, false))
}

func (uc CustomerController) CreateCustomer(c *gin.Context) {

	var customer models.Customer

	c.BindJSON(&customer)

	err := uc.services.CreateCustomer(customer)

	if err == nil {
		c.JSON(201, Response(customer, true, 1, false))
		return
	}

	c.AbortWithStatusJSON(500, Response(false, false, 0, err.Error()))

}

func (uc CustomerController) RemoveCustomer(c *gin.Context) {

	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}
	oid := bson.ObjectIdHex(id)

	isSuccess := uc.services.RemoveCustomer(oid)

	if isSuccess {
		c.JSON(200, Response("Success", true, 1, false))
		return
	}

	c.AbortWithStatusJSON(404, Response(false, false, 0, "Internal error"))

}

func (uc CustomerController) UpdateCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var json models.Customer

	c.BindJSON(&json)

	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}
	oid := bson.ObjectIdHex(id)

	u, err := uc.services.UpdateCustomer(oid, json)

	if err == nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(200, Response(u, true, 1, false))
	} else {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "An error occured"))
	}

}

func (uc CustomerController) AddFavorites(c *gin.Context) {
	customerId := c.Params.ByName("id")
	var productIds []bson.ObjectId

	c.BindJSON(&productIds)

	if !bson.IsObjectIdHex(customerId) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}
	oid := bson.ObjectIdHex(customerId)

	isSuccess, err := uc.services.AddFavorites(oid, productIds)

	if !isSuccess {
		c.AbortWithStatusJSON(500, Response(false, false, 0, err.Error()))
		return
	}

	c.JSON(201, Response(fmt.Sprintf("%d favorite products added for this customer", len(productIds)), true, 1, false))
}
