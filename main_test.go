package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestSetupRouter(t *testing.T) {

	router := SetupRouter()

	w := performRequest(router, "GET", "/customer")
	assert.Equal(t, http.StatusOK, w.Code)

}
