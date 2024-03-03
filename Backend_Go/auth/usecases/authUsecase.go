package usecases

import (
	"errors"
	authError "segmentation/auth/errors"
	"segmentation/auth/models"
	"segmentation/configs"
	"segmentation/dentists/entities"
	"segmentation/dentists/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	RegisterDataProcessing(in *models.RegisterData) error
	CheckData(in *models.RegisterData) error
	LoginDataProcession(in *models.LoginData) (*string, error)
}

type authUsecaseImpl struct {
	dentistRepository repositories.DentistRepository
}

func NewAuthUsecaseImpl(dentistRepository repositories.DentistRepository) AuthUsecase {
	return &authUsecaseImpl{
		dentistRepository: dentistRepository,
	}
}
func (u *authUsecaseImpl) CheckData(in *models.RegisterData) error {
	username := &in.Username
	//result = true -> found username
	//result = false -> Not found username
	//err != nill -> Found error
	if result, err := u.dentistRepository.Search("username", username); result || err != nil {
		return &authError.UsernameAlreadyExistError{}
	} else {
		if errors.Is(err, &authError.ServerInternalError{}) {
			return &authError.ServerInternalError{Err: err}
		}
	}
	return nil
}

func (u *authUsecaseImpl) RegisterDataProcessing(in *models.RegisterData) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &authError.ServerInternalError{Err: err}
	}
	insertDentistData := &entities.InsertDentist{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  string(hashedPassword),
		Status:    "Active",
	}

	if err := u.dentistRepository.InsertDentistData(insertDentistData); err != nil {
		return &authError.ServerInternalError{Err: err}
	}

	return nil
}

func (u *authUsecaseImpl) LoginDataProcession(in *models.LoginData) (*string, error) {
	username := &in.Username
	password := &in.Password
	//result = true -> found username
	//result = false -> Not found username
	//err != nill -> Found error
	if result, err := u.dentistRepository.Search("username", username); !result || err != nil {
		if err != nil {
			return nil, &authError.ServerInternalError{Err: err}
		}
		return nil, &authError.UsernameNotFoundError{}
	}
	if dentist, err := u.dentistRepository.GetDentistDataByKey("username", username); err != nil {
		//return error
		return nil, &authError.ServerInternalError{Err: err}
	} else {
		//compare password
		if err := bcrypt.CompareHashAndPassword([]byte(dentist.Password), []byte(*password)); err != nil {
			return nil, &authError.PasswordIncorrectError{}
		} else {
			//return success
			//generate token
			tokenUsecase := NewTokenUsecaseImpl(configs.GetJwtConfig().SecretKey)
			token, err := tokenUsecase.GenerateToken(&dentist.ID, &dentist.Username)
			if err != nil {
				return nil, &authError.ServerInternalError{Err: err}
			}
			return token, nil
		}
	}
}
