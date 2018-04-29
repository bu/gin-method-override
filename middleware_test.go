package methodOverride

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	// no debug mode
	gin.SetMode(gin.ReleaseMode)

	// create a default
	r := gin.Default()

	// our middle-ware
	r.Use(ProcessMethodOverride(r))

	// routes
	r.POST("/test", testPOST)
	r.GET("/test", testGET)
	r.PUT("/test", testPUT)

	// return to caller
	return r
}

func TestNoDispatchPOST(t *testing.T) {
	// setup
	router := setupRouter()

	// prepare
	RequestURL := "/test"
	RequestMethod := "POST"
	RequestBody := bytes.NewBuffer([]byte("testing=1")) // no method

	ExpectedResponseStatus := 200 // from testPOST
	ExpectedResponse := "0"       // from testPOST

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(RequestMethod, RequestURL, RequestBody)
	router.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
	assert.Equal(t, ExpectedResponse, w.Body.String())
}

func TestDispatchToPUT(t *testing.T) {
	// setup
	router := setupRouter()

	// prepare
	RequestURL := "/test"
	RequestMethod := "POST"
	RequestBody := bytes.NewBuffer([]byte("_method=PUT&testing=1")) // method = PUT

	ExpectedResponseStatus := 200 // from testPOST
	ExpectedResponse := "1"       // from testPOST

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(RequestMethod, RequestURL, RequestBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
	assert.Equal(t, ExpectedResponse, w.Body.String())
}

func testGET(c *gin.Context) {
	c.String(200, "pong")
}

func testPOST(c *gin.Context) {
	c.String(200, "0")
}

func testPUT(c *gin.Context) {
	c.String(200, "1")
}
