package connection

import (
	"fmt"
	"go-app/src/infrastructure/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", configs.Config.DB.User, configs.Config.DB.Password, configs.Config.DB.Host, configs.Config.DB.Port, configs.Config.DB.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
