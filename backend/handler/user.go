package handler

import (
	"backend/domain/model"
	"backend/interfaces"
	"log"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserServiceInterface interfaces.UserServiceInterface
}

func NewUserHandler(
	UserServiceInterface interfaces.UserServiceInterface,
) interfaces.UserHandlerInterface {
	return &UserHandler{
		UserServiceInterface: UserServiceInterface,
	}
}

func (handler *UserHandler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {

		//open connection to database
		err := handler.UserServiceInterface.DatabaseRepository().Connect()
		//fmt.Println("err", err)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection when done
		defer handler.UserServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.UserServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("GetUsers, begin transaction err:", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//return error if id_khach_hang not found
		data, err := handler.UserServiceInterface.GetUsers(tx)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		if err != nil {
			log.Println("GetUsers, err:", err)
			return c.JSON(http.StatusInternalServerError, "failed to get users")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: data,
		})
	}
}

func (handler *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_khach_hang := c.Param("id_khach_hang")

		//validate
		var isNumber = regexp.MustCompile("^[0-9]+$").MatchString
		if id_khach_hang == "" {
			return c.JSON(http.StatusBadRequest, "id_khach_hang cannot be empty")
		} else {
			if !isNumber(id_khach_hang) {
				return c.JSON(http.StatusBadRequest, "id_khach_hang only accept number: [0-9]")
			}

			if len(id_khach_hang) != 12 {
				return c.JSON(http.StatusBadRequest, "id_khach_hang must be 12 characters")
			}
		}

		//open connection to database
		err := handler.UserServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection when done
		defer handler.UserServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.UserServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("GetUser, begin transaction err:", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//return error if id_khach_hang not found
		data, err := handler.UserServiceInterface.GetUser(tx, id_khach_hang)
		if data == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		if err != nil {
			log.Println("GetUser, err:", err)
			return c.JSON(http.StatusInternalServerError, "failed to get user")
		}

		//return success
		return c.JSON(http.StatusOK, Message{
			Error:  "",
			Result: data,
		})
	}
}

func (handler *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var User model.User
		//validate request body
		if err := c.Bind(&User); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body")
		}

		//validate
		var isNumber = regexp.MustCompile("[0-9]+").MatchString
		if User.IDKhachHang == "" {
			return c.JSON(http.StatusBadRequest, "id_khach_hang cannot be empty")
		} else {
			if !isNumber(User.IDKhachHang) {
				return c.JSON(http.StatusBadRequest, "id_khach_hang only accept number: [0-9]")
			}
			if len(User.IDKhachHang) != 12 {
				return c.JSON(http.StatusBadRequest, "id_khach_hang must be 12 characters")
			}
		}

		//open connection to database
		err := handler.UserServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection when done
		defer handler.UserServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.UserServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("CreateUser, begin transaction err:", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.UserServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("CreateUser, rollback transaction err:", errRollback)
				}
			}
		}()

		//check if id_khach_hang exists
		data, err := handler.UserServiceInterface.GetUser(tx, User.IDKhachHang)
		if data != nil {
			return c.JSON(http.StatusConflict, "id_khach_hang already exists")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get user")
		}

		//create user
		err = handler.UserServiceInterface.CreateUser(tx, User)
		if err != nil {
			log.Println("CreateUser, err:", err)
			return c.JSON(http.StatusInternalServerError, "failed to create user")
		}

		//commit transaction
		err = handler.UserServiceInterface.DatabaseRepository().Commit(tx)
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

func (handler *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_khach_hang := c.Param("id_khach_hang")

		var User model.User
		//validate request body
		if err := c.Bind(&User); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body")
		}

		//validate
		var isNumber = regexp.MustCompile("[0-9]+").MatchString
		if id_khach_hang == "" {
			return c.JSON(http.StatusBadRequest, "id_khach_hang cannot be empty")
		} else {
			if !isNumber(id_khach_hang) {
				return c.JSON(http.StatusBadRequest, "id_khach_hang only accept number: [0-9]")
			}

			if len(id_khach_hang) != 12 {
				return c.JSON(http.StatusBadRequest, "id_khach_hang must be 12 characters")
			}
		}

		//open connection to database
		err := handler.UserServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection when done
		defer handler.UserServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.UserServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("UpdateUser, begin transaction err:", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.UserServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("UpdateUser, rollback transaction err:", errRollback)
				}
			}
		}()

		//check if id_khach_hang exists
		data, err := handler.UserServiceInterface.GetUser(tx, id_khach_hang)
		if data == nil {
			return c.JSON(http.StatusNotFound, "id_khach_hang not found")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get user")
		}

		//update user
		err = handler.UserServiceInterface.UpdateUser(tx, User, id_khach_hang)
		if err != nil {
			log.Println("UpdateUsers, err:", err)
			return c.JSON(http.StatusInternalServerError, "failed to update user")
		}

		//commit transaction
		err = handler.UserServiceInterface.DatabaseRepository().Commit(tx)
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

func (handler *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_khach_hang := c.Param("id_khach_hang")

		//validate
		var isNumber = regexp.MustCompile("[0-9]+").MatchString
		if id_khach_hang == "" {
			return c.JSON(http.StatusBadRequest, "id_khach_hang cannot be empty")
		} else {
			if !isNumber(id_khach_hang) {
				return c.JSON(http.StatusBadRequest, "id_khach_hang only accept number: [0-9]")
			}

			if len(id_khach_hang) != 12 {
				return c.JSON(http.StatusBadRequest, "id_khach_hang must be 12 characters")
			}
		}

		//open connection to database
		err := handler.UserServiceInterface.DatabaseRepository().Connect()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to connect to database")
		}

		//close connection when done
		defer handler.UserServiceInterface.DatabaseRepository().Close()

		//begin transaction
		tx, err := handler.UserServiceInterface.DatabaseRepository().Begin()
		if err != nil {
			log.Println("DeleteUser, begin transaction err:", err)
			return c.JSON(http.StatusInternalServerError, "something went wrong")
		}

		//defer rollback transaction if panic
		defer func() {
			if r := recover(); r != nil {
				errRollback := handler.UserServiceInterface.DatabaseRepository().Rollback(tx)
				if errRollback != nil {
					log.Println("DeleteUser, rollback transaction err:", errRollback)
				}
			}
		}()

		//check if id_khach_hang exists
		data, err := handler.UserServiceInterface.GetUser(tx, id_khach_hang)
		if data == nil {
			return c.JSON(http.StatusNotFound, "id_khach_hang not found")
		}
		if err != nil {
			log.Println("DeleteUsers, err:", err)
			return c.JSON(http.StatusInternalServerError, "failed to get user")
		}

		//delete user
		err = handler.UserServiceInterface.DeleteUser(tx, id_khach_hang)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to delete user")
		}

		//commit transaction
		err = handler.UserServiceInterface.DatabaseRepository().Commit(tx)
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

type Message struct {
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
}
