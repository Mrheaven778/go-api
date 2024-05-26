package db

import "github.com/mrheaven778/go-astro-crud/models"

func autoMigrateDB() {
	err := DB.AutoMigrate(models.Task{})
	DB.AutoMigrate(models.User{})

	if err != nil {
		return
	}
}
