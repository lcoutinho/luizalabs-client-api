package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lcoutinho/luizalabs-client-api/models"
	"github.com/lcoutinho/luizalabs-client-api/services"
	"github.com/lcoutinho/luizalabs-client-api/services/auth"
	"github.com/patrickmn/go-cache"
	"gopkg.in/mgo.v2"
)

type JwtController struct {
	cacheClient *cache.Cache
	services    *services.AuthService
}

func NewAuthController(cacheClient *cache.Cache, session *mgo.Session) *JwtController {
	authService := services.NewAuthService(session)
	return &JwtController{cacheClient, authService}
}

func (ac JwtController) RefreshToken(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	if !ac.services.HasUser(user.Username, user.Password) {
		c.AbortWithStatusJSON(401, Response(false, false, 0, "User not found"))
		return
	}

	resource := c.Request.Header.Get("Resource")

	createdToken := ac.GenarateNewToken(resource)

	tokenData := map[string]string{
		"token": createdToken,
	}

	c.JSON(200, Response(tokenData, true, 0, false))
}

func (ac JwtController) GenarateNewToken(resource string) string {

	authService, err := auth.GetAuthStrategyService(ac.cacheClient)

	if err != nil {
		panic(err)
	}

	return authService.GenerateToken(resource)
}
