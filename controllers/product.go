package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"github.com/lcoutinho/luizalabs-client-api/services"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	ProductController struct {
		services *services.ProductService
	}
)

func NewProductController(s *mgo.Session) *ProductController {
	uc := services.NewProductService(s)
	return &ProductController{uc}
}

func (uc ProductController) ProductList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	results, err := uc.services.GetAll(page)

	if err != nil {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Product doesn't exist"))
		return
	}

	c.JSON(200, Response(results, true, len(results), false))

}

func (uc ProductController) GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}
	oid := bson.ObjectIdHex(id)

	u, err := uc.services.GetProduct(oid)
	if err != nil {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Product doesn't exist"))
		return
	}

	c.JSON(200, Response(u, true, 1, false))
}

func (uc ProductController) CreateProduct(c *gin.Context) {

	var product models.Product

	c.BindJSON(&product)

	err := uc.services.CreateProduct(product)

	if err == nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, Response(product, true, 1, false))
		return
	}

	c.AbortWithStatusJSON(500, Response(false, false, 0, err.Error()))

}

func (uc ProductController) RemoveProduct(c *gin.Context) {

	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}

	oid := bson.ObjectIdHex(id)

	isSuccess := uc.services.RemoveProduct(oid)

	if isSuccess {
		c.JSON(200, Response("Success", true, 1, false))
		return
	}

	c.AbortWithStatusJSON(500, Response(false, false, 0, "An error occured"))

}

func (uc ProductController) UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	c.BindJSON(&product)

	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(404, Response(false, false, 0, "Invalid id"))
		return
	}
	oid := bson.ObjectIdHex(id)

	u, err := uc.services.UpdateProduct(oid, product)

	if err == nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(200, Response(u, true, 1, false))
		return
	}
		c.AbortWithStatusJSON(500, Response(false, false, 0, err.Error()))
}
