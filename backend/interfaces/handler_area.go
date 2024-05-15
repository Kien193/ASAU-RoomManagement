package interfaces

import "github.com/labstack/echo/v4"

type AreaHandlerInterface interface {
	GetAreas() echo.HandlerFunc
	GetArea() echo.HandlerFunc
	CreateArea() echo.HandlerFunc
	UpdateArea() echo.HandlerFunc
	DeleteArea() echo.HandlerFunc
}
