package middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"strings"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Can not get request. Error %s", err.Error())
			c.AbortWithStatusJSON(401, err)
			return
		}

		jsonFile := fmt.Sprintf("/%s.json", GetFile(c))

		schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", jsonFile))
		requestJSONLoader := gojsonschema.NewBytesLoader(requestBody)

		result, err := gojsonschema.Validate(schemaLoader, requestJSONLoader)

		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(401, err)
		}

		if result.Valid() != true {

			var formatedResponse []string

			for _, v := range result.Errors() {
				formatedResponse = append(formatedResponse, v.String())

			}

			c.AbortWithStatusJSON(401, formatedResponse)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	}
}

func GetFile(c *gin.Context) string {

	path := c.Request.URL.Path

	segmentsUrl := strings.Split(path, "/")

	resource := segmentsUrl[1]

	c.Set("resource", resource)

	return resource
}
