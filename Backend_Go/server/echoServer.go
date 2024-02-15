package server

import (
	"fmt"
	"segmentation/configs"

	authHandler "segmentation/auth/handlers"
	authUsecase "segmentation/auth/usecases"
	dentistRepository "segmentation/dentists/repositories"

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
	s.app.Validator = &CustomValidator{validator: validator.New()}
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
	authRouters.POST("/login", authHttpHandler.Register)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
