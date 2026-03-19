package interfaces

import "myMarket/internal/models"

type IReview interface {
	GetAllReviews() ([]models.Review, error)
}
