package echoswagger

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webx-top/echo"
	_ "github.com/webx-top/echo-swagger/example/docs"
	mw "github.com/webx-top/echo/middleware"
	otesting "github.com/webx-top/echo/testing"
)

func TestWrapHandler(t *testing.T) {

	router := echo.New()
	router.RouteDebug = true
	router.SetDebug(true)
	router.Use(mw.Recover(), mw.Log())

	router.Get("/*", WrapHandler)

	router.Commit()

	w1 := performRequest("GET", "/index.html", router)
	assert.Equal(t, 200, w1.Code)

	w2 := performRequest("GET", "/doc.json", router)
	assert.Equal(t, 200, w2.Code)

	w3 := performRequest("GET", "/favicon-16x16.png", router)
	assert.Equal(t, 200, w3.Code)

	w4 := performRequest("GET", "/notfound", router)
	assert.Equal(t, 404, w4.Code)

}

func performRequest(method, target string, e *echo.Echo) *httptest.ResponseRecorder {
	return otesting.Request(method, target, e)
}
