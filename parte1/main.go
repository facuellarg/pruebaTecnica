package main

import (
	"ipcom1/parte1/api"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echoServer := echo.New()
	echoServer.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}),
	)
	echoServer.GET("ventas/:date", api.PurchasesByDay)
	if err := echoServer.Start(":8000"); err != nil {
		log.Fatal(err.Error())
	}

}
