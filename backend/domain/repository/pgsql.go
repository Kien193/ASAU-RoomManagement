package repository

type DatabaseRepositoryPGSQLInterface interface {
	DatabaseInterface
	ApiUser() UserRepositoryInterface
}
