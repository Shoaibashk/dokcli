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
	"github.com/mbndr/figlet4go"
	"github.com/pkg/browser"
	"github.com/shoaibashk/dokcli/ui"
)

type Dokcli struct {
	Server      *echo.Echo
	Port        string
	HideBanner  bool
	OpenBrowser bool
}

func NewServer(hideBanner bool, openBrowser bool) *Dokcli {
	return &Dokcli{
		Server:      echo.New().AcquireContext().Echo(),
		HideBanner:  hideBanner,
		OpenBrowser: openBrowser,
	}
}
func (d *Dokcli) SetPort(port string) string {
	d.Port = ":" + port
	return port
}

func (d *Dokcli) Middleware() {

	d.Server.Use(middleware.Logger())
	d.Server.Use(middleware.Recover())
	d.Server.Use(middleware.CORS())

	// fmt.Println(d.HideBanner, d.OpenBrowser, d.Port)
}

func (d *Dokcli) Routing(url string) {
	// Routes
	d.Server.GET("/health", HealthCheck)

	// Route => handler
	d.Server.GET("/spec-url", func(c echo.Context) error {
		resp, err := http.Get(url)
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
		url := "http://localhost" + d.Port + "/"
		fmt.Println(url)
		<-time.After(100 * time.Millisecond)
		browser.OpenURL(url)
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

func DokcliBanner() (string, error) {
	ascii := figlet4go.NewAsciiRender()

	bannerOptions := figlet4go.NewRenderOptions()
	bannerOptions.FontName = "larry3d"
	bannerOptions.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorGreen,
		figlet4go.ColorRed,
		figlet4go.ColorCyan,

		// ...or by an hex string...
		// figlet4go.NewTrueColorFromHexString("885DBA"),
		// ...or by an TrueColor object with rgb values
		// figlet4go.TrueColor{136, 93, 186},

	}

	// figlet4go.TrueColor{255, 198, 211}

	return ascii.RenderOpts("Dok Cli", bannerOptions)
}
