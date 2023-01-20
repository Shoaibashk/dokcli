package api

import (
	// "browser"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/browser"
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

// func PrettyJSON[T pretty](data T) (string, error) {
// 	var prettyJSON bytes.Buffer
// 	switch data. {
// 	case string:
// 		{
// 			if err := json.Indent(&prettyJSON, []byte(data), "", "    "); err != nil {
// 				return "", err
// 			}
// 		}

// 		// case :
// 		// 	{
// 		// 		if err := json.Indent(&prettyJSON, data, "", "    "); err != nil {
// 		// 			return "", err
// 		// 		}
// 		// 	}

// 	}
// 	return prettyJSON.String(), nil
// }

func Server() {
	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/health", HealthCheck)

	// Route => handler
	e.GET("/spec-url", func(c echo.Context) error {
		resp, err := http.Get("https://api.eu.urbanstreet.com/delivery/swagger/v1/swagger.json")

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
	// browser.Open("http://localhost:1212/")
	// go open("http://localhost:1212/")
	go func() {
		<-time.After(100 * time.Millisecond)
		browser.OpenURL("http://localhost:1212/")
	}()
	e.Start(":1212")
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
