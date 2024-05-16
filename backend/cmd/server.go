package cmd

import (
	"backend/domain/repository"
	"backend/interfaces"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

const API_SERVER_DEFAULT_PORT = "8080"

type ApiServer struct {
	UserHandler interfaces.UserHandlerInterface
	AreaHandler interfaces.AreaHandlerInterface

	UserService interfaces.UserServiceInterface
	AreaService interfaces.AreaServiceInterface

	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface
	Database                repository.DatabaseInterface

	Dsn       string `json:"dsn"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	SecretKey string

	echo         *echo.Echo
	groupUserAPI *echo.Group
	groupAreaAPI *echo.Group
}

func (server *ApiServer) Run() {
	port := os.Getenv("API_SERVER_PORT")
	if port == "" {
		port = API_SERVER_DEFAULT_PORT
	}
	server.start(port)
}

func (server *ApiServer) start(port string) {
	errs := server.loadEnv()
	if len(errs) > 0 {
		log.Println(errs)
	}

	server.echo = echo.New()
	server.setMiddleware()
	server.dependencyInjection()
	server.routes()

	log.Printf("Server started at port %s", port)
	server.echo.Logger.Fatal(server.echo.Start(":" + port))
}
