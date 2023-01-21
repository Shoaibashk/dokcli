package server

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/browser"
	"github.com/shoaibashk/dokcli/ui"
)

type Dokcli struct {
	Server      *echo.Echo
	Port        string
	HideBanner  bool
	OpenBrowser bool
}

func (d Dokcli) New(port string, hideBanner bool, openBrowser bool) *Dokcli {
	return &Dokcli{
		Server:      echo.New().AcquireContext().Echo(),
		Port:        port,
		HideBanner:  hideBanner,
		OpenBrowser: openBrowser,
	}
}

func (d *Dokcli) Middleware() {

	d.Server.Use(middleware.Logger())
	d.Server.Use(middleware.Recover())
	d.Server.Use(middleware.CORS())

	// fmt.Println(d.HideBanner, d.OpenBrowser, d.Port)
}

func (d *Dokcli) Routing() {
	// Routes
	d.Server.GET("/health", HealthCheck)

	// Route => handler
	d.Server.GET("/spec-url", func(c echo.Context) error {
		resp, err := http.Get("https://api.eu.urbanstreet.com/delivery/swagger/v1/swagger.json")

		if err == nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			return c.JSONBlob(http.StatusOK, body)
		}

		return echo.ErrInternalServerError

	})

}

func (d *Dokcli) Register() {
	d.Server.HideBanner = d.HideBanner
	d.Server.HidePort = true

	d.Server.Logger.SetLevel(log.INFO)

}
func (d *Dokcli) StartServer() {
	// Start server
	d.Server.StaticFS("/", ui.DistDirFS)
	go func() {
		<-time.After(100 * time.Millisecond)
		browser.OpenURL("http://localhost:1212/")
	}()
	if err := d.Server.Start(d.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Print("Application stopped gracefully")
	}

}

func (d *Dokcli) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := d.Server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Warn("application shutdowned")
		fmt.Println("Application shutdowned")
	}
}

func CreateChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
