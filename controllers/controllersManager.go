package controllers

import "auth-api/services"

type ControllersManager struct {
	UserController
}

// Constructor
func NewControllersManager(serviceMgr *services.ServiceManager) *ControllersManager {
	return &ControllersManager{
		*NewUserController(&serviceMgr.UserService),
	}
}
