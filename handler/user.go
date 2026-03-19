package handler

import (
	"fmt"
	"myMarket/internal/interfaces"
	"myMarket/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserHandler struct {
	repo interfaces.IUser
}

func NewUserHandler(r interfaces.IUser) *UserHandler {
	return &UserHandler{
		repo: r,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	userSaved, err := h.repo.Register(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User registered successfully",
		"user":    userSaved,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var loginReq LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.repo.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message":    "User logged in successfully",
		"token":      token,
		"type":       "Bearer",
		"expired_in": time.Now().Add(time.Hour * 24).Unix(),
	})
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("Erro na conversão:", err)
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.repo.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User retrieved successfully", "user": user})
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Users retrieved successfully", "users": users})
}
