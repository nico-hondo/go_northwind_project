package repositories

import "database/sql"

type RepositoryManager struct {
	CategoryRepository
	UsersRepository
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewCategoryRepository(dbHandler),
		*NewUsersRepository(dbHandler),
	}
}
