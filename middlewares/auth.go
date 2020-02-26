package middlewares

import (
	"github.com/lcoutinho/luizalabs-client-api/controllers"
	"github.com/lcoutinho/luizalabs-client-api/services/auth"
	"github.com/patrickmn/go-cache"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtFilter(cacheClient *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenJwt := c.Request.Header.Get("Authorization")

		if len(tokenJwt) == 0 {
			c.AbortWithStatusJSON(401, controllers.Response(false, false, 0, "Missing Authorization header"))
			return
		}
		validateJwt(tokenJwt, c, cacheClient)
	}
}

func validateJwt(myToken string, c *gin.Context, cacheClient *cache.Cache) {
	authService, _ := auth.GetAuthStrategyService(cacheClient)

	token, err := authService.ValidateToken(myToken)

	if err == nil && token.Valid {

		currentResource := c.MustGet("resource").(string)

		allowedResource := token.Claims.(jwt.MapClaims)["resource"]

		if !strings.Contains(allowedResource.(string), currentResource) {
			c.AbortWithStatusJSON(401, controllers.Response(false, false, 0, "Not allowed this resource"))
			return
		}

		c.Next()
	} else {
		c.AbortWithStatusJSON(403, controllers.Response(false, false, 0, err.Error()))
	}
}
