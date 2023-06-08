package main

import (
	seed_local_data "go-app/config/seed/local/data"
	"go-app/src/infrastructure/configs"
	"go-app/src/infrastructure/gorm/connection"
)

func main() {
	_, err := configs.InitConfig(".env")
	if err != nil {
		panic(err)
	}
	db := connection.DBConnect()

	// seed
	seed_local_data.SeedUsers(db)
}
