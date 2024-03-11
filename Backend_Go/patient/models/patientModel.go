package models

type (
	PatientModel struct {
		ID        uint64  `json:"patient_id"`
		BirthDate string  `json:"birth_date"`
		Gender    string  `json:"gender"`
		Age       uint64  `json:"age"`
		Phone     *string `json:"phone"`
	}

	PatientInsertModel struct {
		BirthDate string `json:"birth_date" validate:"required"`
		Gender    string `json:"gender" validate:"required,gender"`
		Age       uint64 `json:"age" validate:"required,min=0"`
		Phone     string `json:"phone" validate:"omitempty,min=7,max=13"`
	}

	PatientUpdateModel struct {
		BirthDate string `json:"birth_date" validate:"required"`
		Gender    string `json:"gender" validate:"required,gender"`
		Age       uint64 `json:"age" validate:"required,min=0"`
		Phone     string `json:"phone" validate:"omitempty,min=7,max=13"`
	}
)
