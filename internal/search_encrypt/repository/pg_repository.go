package repository

import (
	"context"
	"serch_encrypt/domain"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository{
	return &customerRepository{
		db:db,
	}
}

func (c customerRepository)FetchByRange(ctx context.Context, model interface{} ,minId int, maxId int )error{
	db := c.db.WithContext(ctx)
	return	db.Select(
		"(SELECT DEC_B64('SEC',customers.legal_name) AS legal_name)",
		"(SELECT DEC_B64('SEC',customers.full_name) AS full_name)",
		"(SELECT DEC_B64('SEC',customers.birth_place) AS birth_place)",
		"(SELECT DEC_B64('SEC',customers.surgate_mother_name) AS surgate_mother_name)",
		"(SELECT DEC_B64('SEC',customers.mobile_phone) AS mobile_phone)",
		"(SELECT DEC_B64('SEC',customers.email) AS email)",
		"(SELECT DEC_B64('SEC',customers.residence_address) AS residence_address)",
		"(SELECT DEC_B64('SEC',customers.residence_phone1) AS residence_phone1)",
		"(SELECT DEC_B64('SEC',customers.legal_address) AS legal_address)",
		"(SELECT DEC_B64('SEC',customers.legal_phone1) AS legal_phone1)",
		"(SELECT DEC_B64('SEC',customers.company_address) AS company_address)",
		"(SELECT DEC_B64('SEC',customers.company_phone1) AS company_phone1)",
		"(SELECT DEC_B64('SEC',customers.emergency_contact_name) AS emergency_contact_name)",
		"(SELECT DEC_B64('SEC',customers.emergency_contact_address) AS emergency_contact_address)",
		"(SELECT DEC_B64('SEC',customers.emergency_contact_home_phone1) AS emergency_contact_home_phone1)",
		"(SELECT DEC_B64('SEC',customers.emergency_contact_office_phone1) AS emergency_contact_office_phone1)",
		"(SELECT DEC_B64('SEC',customers.emergency_contact_mobile_phone) AS emergency_contact_mobile_phone)",
		).Where("id >= ? AND id id <= ?", minId,maxId ).Model(&model).Error
}

func (c customerRepository)GetCountAll(ctx context.Context,Count int64) (error){
	db := c.db.WithContext(ctx)
	return db.Model(&domain.Customer{}).Count(&Count).Error
}

