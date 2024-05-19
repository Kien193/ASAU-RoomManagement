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

	// Area API
	server.groupAreaAPI = server.echo.Group("/api/v1/areas")
	server.groupAreaAPI.Use()
	{
		server.groupAreaAPI.GET("", server.AreaHandler.GetAreas())
		server.groupAreaAPI.GET("/:id_khu_vuc", server.AreaHandler.GetArea())
		server.groupAreaAPI.POST("", server.AreaHandler.CreateArea())
		server.groupAreaAPI.PUT("/:id_khu_vuc", server.AreaHandler.UpdateArea())
		server.groupAreaAPI.DELETE("/:id_khu_vuc", server.AreaHandler.DeleteArea())
	}

	// Room API
	server.groupRoomAPI = server.echo.Group("/api/v1/rooms")
	server.groupRoomAPI.Use()
	{
		server.groupRoomAPI.GET("", server.RoomHandler.GetRooms())
		server.groupRoomAPI.GET("/:id_phong", server.RoomHandler.GetRoom())
		server.groupRoomAPI.POST("", server.RoomHandler.CreateRoom())
		server.groupRoomAPI.PUT("/:id_phong", server.RoomHandler.UpdateRoom())
		server.groupRoomAPI.DELETE("/:id_phong", server.RoomHandler.DeleteRoom())
	}
}
