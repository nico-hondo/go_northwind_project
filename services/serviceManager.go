package services

import "codeid.northwind/repositories"

type ServiceManager struct {
	CategoryService
	UsersService
}

// constructor baru
func NewServiceManager(repoMgr *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		// CategoryService: *NewCategoryService(&repoMgr.CategoryRepository),
		CategoryService: *NewCategoryService(repoMgr),
		UsersService:    *NewUsersService(repoMgr), //&repoMgr.UserRepository,
	}
}
