package handler

import (
	"myMarket/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	repo interfaces.IReview
}

func NewReviewHandler(r interfaces.IReview) *ReviewHandler {
	return &ReviewHandler{
		repo: r,
	}
}

func (h *ReviewHandler) GetAllReviews(ctx *gin.Context) {
	reviews, err := h.repo.GetAllReviews()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve reviews"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Reviews retrieved successfully", "reviews": reviews})
}
