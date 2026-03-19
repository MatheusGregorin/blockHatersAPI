package repository

import (
	"fmt"
	"myMarket/internal/database"
	"myMarket/internal/models"
)

type ReviewMysqlRepository struct{}

func NewReviewMysqlRepository() *ReviewMysqlRepository {
	return &ReviewMysqlRepository{}
}

func (rr *ReviewMysqlRepository) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	result := database.DB.Preload("Merchant").Find(&reviews)
	if result.Error != nil {
		return nil, fmt.Errorf("could not fetch reviews")
	}

	return reviews, nil
}
