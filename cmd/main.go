package main

import (
	"notification_service/config"
	"notification_service/database"
	"notification_service/server"
)

func main() {
	// Load environment variables
	config.LoadEnv()
	database.LoadDB()

	_ = server.NewGRPCServer().Start()
}
