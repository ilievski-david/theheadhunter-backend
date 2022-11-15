package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilievski-david/theheadhunter-backend/crud"
	"github.com/ilievski-david/theheadhunter-backend/models"
)

type Handler interface {
	AddColor(context *gin.Context)
	GetColors(context *gin.Context)
	RemoveColor(context *gin.Context)
}

type handler struct {
	c_db crud.Database
}

func NewHandler(_crud crud.Database) Handler {
	return &handler{
		c_db: _crud,
	}
}

func (h *handler) GetColors(context *gin.Context) {
	userToken := context.Param("token")
	var colors []models.Color
	colors, err := h.c_db.QueryColors(userToken)
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
	context.JSON(code, gin.H{"code": code})

}

func (h *handler) RemoveColor(context *gin.Context) {
	var newColorRemove models.ColorRemove
	err := context.BindJSON(&newColorRemove)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.c_db.DeleteColor(newColorRemove)
	if err != nil {
		// implement no color found
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": 200})
}

func (h *handler) userHandler(userToken string) error {
	err := h.c_db.IgnoreOrInsertUser(userToken)
	return err
}

func (h *handler) colorHandler(newColorPost models.ColorPost) int {
	if newColorPost.Name == "" {
		return 402
	} else if len(newColorPost.Name) > 20 {
		//return 403
	}

	colors, err := h.c_db.QueryColors(newColorPost.UserToken)
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
	h.c_db.InsertColor(newColorPost)
	return 200
}
