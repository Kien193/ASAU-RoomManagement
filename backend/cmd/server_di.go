package cmd

import (
	"ASAU-user-api/handler"
	"ASAU-user-api/infrastructure"
	"ASAU-user-api/service"

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
