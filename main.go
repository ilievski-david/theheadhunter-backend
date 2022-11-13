package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilievski-david/theheadhunter-backend/initializers"
	"github.com/ilievski-david/theheadhunter-backend/models"
)

type colorPost struct {
	UserToken string `json:"userToken"`
	Hex       string `json:"hex"`
	Name      string `json:"name"`
}

type colorGet struct {
	UserToken string `json:"userToken"`
}

type colorRemove struct {
	UserToken string `json:"userToken"`
	ID        int    `json:"id"`
}

func getColors(context *gin.Context) {
	var newColorGet colorGet
	err := context.BindJSON(&newColorGet)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		return
	}
	//newColorGet.UserToken
	var colors []models.Color
	err = initializers.DB.Raw("SELECT * FROM colors WHERE user_token = ?", newColorGet.UserToken).Scan(&colors).Error
	if err != nil {
		context.JSON(402, gin.H{"error": err.Error()})
		return
	}
	// print in line colors
	fmt.Println(colors)
	context.JSON(http.StatusOK, colors)

}

func addColor(context *gin.Context) {
	var newColorPost colorPost
	err := context.BindJSON(&newColorPost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userHandler(newColorPost.UserToken)
	code := colorHandler(newColorPost)
	// how to send status codes?
	context.JSON(code, gin.H{"code": code})

}

func removeColor(context *gin.Context) {
	var newColorRemove colorRemove
	err := context.BindJSON(&newColorRemove)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Exec(fmt.Sprintf("DELETE FROM colors WHERE id=%d AND user_token='%s'", newColorRemove.ID, newColorRemove.UserToken))
	if result.RowsAffected == 0 {
		// implement no color found
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": 200})
}

func colorHandler(newColorPost colorPost) int {
	if newColorPost.Name == "" {
		return http.StatusBadRequest
	} else if len(newColorPost.Name) > 20 {
		return 403
	}
	var colors []models.Color
	err := initializers.DB.Raw("SELECT * FROM colors WHERE user_token = ?", newColorPost.UserToken).Scan(&colors).Error
	if err != nil {
		return 404
	}

	for _, color := range colors {
		if color.Name == newColorPost.Name {
			return 405
		}
		if color.Hex == newColorPost.Hex {
			return 406
		}
	}
	initializers.DB.Create(&models.Color{UserToken: newColorPost.UserToken, Hex: newColorPost.Hex, Name: newColorPost.Name})
	return 200
}

func userHandler(userToken string) {
	result := initializers.DB.FirstOrCreate(&models.User{}, models.User{UserToken: userToken})
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/getColors", getColors)
	router.POST("/addColor", addColor)
	router.DELETE("/removeColor", removeColor)
	enviorment := os.Getenv("ENVIORMENT")
	if enviorment == "production" {
		ssl_folder := os.Getenv("SERVER_SSL")
		router.RunTLS(":8080", ssl_folder+"/server.crt", ssl_folder+"/server.key")
	} else {
		router.Run(":8080")
	}

}
