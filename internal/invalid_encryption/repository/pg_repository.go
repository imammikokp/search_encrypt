package repository

import (
	"context"
	"serch_encrypt/domain"

	"gorm.io/gorm"
)

type invalidEncryptionRepository struct {
	db *gorm.DB
}


func NewInvalidEncryptReposiotry(db *gorm.DB) domain.InvalidEncryptionRepository{
	return &invalidEncryptionRepository{
		db:db,
	}
}

func (c invalidEncryptionRepository)Insert(ctx context.Context, invalidEncrypt domain.InvalidEncryption)error{
	db:= c.db.WithContext(ctx)
	return db.Create(&invalidEncrypt).Error
}

