package server

import (
	"fmt"
	"segmentation/configs"

	authHandler "segmentation/auth/handlers"
	authUsecase "segmentation/auth/usecases"
	dentistHandler "segmentation/dentists/handlers"
	dentistRepository "segmentation/dentists/repositories"
	dentistUsecase "segmentation/dentists/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *configs.Config
}

func NewEchoServer(cfg *configs.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	// initialize routers here
	s.initializeAuthHttpHandler()
	s.initializeDentistHttpHandler()
	s.app.Validator = NewCustomValidator()
	s.app.Use(middleware.Logger())
	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeAuthHttpHandler() {
	dentistPosgresRepository := dentistRepository.NewDentistPostgresRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecaseImpl(dentistPosgresRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(authUsecase)

	authRouters := s.app.Group("v1/auth")
	authRouters.POST("/register", authHttpHandler.Register)
	authRouters.POST("/login", authHttpHandler.Login)
}

func (s *echoServer) initializeDentistHttpHandler() {
	dentistPosgresRepository := dentistRepository.NewDentistPostgresRepository(s.db)
	dentistUsecase := dentistUsecase.NewDentistUsecaseImpl(dentistPosgresRepository)
	dentistHttpHandler := dentistHandler.NewDentistHttpHandler(dentistUsecase)

	dentistRouters := s.app.Group("v1/dentist")
	dentistRouters.Use(TokenAuthentication(dentistPosgresRepository))
	dentistRouters.POST("/update/:id", dentistHttpHandler.UpdateDentist)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func genderValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == "female" || value == "male"
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("gender", genderValidation)
	return &CustomValidator{validator: v}
}
