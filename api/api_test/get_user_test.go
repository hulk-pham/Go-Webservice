package api_test

import (
	"encoding/json"
	"hulk/go-webservice/api"
	"hulk/go-webservice/common"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestListUserHandler(t *testing.T) {
	r := api.InitRouter()
	common.InitDB()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/", nil)
	r.ServeHTTP(w, req)

	var result common.JSONResult
	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, result.Data)
}
