package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gin-gonic/gin"
)

func testHealthy(t *testing.T) {
	engine := setupTestcase()
	recorder := requestHTTP(engine, "GET", "/healthy")

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func setupTestcase() *gin.Engine {
	engine := gin.New()

	controller := NewController()
	controller.RegisterRoutes(engine)
	
	return engine
}

func requestHTTP(router http.Handler, method, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	return recorder
}