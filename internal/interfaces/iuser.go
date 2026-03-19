package interfaces

import "myMarket/internal/models"

type IUser interface {
	Register(user *models.User) (*models.User, error)
	Login(email string, password string) (string, error)
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}
