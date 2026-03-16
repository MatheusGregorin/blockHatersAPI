package handler

import (
	"fmt"
	"myMarket/internal/interfaces"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=10"`
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
	var loginReq LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	userSaved, err := h.repo.Register(loginReq.Username, loginReq.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User registered successfully",
		"user":    userSaved,
	})
	return
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var loginReq LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
	}

	token, err := h.repo.Login(loginReq.Username, loginReq.Password)
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
