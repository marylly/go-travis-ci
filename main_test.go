package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(router *gin.Engine, method, path string) *httptest.ResponseRecorder {
	request, _ 		:= http.NewRequest(method, path, nil)
	responseWriter 	:= httptest.NewRecorder()
	router.ServeHTTP(responseWriter, request)
	return responseWriter
 }

func TestHealthCheck(t *testing.T) {
	application := Routes()
	response 	:= performRequest(application, "GET", "/health-check")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Service OK! =]", response.Body.String())
}
