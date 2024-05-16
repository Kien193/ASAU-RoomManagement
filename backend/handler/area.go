package handler

import (
	"backend/domain/model"
	"backend/interfaces"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AreaHandler struct {
	AreaServiceInterface interfaces.AreaServiceInterface
}

func NewAreaHandler(
	AreaServiceInterface interfaces.AreaServiceInterface,
) interfaces.AreaHandlerInterface {
	return &AreaHandler{
		AreaServiceInterface: AreaServiceInterface,
	}
}

func (handler *AreaHandler) GetAreas() echo.HandlerFunc {
	return func(c echo.Context) error {

		//open connection to database
		err := handler.AreaServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.AreaServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.AreaServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("GetAreas, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//return error if id_khu_vuc not found
		data, err := handler.AreaServiceInterface.GetAreas(tx)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		if err != nil {
			log.Println("GetAreas, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to get areas")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: data,
		})
	}
}

func (handler *AreaHandler) GetArea() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_khu_vuc := c.Param("id_khu_vuc")

		//validate

		//open connection to database
		err := handler.AreaServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.AreaServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.AreaServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("GetArea, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//return error if id_khu_vuc not found
		data, err := handler.AreaServiceInterface.GetArea(tx, id_khu_vuc)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		if err != nil {
			log.Println("GetArea, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to get area")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: data,
		})
	}
}

func (handler *AreaHandler) CreateArea() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Area model.Area
		if err := c.Bind(&Area); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid Body")
		}

		//validate

		//open connection to database
		err := handler.AreaServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.AreaServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.AreaServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("CreateArea, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.AreaServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("CreateUser, rollback transaction err:", errRollback)
				}
			}
		}()

		//check if id_khu_vuc already exists
		data, err := handler.AreaServiceInterface.GetArea(tx, Area.IDKhuVuc)
		if data != nil {
			return c.JSON(http.StatusConflict, "id_khu_vuc already exists")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get area")
		}

		//create area
		err = handler.AreaServiceInterface.CreateArea(tx, Area)
		if err != nil {
			log.Println("CreateArea, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to create area")
		}

		//commit transaction
		err = handler.AreaServiceInterface.DatabaseRepository().Commit(tx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to commit transaction")
		}

		//return success
		return c.JSON(http.StatusCreated, Message{
			Error:  "",
			Result: nil,
		})
	}
}

func (handler *AreaHandler) UpdateArea() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_khu_vuc := c.Param("id_khu_vuc")

		var Area model.Area
		//validate request body
		if err := c.Bind(&Area); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid Body")
		}

		//validate

		//open connection to database
		err := handler.AreaServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.AreaServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.AreaServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("UpdateArea, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.AreaServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("UpdateArea, rollback transaction err:", errRollback)
				}
			}
		}()

		//return error if id_khu_vuc not found
		data, err := handler.AreaServiceInterface.GetArea(tx, Area.IDKhuVuc)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}
		if err != nil {
			log.Println("UpdateArea, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to get area")
		}

		//update area
		err = handler.AreaServiceInterface.UpdateArea(tx, Area, id_khu_vuc)
		if err != nil {
			log.Println("UpdateArea, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to update area")
		}

		//commit transaction
		err = handler.AreaServiceInterface.DatabaseRepository().Commit(tx)
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

func (handler *AreaHandler) DeleteArea() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_khu_vuc := c.Param("id_khu_vuc")

		//validate

		//open connection to database
		err := handler.AreaServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection to database
		defer handler.AreaServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.AreaServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("DeleteArea, begin transaction err: ", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.AreaServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("DeleteArea, rollback transaction err:", errRollback)
				}
			}
		}()

		//return error if id_khu_vuc not found
		data, err := handler.AreaServiceInterface.GetArea(tx, id_khu_vuc)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}
		if err != nil {
			log.Println("DeleteArea, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to get area")
		}

		//delete area
		err = handler.AreaServiceInterface.DeleteArea(tx, id_khu_vuc)
		if err != nil {
			log.Println("DeleteArea, err: ", err)
			return c.JSON(http.StatusInternalServerError, "failed to delete area")
		}

		//commit transaction
		err = handler.AreaServiceInterface.DatabaseRepository().Commit(tx)
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
