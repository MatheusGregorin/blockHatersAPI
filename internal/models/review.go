package models

import "time"

type Review struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string     `gorm:"type:varchar(150);not null" json:"title"`
	Consumer   string     `gorm:"type:varchar(150);not null" json:"consumer"`
	Phone      string     `gorm:"type:varchar(50)" json:"phone"`
	Stars      uint8      `gorm:"not null;index" json:"stars"`
	Comment    string     `gorm:"type:longtext;not null" json:"comment"`
	Status     string     `gorm:"type:enum('pending','approved','rejected');default:'pending';not null" json:"status"`
	MerchantID uint       `gorm:"index;not null" json:"merchant_id"`
	Merchant   *Merchant  `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
