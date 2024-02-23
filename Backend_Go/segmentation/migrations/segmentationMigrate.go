package main

import (
	"segmentation/configs"
	"segmentation/database"
	"segmentation/segmentation/entities"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	segmentationMigrate(db)
}

func segmentationMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Image{})
	db.GetDb().Migrator().CreateTable(&entities.Bitewing{})
	db.GetDb().Migrator().CreateTable(&entities.Result{})
	db.GetDb().Migrator().CreateTable(&entities.Tooth{})

}
