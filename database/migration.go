package database

import (
	"fmt"
	"indochat/models"
	"indochat/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Customers{},
		&models.Categories{},
		&models.Products{},
		&models.Orders{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
