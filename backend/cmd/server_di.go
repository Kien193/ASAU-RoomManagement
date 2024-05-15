package cmd

import (
	"backend/handler"
	"backend/infrastructure"
	"backend/service"

	_ "github.com/lib/pq"
)

func (server *ApiServer) dependencyInjection() {
	server.Database = infrastructure.NewDatabase(
		server.Dsn,
		server.Username,
		server.Password,
	)
	server.DatabaseRepositoryPGSQL = infrastructure.NewDatabaseRepositoryPGSQL(server.Database)
	server.UserService = service.NewUserService(server.DatabaseRepositoryPGSQL)
	server.UserHandler = handler.NewUserHandler(server.UserService)
}
