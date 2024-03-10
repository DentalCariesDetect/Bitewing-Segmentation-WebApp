package main

import (
	"segmentation/configs"
	"segmentation/database"
	"segmentation/dentists/entities"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	dentistMigrate(db)
}

func dentistMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Dentist{})
}
