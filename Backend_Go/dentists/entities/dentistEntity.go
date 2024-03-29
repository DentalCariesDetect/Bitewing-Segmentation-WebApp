package entities

import (
	"gorm.io/gorm"
)

type (
	Dentist struct {
		gorm.Model
		FirstName    string  `json:"first_name"`
		LastName     string  `json:"last_name"`
		Username     string  `json:"username"`
		Password     string  `json:"password"`
		Gender       *string `json:"gender"`
		StartDate    *string `json:"start_date"`
		YearExp      *uint64 `json:"year_ext"`
		HospitalName *string `json:"hospital_name"`
		Phone        *string `json:"phone"`
		Status       string  `json:"status"`
	}

	UpdateDentist struct {
		gorm.Model
		FirstName    string  `json:"first_name"`
		LastName     string  `json:"last_name"`
		Gender       *string `json:"gender"`
		StartDate    *string `json:"start_date"`
		YearExp      *uint64 `json:"year_ext"`
		HospitalName *string `json:"hospital_name"`
		Phone        *string `json:"phone"`
		Status       string  `json:"status"`
	}
	InsertDentist struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Status    string `json:"status"`
	}

	// DentistResponse struct {
	// 	DentistId    uint32    `json:"dentist_id"`
	// 	FirstName    string    `json:"first_name"`
	// 	LastName     string    `json:"last_name"`
	// 	Gender       string    `json:"gender"`
	// 	StartDate    time.Time `json:"start_date"`
	// 	YearExp      uint32    `json:"year_ext"`
	// 	HospitalName string    `json:"hospital_name"`
	// 	Phone        string    `json:"phone"`
	// 	CreateOn     time.Time `json:"create_on"`
	// 	UpdateOn     time.Time `json:"update_on"`
	// }
)
