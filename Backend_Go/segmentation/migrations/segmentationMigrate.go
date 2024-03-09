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
	db.GetDb().AutoMigrate(&entities.Image{}, &entities.Bitewing{}, &entities.Result{}, &entities.Tooth{})

}
