package main

import (
	"github.com/admpub/log"
	"github.com/webx-top/echo"
	echoswagger "github.com/webx-top/echo-swagger"
	_ "github.com/webx-top/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/webx-top/echo/engine/standard"
	mw "github.com/webx-top/echo/middleware"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	defer log.Close()
	e := echo.New()
	e.Use(mw.Recover(), mw.Log())

	e.Get("/swagger/*", echoswagger.WrapHandler)
	/*
		Or can use EchoWrapHandler func with configurations.
		url := echoswagger.URL("http://localhost:1323/swagger/doc.json") //The url pointing to API definition
		e.Get("/swagger/*", echoswagger.EchoWrapHandler(url))
	*/

	e.Get("/", func(c echo.Context) error {
		return c.Redirect("/swagger/index.html")
	})

	e.Run(standard.New(":1323"))
}
