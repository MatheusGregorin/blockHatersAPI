package interfaces

import "myMarket/internal/models"

type IMerchant interface {
	GetAllMerchants() ([]models.Merchant, error)
	CreateMerchant(merchant *models.Merchant) (*models.Merchant, error)
}
