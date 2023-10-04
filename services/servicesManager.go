package services

import "auth-api/repositories"

type ServiceManager struct {
	UserService
}

// Constructor
func NewServicesManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		UserService: *NewUserService(&repoMgr.UserRepository),
	}
}
