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

func (ur *UserMysqlRepository) Register(user *models.User) (*models.User, error) {
	var existingUser models.User

	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return nil, fmt.Errorf("email already exists")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generating hash password")
	}

	user.Password = string(hashPass)
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}

	user.Password = "" // Limpa a senha antes de retornar o usuário

	return user, nil
}

func (ur *UserMysqlRepository) Login(email string, password string) (string, error) {
	var user models.User

	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("user not found")
	}

	// Verifica se a senha fornecida corresponde ao hash armazenado
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Printf("Erro ao verificar senha! Banco: [%s] | Recebido: [%s] | Erro: %v\n", user.Password, password, err)
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

func (ur *UserMysqlRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Preload("Merchant").Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("could not fetch users")
	}

	// Limpa a senha de todos os usuários
	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}
