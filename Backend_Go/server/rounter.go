package server

import (
	authHandler "segmentation/auth/handlers"
	authUsecase "segmentation/auth/usecases"
	dentistHandler "segmentation/dentists/handlers"
	dentistRepository "segmentation/dentists/repositories"
	dentistUsecase "segmentation/dentists/usecases"
)

func (s *echoServer) initializeRouters() {
	s.initializeAuthHttpHandler()
	s.initializeDentistHttpHandler()
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
	dentistRouters.GET("/:id", dentistHttpHandler.GetDentistById)
	dentistRouters.GET("/", dentistHttpHandler.GetDentistAll)
	dentistRouters.DELETE("/:id", dentistHttpHandler.DeleteDentist)
}
