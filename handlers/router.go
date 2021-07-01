package handlers

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/phuwn/coinchameleon/data/db"
)

// Router - handling routes for incoming request
func Router() *echo.Echo {
	r := echo.New()
	r.Pre(mw.RemoveTrailingSlash())

	r.GET("/healthz", healthz)
	{
		r.GET("coins/:id/price", getPriceData)
	}

	return r
}

func healthz(c echo.Context) error {
	err := db.Healthz()
	if err != nil {
		return err
	}
	return c.String(200, "ok")
}
