package testutil

import (
	"go-app/src/infrastructure/configs"
	"go-app/src/infrastructure/gorm/connection"

	"gorm.io/gorm"
)

func SetUp() *gorm.DB {
	_, err := configs.InitConfig(".env.test")
	if err != nil {
		panic(err)
	}

	db := connection.DBConnect()
	return db
}

func RefreshTables(db *gorm.DB, models []interface{}) error {
	for _, table := range models {
		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	for _, table := range models {
		if err := db.AutoMigrate(table); err != nil {
			return err
		}
	}

	return nil
}
