package repository

import (
	"context"
	"serch_encrypt/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	
	return db.Model(&domain.InvalidEncryption{}).Create(&invalidEncrypt).Error
}

func (c invalidEncryptionRepository)Inserts(ctx context.Context,  field []domain.InvalidEncryption)error{
	db:= c.db.WithContext(ctx)
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&field).Error
}

func (c invalidEncryptionRepository)UpdateByCustomerId(ctx context.Context, field map[string]interface{}, customerId int)error{
	db:=c.db.WithContext(ctx)
	return db.Model(&domain.InvalidEncryption{}).Where("customer_id").Updates(field).Error
}


