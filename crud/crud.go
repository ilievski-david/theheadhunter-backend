package crud

import (
	"github.com/ilievski-david/theheadhunter-backend/models"
	"gorm.io/gorm"
)

type Database interface {
	InsertColor(newColor models.ColorPost) error
	IgnoreOrInsertUser(token string) error
	DeleteColor(newColorRemove models.ColorRemove) error
	QueryColors(userToken string) ([]models.Color, error)
}

type database struct {
	DB *gorm.DB
}

func NewCrudDatabase(_db *gorm.DB) Database {
	return &database{
		DB: _db,
	}
}

func (d *database) InsertColor(newColorPost models.ColorPost) error {
	err := d.DB.Create(&models.Color{UserToken: newColorPost.UserToken, Hex: newColorPost.Hex, Name: newColorPost.Name}).Error
	return err
}

func (d *database) IgnoreOrInsertUser(token string) error {
	err := d.DB.FirstOrCreate(&models.User{}, models.User{UserToken: token}).Error
	return err
}

func (d *database) DeleteColor(newColorRemove models.ColorRemove) error {
	err := d.DB.Exec("DELETE FROM colors WHERE id= ? AND user_token= ?", newColorRemove.ID, newColorRemove.UserToken).Error
	return err
}

func (d *database) QueryColors(userToken string) ([]models.Color, error) {
	var colors []models.Color
	err := d.DB.Raw("SELECT * FROM colors WHERE user_token = ?", userToken).Scan(&colors).Error
	return colors, err
}
