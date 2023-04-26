package domain

import (
	"context"

	"gorm.io/gorm"
)

type InvalidEncryption struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID int  `gorm:"primarykey;autoIncrement:true"`
	gorm.Model
}

type InvalidEncryptionRepository interface{
	Insert(ctx context.Context, invalidEncrypt InvalidEncryption)error
}

