package interfaces

import "myMarket/internal/models"

type IUser interface {
	Register(username string, password string) (*models.User, error)
	Login(username string, password string) (string, error)
	GetUserByID(id uint) (*models.User, error)
}
