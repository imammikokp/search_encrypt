package repository

import (
	"context"
	"search_encrypt/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type invalidHistoryEncryptionRepository struct {
	db *gorm.DB
}


func NewInvalidHistoryEncryptReposiotry(db *gorm.DB) domain.InvalidHistoryEncryptionRepository{
	return &invalidHistoryEncryptionRepository{
		db:db,
	}
}

func (c invalidHistoryEncryptionRepository)Insert(ctx context.Context, invalidHistoryEncrypt domain.InvalidHistoryEncryption)error{
	db:= c.db.WithContext(ctx)
	
	return db.Model(&domain.InvalidHistoryEncryption{}).Create(&invalidHistoryEncrypt).Error
}

func (c invalidHistoryEncryptionRepository)Inserts(ctx context.Context,  field []domain.InvalidHistoryEncryption)error{
	db:= c.db.WithContext(ctx)
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&field).Error
}

func (c invalidHistoryEncryptionRepository)UpdateByCustomerId(ctx context.Context, field map[string]interface{}, customerId int)error{
	db:=c.db.WithContext(ctx)
	return db.Model(&domain.InvalidHistoryEncryption{}).Where("customer_id").Updates(field).Error
}


