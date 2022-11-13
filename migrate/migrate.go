package main

import (
	"github.com/ilievski-david/theheadhunter-backend/initializers"
	"github.com/ilievski-david/theheadhunter-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.Migrator().DropTable(&models.User{})
	initializers.DB.Migrator().DropTable(&models.Color{})
	initializers.DB.AutoMigrate(&models.User{}, &models.Color{})

}
