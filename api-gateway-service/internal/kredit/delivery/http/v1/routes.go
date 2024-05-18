package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *kreditHandler) MapRoutes() {
	h.group.POST("/konsumen", h.CreateKonsumen())
	h.group.POST("/limit", h.CreateLimit())
	h.group.GET("/limit/:id", h.GetLimit())
	h.group.POST("/transaksi", h.CreateTransaksi())
	h.group.GET("/transaksi/:id", h.GetTransaksi())
	h.group.Any("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
