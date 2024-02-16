package usecases

import (
	"errors"
	auth_error "segmentation/auth/entities"
	"segmentation/auth/models"
	"segmentation/configs"
	"segmentation/dentists/entities"
	"segmentation/dentists/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		return &auth_error.UsernameAlreadyExistError{}
	} else {
		if errors.Is(err, &auth_error.ServerInternalError{}) {
			return &auth_error.ServerInternalError{Err: err}
		}
	}
	return nil
}

func (u *authUsecaseImpl) RegisterDataProcessing(in *models.RegisterData) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &auth_error.ServerInternalError{Err: err}
	}
	insertDentistData := &entities.InsertDentist{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  string(hashedPassword),
		CreateOn:  time.Now(),
		Status:    "Active",
	}

	if err := u.dentistRepository.InsertDentistData(insertDentistData); err != nil {
		return &auth_error.ServerInternalError{Err: err}
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
		return nil, &auth_error.UsernameNotFoundError{}
	} else {
		if errors.Is(err, &auth_error.ServerInternalError{}) {
			return nil, &auth_error.ServerInternalError{Err: err}
		}
	}
	if dentist, err := u.dentistRepository.GetDentistDataByKey("username", username); err != nil {
		//return error
		return nil, &auth_error.ServerInternalError{Err: err}
	} else {
		//compare password
		if err := bcrypt.CompareHashAndPassword([]byte(dentist.Password), []byte(*password)); err != nil {
			return nil, &auth_error.PasswordIncorrectError{}
		} else {
			//return success
			//generate token
			token, err := generateToken(&dentist.DentistId, &dentist.Username)
			if err != nil {
				return nil, &auth_error.ServerInternalError{Err: err}
			}
			return token, nil
		}
	}

}

func generateToken(id *uint32, username *string) (*string, error) {

	// Load the jwt config
	config := configs.GetJwtConfig()
	key := []byte(config.SecretKey)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	clams := token.Claims.(jwt.MapClaims)
	clams["authorized"] = true
	clams["dentist_id"] = *id
	clams["username"] = *username
	clams["exp"] = time.Now().Add(time.Hour * 2).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
