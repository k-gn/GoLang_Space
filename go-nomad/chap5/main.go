package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerHome) // get request
	e.POST("/scrap", handlerScrap)
	e.Logger.Fatal(e.Start(":1323")) // port
}

func handlerHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! [GET]")
}

func handlerScrap(c echo.Context) error {
	defer os.Remove("cards.csv")
	fmt.Println(c.FormValue("term"))
	return c.Attachment("cards.csv", "cards.csv") // return file
}
