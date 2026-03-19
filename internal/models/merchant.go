package models

import "time"

type Merchant struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(150);not null" json:"name"`
	Cnpj      string     `gorm:"type:varchar(18);not null" json:"cnpj"`
	Status    string     `gorm:"type:enum('pending','approved','rejected');default:'pending';not null;index" json:"status"`
	Plan      string     `gorm:"type:varchar(255);not null" json:"plan"`
	ValuePlan string     `gorm:"type:varchar(255);not null" json:"value_plan"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
