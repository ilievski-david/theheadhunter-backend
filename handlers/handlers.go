package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilievski-david/theheadhunter-backend/models"
	"gorm.io/gorm"
)

type Handler interface {
	AddColor(context *gin.Context)
	GetColors(context *gin.Context)
	RemoveColor(context *gin.Context)
}

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return &handler{
		DB: db,
	}
}

func (h *handler) GetColors(context *gin.Context) {
	userToken := context.Param("token")
	var colors []models.Color
	err := h.DB.Raw("SELECT * FROM colors WHERE user_token = ?", userToken).Scan(&colors).Error
	if err != nil {
		context.JSON(402, gin.H{"error": err.Error()})
		return
	}
	// print in line colors
	fmt.Println(colors)
	context.JSON(http.StatusOK, colors)

}

func (h *handler) AddColor(context *gin.Context) {
	var newColorPost models.ColorPost
	err := context.BindJSON(&newColorPost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.userHandler(newColorPost.UserToken)
	code := h.colorHandler(newColorPost)
	// how to send status codes?
	context.JSON(code, gin.H{"code": code})

}

func (h *handler) RemoveColor(context *gin.Context) {
	var newColorRemove models.ColorRemove
	err := context.BindJSON(&newColorRemove)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Exec("DELETE FROM colors WHERE id= ? AND user_token= ?", newColorRemove.ID, newColorRemove.UserToken)
	if result.RowsAffected == 0 {
		// implement no color found
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": 200})
}

func (h *handler) userHandler(userToken string) {
	result := h.DB.FirstOrCreate(&models.User{}, models.User{UserToken: userToken})
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func (h *handler) colorHandler(newColorPost models.ColorPost) int {
	if newColorPost.Name == "" {
		return 402
	} else if len(newColorPost.Name) > 20 {
		return 403
	}
	var colors []models.Color
	err := h.DB.Raw("SELECT * FROM colors WHERE user_token = ?", newColorPost.UserToken).Scan(&colors).Error
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
	h.DB.Create(&models.Color{UserToken: newColorPost.UserToken, Hex: newColorPost.Hex, Name: newColorPost.Name})
	return 200
}
