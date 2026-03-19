package handler

import (
	"myMarket/internal/interfaces"
	"myMarket/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateMerchantRequest struct {
	Name      string `json:"name" binding:"required"`
	Cnpj      string `json:"cnpj" binding:"required,max=18"`
	Status    string `json:"status" binding:"omitempty,oneof=pending approved rejected"`
	Plan      string `json:"plan" binding:"required"`
	ValuePlan string `json:"value_plan" binding:"required"`
}

type MerchantHandler struct {
	repo interfaces.IMerchant
}

func NewMerchantHandler(r interfaces.IMerchant) *MerchantHandler {
	return &MerchantHandler{
		repo: r,
	}
}

func (h *MerchantHandler) GetAllMerchants(ctx *gin.Context) {
	merchants, err := h.repo.GetAllMerchants()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve merchants"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Merchants retrieved successfully", "merchants": merchants})
}

func (h *MerchantHandler) CreateMerchant(ctx *gin.Context) {
	var req CreateMerchantRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	merchant := models.Merchant{
		Name:      req.Name,
		Cnpj:      req.Cnpj,
		Status:    req.Status,
		Plan:      req.Plan,
		ValuePlan: req.ValuePlan,
	}

	if merchant.Status == "" {
		merchant.Status = "pending"
	}

	created, err := h.repo.CreateMerchant(&merchant)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message":  "Merchant created successfully",
		"merchant": created,
	})
}
