package domain

import (
	"context"

	"gorm.io/gorm"
)

type InvalidEncryption struct {
	ID                           uint `gorm:"primaryKey;autoIncrement:true"`
	CustomerID                   int  `gorm:"column:customer_id;index:idx_customer_id;uniqueIndex"`
	LegalName                    bool `gorm:"column:legal_name"`
	FullName                     bool `gorm:"column:full_name"`
	BirthPlace                   bool `gorm:"column:birth_place"`
	SurgateMotherName            bool `gorm:"column:surgate_mother_name"`
	MobilePhone                  bool `gorm:"column:mobile_phone"`
	Email                        bool `gorm:"column:email"`
	ResidenceAddress             bool `gorm:"column:residence_address"`
	ResidencePhone1              bool `gorm:"column:residence_phone1"`
	LegalAddress                 bool `gorm:"column:legal_address"`
	LegalPhone1                  bool `gorm:"column:legal_phone1"`
	CompanyAddress               bool `gorm:"column:company_address"`
	CompanyPhone1                bool `gorm:"column:company_phone1"`
	EmergencyContactName         bool `gorm:"column:emergency_contact_name"`
	EmergencyContactAddress      bool `gorm:"column:emergency_contact_address"`
	EmergencyContactHomePhone1   bool `gorm:"column:emergency_contact_home_phone1"`
	EmergencyContactMobilePhone  bool `gorm:"column:emergency_contact_mobile_phone"`
	EmergencyContactOfficePhone1 bool `gorm:"column:emergency_contact_office_phone1"`
	HasCheckedInvalidField       bool `gorm:"column:has_checked_invalid_field"`
	HasUpdateToValidField        bool `gorm:"column:has_update_to_valid_field"`
	gorm.Model
}

func (r InvalidEncryption) TableName() string {
	return "invalid_encryptions"
}

type InvalidHistoryEncryption struct {
	ID                           uint `gorm:"primaryKey;autoIncrement:true"`
	HistoryID                    int  `gorm:"column:history_id;index:idx_history_id;uniqueIndex"`
	CustomerID                   int  `gorm:"column:customer_id"`
	LegalName                    bool `gorm:"column:legal_name"`
	FullName                     bool `gorm:"column:full_name"`
	BirthPlace                   bool `gorm:"column:birth_place"`
	SurgateMotherName            bool `gorm:"column:surgate_mother_name"`
	MobilePhone                  bool `gorm:"column:mobile_phone"`
	Email                        bool `gorm:"column:email"`
	ResidenceAddress             bool `gorm:"column:residence_address"`
	ResidencePhone1              bool `gorm:"column:residence_phone1"`
	LegalAddress                 bool `gorm:"column:legal_address"`
	LegalPhone1                  bool `gorm:"column:legal_phone1"`
	CompanyAddress               bool `gorm:"column:company_address"`
	CompanyPhone1                bool `gorm:"column:company_phone1"`
	EmergencyContactName         bool `gorm:"column:emergency_contact_name"`
	EmergencyContactAddress      bool `gorm:"column:emergency_contact_address"`
	EmergencyContactHomePhone1   bool `gorm:"column:emergency_contact_home_phone1"`
	EmergencyContactMobilePhone  bool `gorm:"column:emergency_contact_mobile_phone"`
	EmergencyContactOfficePhone1 bool `gorm:"column:emergency_contact_office_phone1"`
	HasCheckedInvalidField       bool `gorm:"column:has_checked_invalid_field"`
	HasUpdateToValidField        bool `gorm:"column:has_update_to_valid_field"`
	gorm.Model
}

func (r InvalidHistoryEncryption) TableName() string {
	return "invalid_history_encryptions"
}

type InvalidEncryptionRepository interface {
	Insert(ctx context.Context, invalidEncrypt InvalidEncryption) error
	UpdateByCustomerId(ctx context.Context, field map[string]interface{}, customerId int) error
	Inserts(ctx context.Context, field []InvalidEncryption) error
}

type InvalidHistoryEncryptionRepository interface {
	Insert(ctx context.Context, invalidEncrypt InvalidHistoryEncryption) error
	UpdateByCustomerId(ctx context.Context, field map[string]interface{}, customerId int) error
	Inserts(ctx context.Context, field []InvalidHistoryEncryption) error
}
