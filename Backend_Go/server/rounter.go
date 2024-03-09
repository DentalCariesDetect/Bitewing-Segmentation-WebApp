package server

import (
	authHandler "segmentation/auth/handlers"
	authUsecase "segmentation/auth/usecases"

	dentistHandler "segmentation/dentists/handlers"
	dentistRepository "segmentation/dentists/repositories"
	dentistUsecase "segmentation/dentists/usecases"

	patientHandler "segmentation/patient/handlers"
	patientRepository "segmentation/patient/repositories"
	patientUsecase "segmentation/patient/usecases"
)

func (s *echoServer) initializeRouters() {
	s.initializeAuthHttpHandler()
	s.initializeDentistHttpHandler()
	s.initializePatientHttpHandler()
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
	dentistRouters.PUT("/:id", dentistHttpHandler.UpdateDentist)
	dentistRouters.GET("/:id", dentistHttpHandler.GetDentistById)
	dentistRouters.GET("/", dentistHttpHandler.GetDentistAll)
	dentistRouters.DELETE("/:id", dentistHttpHandler.DeleteDentist)
}

func (s *echoServer) initializePatientHttpHandler() {
	patientPosgresRepository := patientRepository.NewPatientPostgresRepository(s.db)
	patientUsecase := patientUsecase.NewPatientUsecaseImpl(patientPosgresRepository)
	patientHttpHandler := patientHandler.NewPatientHttpHandler(patientUsecase)

	patientRouters := s.app.Group("v1/patient")

	patientRouters.Use(TokenAuthentication(dentistRepositoryForAuth(s)))
	patientRouters.POST("/", patientHttpHandler.CreatePatient)
	patientRouters.GET("/:id", patientHttpHandler.GetPatientById)
	patientRouters.GET("/", patientHttpHandler.GetPatientsByDentist)
	patientRouters.DELETE("/:id", patientHttpHandler.DeletePatient)
	patientRouters.PUT("/:id", patientHttpHandler.UpdatePatient)
}

func dentistRepositoryForAuth(s *echoServer) dentistRepository.DentistRepository {
	return dentistRepository.NewDentistPostgresRepository(s.db)
}
