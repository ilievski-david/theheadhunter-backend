package main

import (
	"github.com/ilievski-david/theheadhunter-backend/initializers"
	"github.com/ilievski-david/theheadhunter-backend/models"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	_db, err := initializers.ConnectToDB()
	if err != nil {
		panic(err)
	}
	database = _db
}

func main() {
	database.Migrator().DropTable(&models.User{})
	database.Migrator().DropTable(&models.Color{})
	database.AutoMigrate(&models.User{}, &models.Color{})

}
