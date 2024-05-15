package interfaces

import "github.com/labstack/echo/v4"

type UserHandlerInterface interface {
	GetUsers() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}
