package model

import "github.com/guregu/null"

type Bank struct {
	ID          string      `gorm:"column:id;type:uuid;primaryKey"`
	Code        string      `gorm:"column:code"`
	Name        string      `gorm:"column:name"`
	Description null.String `gorm:"column:description"` // Nullable field
	IsActive    bool        `gorm:"column:is_active"`
	Country     string      `gorm:"column:country"`
	Currency    string      `gorm:"column:currency"`
	Type        string      `gorm:"column:type"`
	Count       int         `gorm:"column:count"`
}

func (b *Bank) TableName() string {
	return "banks"
}

type BankList []Bank
