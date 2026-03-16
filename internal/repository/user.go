package repository

import (
	"fmt"
	"myMarket/internal/database"
	"myMarket/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserMysqlRepository struct{}

func NewUserMysqlRepository() *UserMysqlRepository {
	return &UserMysqlRepository{}
}

func (ur *UserMysqlRepository) Register(username string, password string) (*models.User, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error == nil {
		return nil, fmt.Errorf("username already exists")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generating hash password")
	}

	// Verifica se a senha fornecida corresponde ao hash armazenado
	if err := bcrypt.CompareHashAndPassword(hashPass, []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	user.Username = username
	user.Password = string(hashPass)
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	user.Password = "" // Limpa a senha antes de retornar o usuário

	return &user, nil
}

func (ur *UserMysqlRepository) Login(username string, password string) (string, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("user not found")
	}

	// Verifica se a senha fornecida corresponde ao hash armazenado
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expira em 24h
	})

	var jwtKey = []byte(os.Getenv("TOKEN_JWT"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ur *UserMysqlRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found")
	}

	user.Password = "" // Limpa a senha antes de retornar o usuário

	return &user, nil
}
