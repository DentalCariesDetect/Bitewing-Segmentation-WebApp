package models

type UpdateModel struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	HospitalName string `json:"hospital_name" `
	Phone        string `json:"phone" validate:"omitempty,min=7,max=13"`
	Gender       string `json:"gender" validate:"omitempty,gender"`
	StartDate    string `json:"start_date"  `
	YearExp      uint64 `json:"year_ext" validate:"omitempty,min=0"`
}

type DentistModel struct {
	ID           uint64  `json:"dentist_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Username     string  `json:"username"`
	HospitalName *string `json:"hospital_name"`
	Phone        *string `json:"phone"`
	Gender       *string `json:"gender"`
	StartDate    *string `json:"start_date"`
	YearExp      *uint64 `json:"year_ext"`
}
