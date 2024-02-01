package repository

import (
	"context"
	"search_encrypt/domain"

	"gorm.io/gorm"
)

type searchHistoryEncryptRepository struct {
	db *gorm.DB
}

func NewSearchHistoryEncryptRepository(db *gorm.DB) domain.SearchHistoryEncryptRepository {
	return &searchHistoryEncryptRepository{
		db: db,
	}
}

func (c searchHistoryEncryptRepository) FetchByRange(ctx context.Context, model interface{}, search []int, minId int, maxId int) error {
	db := c.db.WithContext(ctx)
	db = db.Select(
		"customers.customer_id AS customers_id",
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
	)

	if search != nil {
		if len(search) != 0 {

			var idRangeLoad []int
			for _, v := range search {
				if v >= minId && v <= maxId {
					idRangeLoad = append(idRangeLoad, v)
				}
			}
			db = db.Where(`id in (select a.id from customers a where id >= ? AND id <= ? EXCEPT SELECT b.id FROM customers b where id in(?))`, minId, maxId, idRangeLoad)

		} else {
			db = db.Where("id >= ? AND  id <= ?", minId, maxId)
		}
	} else {
		db = db.Where("id >= ? AND  id <= ?", minId, maxId)
	}

	return db.Find(model).Error

}

func (c searchHistoryEncryptRepository) FindById(ctx context.Context, model interface{}, id int) error {
	db := c.db.WithContext(ctx)
	return db.Select(
		"customers.customer_id",
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
	).Where("id = ?", id).First(model).Error
}

func (c searchHistoryEncryptRepository) GetCountAll(ctx context.Context) (int64, error) {
	db := c.db.WithContext(ctx)
	var count int64
	return count, db.Model(&domain.EncryptHistoryCustomer{}).Count(&count).Error
}
