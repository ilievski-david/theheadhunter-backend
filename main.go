package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilievski-david/theheadhunter-backend/crud"
	"github.com/ilievski-david/theheadhunter-backend/handlers"
	"github.com/ilievski-david/theheadhunter-backend/initializers"
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
	c_db := crud.NewCrudDatabase(database)
	h := handlers.NewHandler(c_db)
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/getColors/:token", h.GetColors)
	router.POST("/addColor", h.AddColor)
	router.DELETE("/removeColor", h.RemoveColor)
	enviorment := os.Getenv("ENVIORMENT")
	if enviorment == "production" {
		ssl_folder := os.Getenv("SERVER_SSL")
		router.RunTLS(":8080", ssl_folder+"/server.crt", ssl_folder+"/server.key")
	} else {
		router.Run(":8080")
	}

}
