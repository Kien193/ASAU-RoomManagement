package interfaces

import "github.com/labstack/echo/v4"

type RoomHandlerInterface interface {
	GetRooms() echo.HandlerFunc
	GetRoom() echo.HandlerFunc
	CreateRoom() echo.HandlerFunc
	UpdateRoom() echo.HandlerFunc
	DeleteRoom() echo.HandlerFunc
}
