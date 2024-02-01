package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type Customer struct {
	ID                                 int             `gorm:"primarykey;autoIncrement:true"`
	MobilePhone                        string          `gorm:"type:varchar(50);column:mobile_phone"`
	Email                              string          `gorm:"type:varchar(100);column:email;index:idx_customers_varchar;index:idx_customers_email,unique"`
	Religion                           string          `gorm:"type:varchar(10);column:religion"`
	MaritalStatus                      string          `gorm:"type:varchar(10);column:marital_status"`
	NumOfDependence                    int             `gorm:"column:num_of_dependence"`
	NoKK                               string          `gorm:"type:varchar(20);column:no_kk"`
	Education                          string          `gorm:"type:varchar(10);column:education"`
	ProfessionId                       string          `gorm:"type:varchar(10);column:profession_id"`
	Nationality                        string          `gorm:"type:char(3);column:nationality"`
	WNACountry                         string          `gorm:"type:varchar(50);column:wna_country"`
	HomeStatus                         string          `gorm:"type:varchar(10);column:home_status"`
	RentFinishDate                     sql.NullTime    `gorm:"column:rent_finish_date"`
	HomeLocation                       string          `gorm:"type:varchar(10);column:home_location"`
	HomePrice                          sql.NullFloat64 `gorm:"column:home_price"`
	StaySinceYear                      sql.NullInt32   `gorm:"column:stay_since_year"`
	StaySinceMonth                     sql.NullInt32   `gorm:"column:stay_since_month"`
	NumOfAssetOwned                    int             `gorm:"column:num_of_asset_owned"`
	LivingCostAmount                   sql.NullFloat64 `gorm:"column:living_cost_amount"`
	BankID                             string          `gorm:"type:varchar(5);column:bank_id"`
	AccountNo                          string          `gorm:"type:varchar(20);column:account_no"`
	AccountName                        string          `gorm:"type:varchar(50);column:account_name"`
	CreatedAt                          time.Time       `gorm:"column:created_at"`
	CreatedBy                          string          `gorm:"type:varchar(255);column:created_by"`
	UpdatedAt                          time.Time       `gorm:"column:updated_at"`
	UpdatedBy                          string          `gorm:"type:varchar(255);column:updated_by"`
	DeletedAt                          gorm.DeletedAt  `gorm:"index"`
	ResidenceAddress                   string          `gorm:"type:varchar(200);column:residence_address"`
	ResidenceRt                        string          `gorm:"type:char(3);column:residence_rt"`
	ResidenceRw                        string          `gorm:"type:char(3);column:residence_rw"`
	ResidenceKelurahan                 string          `gorm:"type:varchar(30);column:residence_kelurahan"`
	ResidenceKecamatan                 string          `gorm:"type:varchar(30);column:residence_kecamatan"`
	ResidenceCity                      string          `gorm:"type:varchar(30);column:residence_city"`
	ResidenceZipCode                   string          `gorm:"type:char(5);column:residence_zip_code"`
	ResidenceAreaPhone1                string          `gorm:"type:varchar(4);column:residence_area_phone1"`
	ResidencePhone1                    string          `gorm:"type:varchar(50);column:residence_phone1"`
	LegalAddress                       string          `gorm:"type:varchar(200);column:legal_address"`
	LegalRt                            string          `gorm:"type:char(3);column:legal_rt"`
	LegalRw                            string          `gorm:"type:char(3);column:legal_rw"`
	LegalKelurahan                     string          `gorm:"type:varchar(30);column:legal_kelurahan"`
	LegalKecamatan                     string          `gorm:"type:varchar(30);column:legal_kecamatan"`
	LegalCity                          string          `gorm:"type:varchar(30);column:legal_city"`
	LegalZipCode                       string          `gorm:"type:char(5);column:legal_zip_code"`
	LegalAreaPhone1                    string          `gorm:"type:varchar(4);column:legal_area_phone1"`
	LegalPhone1                        string          `gorm:"type:varchar(50);column:legal_phone1"`
	CompanyName                        string          `gorm:"type:varchar(64);column:company_name"`
	CompanyAddress                     string          `gorm:"type:varchar(200);column:company_address"`
	CompanyRt                          string          `gorm:"type:char(3);column:company_rt"`
	CompanyRw                          string          `gorm:"type:char(3);column:company_rw"`
	CompanyKelurahan                   string          `gorm:"type:varchar(30);column:company_kelurahan"`
	CompanyKecamatan                   string          `gorm:"type:varchar(30);column:company_kecamatan"`
	CompanyCity                        string          `gorm:"type:varchar(30);column:company_city"`
	CompanyZipCode                     string          `gorm:"type:char(5);column:company_zip_code"`
	CompanyAreaPhone1                  string          `gorm:"type:varchar(10);column:company_area_phone1"`
	CompanyPhone1                      string          `gorm:"type:varchar(50);column:company_phone1"`
	JobPosition                        string          `gorm:"type:varchar(10);column:job_position"`
	JobType                            string          `gorm:"type:varchar(10);column:job_type"`
	IndustryTypeId                     string          `gorm:"type:varchar(10);column:industry_type_id"`
	MainBusinessSinceYear              sql.NullInt32   `gorm:"column:main_business_since_year"`
	EmploymentSinceMonth               sql.NullInt32   `gorm:"column:employment_since_month"`
	EmploymentSinceYear                sql.NullInt32   `gorm:"column:employment_since_year"`
	MonthlyFixedIncome                 sql.NullFloat64 `gorm:"column:monthly_fixed_income"`
	MonthlyVariableIncome              sql.NullFloat64 `gorm:"column:monthly_variable_income"`
	SpouseIncome                       sql.NullFloat64 `gorm:"column:spouse_income"`
	EmergencyContactName               string          `gorm:"type:varchar(50);column:emergency_contact_name"`
	EmergencyContactRelationship       string          `gorm:"type:varchar(10);column:emergency_contact_relationship"`
	EmergencyContactAddress            string          `gorm:"type:varchar(200);column:emergency_contact_address"`
	EmergencyContactRt                 string          `gorm:"type:char(3);column:emergency_contact_rt"`
	EmergencyContactRw                 string          `gorm:"type:char(3);column:emergency_contact_rw"`
	EmergencyContactKelurahan          string          `gorm:"type:varchar(30);column:emergency_contact_kelurahan"`
	EmergencyContactKecamatan          string          `gorm:"type:varchar(30);column:emergency_contact_kecamatan"`
	EmergencyContactCity               string          `gorm:"type:varchar(30);column:emergency_contact_city"`
	EmergencyContactZipCode            string          `gorm:"type:char(5);column:emergency_contact_zip_code"`
	EmergencyContactHomePhoneArea1     string          `gorm:"type:varchar(4);column:emergency_contact_home_phone_area1"`
	EmergencyContactHomePhone1         string          `gorm:"type:varchar(50);column:emergency_contact_home_phone1"`
	EmergencyContactOfficePhoneArea1   string          `gorm:"type:varchar(4);column:emergency_contact_office_phone_area1"`
	EmergencyContactOfficePhone1       string          `gorm:"type:varchar(50);column:emergency_contact_office_phone1"`
	EmergencyContactMobilePhone        string          `gorm:"type:varchar(50);column:emergency_contact_mobile_phone"`
	CustomerIdConfins                  sql.NullString  `gorm:"type:varchar(20);column:customer_id_confins;index:idx_customers_varchar"`
	LegalName                          string          `gorm:"type:varchar(150);column:legal_name;index:idx_customers_varchar"`
	FullName                           string          `gorm:"type:varchar(150);column:full_name"`
	PersonalCustomerType               string          `gorm:"type:char(1);column:personal_customer_type"`
	IdType                             string          `gorm:"type:varchar(10);column:id_type"`
	IdNumber                           string          `gorm:"type:varchar(20);column:id_number;index:idx_customers_varchar"`
	ExpiredDate                        sql.NullTime    `gorm:"column:expired_date"`
	Gender                             string          `gorm:"type:char(1);column:gender"`
	BirthPlace                         string          `gorm:"type:varchar(100);column:birth_place"`
	BirthDate                          sql.NullTime    `gorm:"column:birth_date"`
	PersonalNPWPNumber                 string          `gorm:"type:varchar(25);column:personal_npwp_number"`
	SurgateMotherName                  string          `gorm:"type:varchar(100);column:surgate_mother_name"`
	Pin                                string          `gorm:"type:varchar(100);column:pin"`
	CreatedDatePin                     sql.NullTime    `gorm:"column:created_date_pin"`
	UpdatedDatePin                     sql.NullTime    `gorm:"column:updated_date_pin"`
	IsDataVerified                     bool            `gorm:"column:is_data_verified"`
	IsFaceVerified                     bool            `gorm:"column:is_face_verified"`
	IsLivenessVerified                 bool            `gorm:"column:is_liveness_verified"`
	IsIdForgeryVerified                bool            `gorm:"column:is_id_forgery_verified"`
	IsHumanVerified                    bool            `gorm:"column:is_human_verified"`
	LatestDateVerificationProcess      sql.NullTime    `gorm:"column:latest_date_verification_process"`
	ExpiredDateVerificationProcess     sql.NullTime    `gorm:"column:expired_date_verification_process"`
	ExpiredVerification                sql.NullTime    `gorm:"column:expired_verification"`
	IsNotActiveExpiredVerificationStep bool            `gorm:"column:is_not_active_expired_verification_step"`
	CustomerTransactionType            string          `gorm:"column:customer_transaction_type"`
	IsLockAccount                      bool            `gorm:"column:is_lock_account"`
	AttemptCSVerify                    int             `gorm:"column:attempt_cs_verify"`
	IncorrectAttemptPin                int             `gorm:"column:incorrect_attempt_pin"`
	IncorrectAttemptPinDatetime        sql.NullTime    `gorm:"column:incorrect_attempt_pin_datetime"`
	IncorrectAttemptChangePin          int             `gorm:"column:incorrect_attempt_change_pin"`
	IncorrectAttemptChangePinDatetime  sql.NullTime    `gorm:"column:incorrect_attempt_change_pin_datetime"`
	ForceForgetPin                     bool            `gorm:"column:force_forget_pin"`
	DatetimeIncorrectCSLatestAttempt   sql.NullTime    `gorm:"datetime_incorrect_cs_latest_attempt"`
	SourceId                           string          `gorm:"type:varchar(60);column:source_id"`
	AgreeToAcceptOtherOffering         sql.NullBool    `gorm:"column:agree_to_accept_other_offering"`
}

func (r Customer) TableName() string {
	return "customers"
}

type EncryptCustomer struct {
	ID                           int    `gorm:"primarykey;autoIncrement:true"`
	LegalName                    string `gorm:"type:varchar(150);column:legal_name;index:idx_customers_varchar"`
	FullName                     string `gorm:"type:varchar(150);column:full_name"`
	BirthPlace                   string `gorm:"type:varchar(100);column:birth_place"`
	SurgateMotherName            string `gorm:"type:varchar(100);column:surgate_mother_name"`
	MobilePhone                  string `gorm:"type:varchar(50);column:mobile_phone"`
	Email                        string `gorm:"type:varchar(100);column:email;index:idx_customers_varchar;index:idx_customers_email,unique"`
	ResidenceAddress             string `gorm:"type:varchar(200);column:residence_address"`
	ResidencePhone1              string `gorm:"type:varchar(50);column:residence_phone1"`
	LegalAddress                 string `gorm:"type:varchar(200);column:legal_address"`
	LegalPhone1                  string `gorm:"type:varchar(50);column:legal_phone1"`
	CompanyAddress               string `gorm:"type:varchar(200);column:company_address"`
	CompanyPhone1                string `gorm:"type:varchar(50);column:company_phone1"`
	EmergencyContactName         string `gorm:"type:varchar(50);column:emergency_contact_name"`
	EmergencyContactAddress      string `gorm:"type:varchar(200);column:emergency_contact_address"`
	EmergencyContactHomePhone1   string `gorm:"type:varchar(50);column:emergency_contact_home_phone1"`
	EmergencyContactMobilePhone  string `gorm:"type:varchar(50);column:emergency_contact_mobile_phone"`
	EmergencyContactOfficePhone1 string `gorm:"type:varchar(50);column:emergency_contact_office_phone1"`
}

func (r EncryptCustomer) TableName() string {
	return "customers"
}

type EncryptHistoryCustomer struct {
	ID                           int    `gorm:"primarykey;autoIncrement:true"`
	CustomerId                   int    `gorm:"column:customer_id"`
	LegalName                    string `gorm:"type:varchar(150);column:legal_name;index:idx_customers_varchar"`
	FullName                     string `gorm:"type:varchar(150);column:full_name"`
	BirthPlace                   string `gorm:"type:varchar(100);column:birth_place"`
	SurgateMotherName            string `gorm:"type:varchar(100);column:surgate_mother_name"`
	MobilePhone                  string `gorm:"type:varchar(50);column:mobile_phone"`
	Email                        string `gorm:"type:varchar(100);column:email;index:idx_customers_varchar;index:idx_customers_email,unique"`
	ResidenceAddress             string `gorm:"type:varchar(200);column:residence_address"`
	ResidencePhone1              string `gorm:"type:varchar(50);column:residence_phone1"`
	LegalAddress                 string `gorm:"type:varchar(200);column:legal_address"`
	LegalPhone1                  string `gorm:"type:varchar(50);column:legal_phone1"`
	CompanyAddress               string `gorm:"type:varchar(200);column:company_address"`
	CompanyPhone1                string `gorm:"type:varchar(50);column:company_phone1"`
	EmergencyContactName         string `gorm:"type:varchar(50);column:emergency_contact_name"`
	EmergencyContactAddress      string `gorm:"type:varchar(200);column:emergency_contact_address"`
	EmergencyContactHomePhone1   string `gorm:"type:varchar(50);column:emergency_contact_home_phone1"`
	EmergencyContactMobilePhone  string `gorm:"type:varchar(50);column:emergency_contact_mobile_phone"`
	EmergencyContactOfficePhone1 string `gorm:"type:varchar(50);column:emergency_contact_office_phone1"`
}

func (r EncryptHistoryCustomer) TableName() string {
	return "customers"
}

type SearchEncryptRepository interface {
	FetchByRange(ctx context.Context, model interface{}, search []int, minId int, maxId int) error
	GetCountAll(ctx context.Context) (int64, error)
	FindById(ctx context.Context, model interface{}, id int) error
}

type SearchHistoryEncryptRepository interface {
	FetchByRange(ctx context.Context, model interface{}, search []int, minId int, maxId int) error
	GetCountAll(ctx context.Context) (int64, error)
	FindById(ctx context.Context, model interface{}, id int) error
}

type SearchEncryptUseCase interface {
	CheckLength() (int64, error)
	FindAndSaveInvalidEncryptByRange(minId int, maxId int) (validAmount int, invalidAmount int, funcError error)
}

type SearchHistoryEncryptUseCase interface {
	CheckLength() (int64, error)
	FindAndSaveInvalidEncryptByRange(minId int, maxId int) (validAmount int, invalidAmount int, funcError error)
}

type SearchEncryptCmdHandler interface {
	CountAll() *cobra.Command
	FindInvalidEncryptByRange() *cobra.Command
	CountAllHistory() *cobra.Command
	FindInvalidHistoryEncryptByRange() *cobra.Command
}
