package repository

import (
	"fmt"
	"myMarket/internal/database"
	"myMarket/internal/models"
)

type MerchantMysqlRepository struct{}

func NewMerchantMysqlRepository() *MerchantMysqlRepository {
	return &MerchantMysqlRepository{}
}

func (mr *MerchantMysqlRepository) GetAllMerchants() ([]models.Merchant, error) {
	var merchants []models.Merchant
	result := database.DB.Find(&merchants)
	if result.Error != nil {
		return nil, fmt.Errorf("could not fetch merchants")
	}

	return merchants, nil
}

func (mr *MerchantMysqlRepository) CreateMerchant(merchant *models.Merchant) (*models.Merchant, error) {
	var existing models.Merchant
	if err := database.DB.Where("cnpj = ?", merchant.Cnpj).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("merchant with this CNPJ already exists")
	}

	if err := database.DB.Create(merchant).Error; err != nil {
		return nil, fmt.Errorf("failed to create merchant: %v", err)
	}

	return merchant, nil
}
