package models

import "time"

type User struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string     `gorm:"type:varchar(150);not null" json:"name"`
	Email           string     `gorm:"type:varchar(191);not null" json:"email"`
	Cpf             string     `gorm:"type:varchar(15);not null" json:"cpf"`
	Phone           string     `gorm:"type:varchar(15)" json:"phone"`
	Password        string     `gorm:"type:varchar(255);not null" json:"password"`
	Status          string     `gorm:"type:enum('pending','approved','rejected');default:'pending';not null;index" json:"status"`
	MerchantID      uint       `gorm:"index;not null" json:"merchant_id"`
	Merchant        *Merchant  `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	RememberToken   string     `gorm:"type:varchar(100)" json:"remember_token"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
