package repository

type DatabaseRepositoryPGSQLInterface interface {
	DatabaseInterface
	ApiUser() UserRepositoryInterface
	ApiArea() AreaRepositoryInterface
	ApiRoom() RoomRepositoryInterface
}
