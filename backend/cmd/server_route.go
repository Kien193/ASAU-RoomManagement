package cmd

func (server *ApiServer) routes() {
	// User API
	server.groupUserAPI = server.echo.Group("/api/v1/users")
	server.groupUserAPI.Use()
	{
		server.groupUserAPI.GET("", server.UserHandler.GetUsers())
		server.groupUserAPI.GET("/:id_khach_hang", server.UserHandler.GetUser())
		server.groupUserAPI.POST("", server.UserHandler.CreateUser())
		server.groupUserAPI.PUT("/:id_khach_hang", server.UserHandler.UpdateUser())
		server.groupUserAPI.DELETE("/:id_khach_hang", server.UserHandler.DeleteUser())
	}
}
