package products

import "github.com/labstack/echo/v4"

type HttpDelivery interface {
	CreateKonsumen() echo.HandlerFunc
	CreateLimit() echo.HandlerFunc
	CreateTransaksi() echo.HandlerFunc

	GetLimit() echo.HandlerFunc
	GetTransaksi() echo.HandlerFunc
}
