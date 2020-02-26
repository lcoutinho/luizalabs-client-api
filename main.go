package main

import (
	// Standard library packages

	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/lcoutinho/luizalabs-client-api/db"

	"github.com/lcoutinho/luizalabs-client-api/middlewares"
	"github.com/patrickmn/go-cache"
	"time"

	// Third party packages
	"github.com/gin-gonic/gin"
	"github.com/lcoutinho/luizalabs-client-api/controllers"
)

func main() {

	// Get a user resource
	router := SetupRouter()

	router.Run(":8000")
}

func SetupRouter() *gin.Engine {

	sessionDB := db.GetSession()
	cacheClient := cache.New(config.JWT_TIME_EXPIRE, 1*time.Minute)

	cc := controllers.NewCustomerController(sessionDB)
	pc := controllers.NewProductController(sessionDB)
	ac := controllers.NewAuthController(cacheClient, sessionDB)

	router := gin.Default()

	router.Use(middlewares.Commom())
	router.POST("/refresh", ac.RefreshToken)
	router.GET("/customer", cc.CustomerList)
	router.GET("/customer/:id", cc.GetCustomer)
	router.GET("/product", pc.ProductList)
	router.GET("/product/:id", pc.GetProduct)
	router.POST("/customer/favorites/:id", cc.AddFavorites)
	router.Use(middlewares.JwtFilter(cacheClient))
	router.DELETE("/product/:id", pc.RemoveProduct)
	router.DELETE("/customer/:id", cc.RemoveCustomer)
	router.Use(middlewares.Validator())
	router.POST("/customer", cc.CreateCustomer)
	router.PUT("/customer/:id", cc.UpdateCustomer)
	router.POST("/product", pc.CreateProduct)
	router.PUT("/product/:id", pc.UpdateProduct)

	return router
}
