package main

import (
	"segmentation/configs"
	"segmentation/database"
	"segmentation/patient/entities"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	patientMigrate(db)
}

func patientMigrate(db database.Database) {
	db.GetDb().AutoMigrate(&entities.Patient{})
}
