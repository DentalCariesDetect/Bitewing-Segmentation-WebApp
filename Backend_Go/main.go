package main

import (
	"segmentation/configs"
	"segmentation/database"
	"segmentation/server"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
