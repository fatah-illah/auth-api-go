package repositories

import "database/sql"

type RepositoriesManager struct {
	UserRepository
}

// Constructor
func NewRepositoriesManager(dbHandler *sql.DB, jwtSecret []byte) *RepositoriesManager {
	return &RepositoriesManager{
		*NewUserRepository(dbHandler, jwtSecret),
	}
}
