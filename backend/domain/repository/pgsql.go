package repository

type DatabaseRepositoryPGSQLInterface interface {
	DatabaseInterface
	ApiUser() UserRepositoryInterface
	ApiArea() AreaRepositoryInterface
}
