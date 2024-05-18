package echo

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewEcho() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger(),
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize:         1 << 10,
			DisablePrintStack: true,
			DisableStackAll:   true,
		}),
		middleware.RequestID(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
			Skipper: func(c echo.Context) bool {
				return strings.Contains(c.Request().URL.Path, "swagger")
			},
		}),
		middleware.BodyLimit("2M"),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			ErrorMessage: "",
			Timeout:      30 * time.Second,
		}),
	)

	e.Server.ReadTimeout = 15 * time.Second
	e.Server.WriteTimeout = 15 * time.Second
	e.Server.MaxHeaderBytes = 1 << 20

	return e
}
