package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shoaibashk/dokcli/ui"
)

type pretty interface {
	[]byte | string
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func PrettyJSON[T pretty](data T) (string, error) {
	var prettyJSON bytes.Buffer
	// var byteData string

	if err := json.Indent(&prettyJSON, []byte(data), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func Server() {
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Route => handler
	e.GET("/spec-url", func(c echo.Context) error {
		resp, err := http.Get("https://petstore.swagger.io/v2/swagger")

		if err == nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			// myString, err := PrettyJSON(body)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// fmt.Println(myString)
			return c.JSONBlob(http.StatusOK, body)
		}

		// b, _ := ioutil.ReadAll(resp.Request.Body)

		return echo.ErrInternalServerError

	})

	// e.GET("/users/:id", handler.GetUser)

	// e.POST("/users", handler.AddCat)

	// Start server
	e.StaticFS("/", ui.DistDirFS)
	e.Logger.Fatal(e.Start(":1212"))
}
