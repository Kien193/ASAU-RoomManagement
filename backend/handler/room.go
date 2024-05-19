package handler

import (
	"backend/domain/model"
	"backend/interfaces"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	RoomServiceInterface interfaces.RoomServiceInterface
}

func NewRoomHandler(
	RoomServiceInterface interfaces.RoomServiceInterface,
) interfaces.RoomHandlerInterface {
	return &RoomHandler{
		RoomServiceInterface: RoomServiceInterface,
	}
}

func (handler *RoomHandler) GetRooms() echo.HandlerFunc {
	return func(c echo.Context) error {

		//open connection to database
		err := handler.RoomServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.RoomServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.RoomServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("GetRooms, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//return error if id_phong not found
		data, err := handler.RoomServiceInterface.GetRooms(tx)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		if err != nil {
			log.Println("GetRooms, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to get rooms")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: data,
		})
	}
}

func (handler *RoomHandler) GetRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_phong := c.Param("id_phong")

		//validate

		//open connection to database
		err := handler.RoomServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.RoomServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.RoomServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("GetRoom, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//return error if id_phong not found
		data, err := handler.RoomServiceInterface.GetRoom(tx, id_phong)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		if err != nil {
			log.Println("GetRoom, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to get room")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: data,
		})
	}
}

func (handler *RoomHandler) CreateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Room model.Room
		if err := c.Bind(&Room); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body")
		}

		//validate

		//open connection to database
		err := handler.RoomServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.RoomServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.RoomServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("CreateRoom, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.RoomServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("CreateRoom, rollback transaction err:", errRollback)
				}
			}
		}()

		//check if id_phong already exists
		data, err := handler.RoomServiceInterface.GetRoom(tx, Room.IDPhong)
		if data != nil {
			return c.JSON(http.StatusConflict, "id_phong already exists")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get room")
		}

		//create room
		err = handler.RoomServiceInterface.CreateRoom(tx, Room)
		if err != nil {
			log.Println("CreateRoom, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to create room")
		}

		//commit transaction
		err = handler.RoomServiceInterface.DatabaseRepository().Commit(tx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to commit transaction")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: nil,
		})
	}
}

func (handler *RoomHandler) UpdateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_phong := c.Param("id_phong")

		var Room model.Room
		//validate request body
		if err := c.Bind(&Room); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body")
		}

		//validate

		//open connection to database
		err := handler.RoomServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.RoomServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.RoomServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("UpdateRoom, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.RoomServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("UpdateRoom, rollback transaction err:", errRollback)
				}
			}
		}()

		//check if id_phong not found
		data, err := handler.RoomServiceInterface.GetRoom(tx, id_phong)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get room")
		}

		//update room
		err = handler.RoomServiceInterface.UpdateRoom(tx, Room, id_phong)
		if err != nil {
			log.Println("UpdateRoom, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to update room")
		}

		//commit transaction
		err = handler.RoomServiceInterface.DatabaseRepository().Commit(tx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to commit transaction")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: nil,
		})
	}
}

func (handler *RoomHandler) DeleteRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_phong := c.Param("id_phong")

		//validate

		//open connection to database
		err := handler.RoomServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.RoomServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.RoomServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("DeleteRoom, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.RoomServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("DeleteRoom, rollback transaction err:", errRollback)
				}
			}
		}()

		//return error if id_phong not found
		data, err := handler.RoomServiceInterface.GetRoom(tx, id_phong)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get room")
		}

		//delete room
		err = handler.RoomServiceInterface.DeleteRoom(tx, id_phong)
		if err != nil {
			log.Println("DeleteRoom, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to delete room")
		}

		//commit transaction
		err = handler.RoomServiceInterface.DatabaseRepository().Commit(tx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to commit transaction")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: nil,
		})
	}
}
