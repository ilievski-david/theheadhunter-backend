package models

type User struct {
	ID        uint   `gorm:"primaryKey", autoIncrement:"true", json:"id"`
	UserToken string `json:"userToken"`
}

type Color struct {
	ID        uint   `gorm:"primaryKey, autoIncrement:"true", json:"id"`
	UserToken string `json:"userToken"`
	Hex       string `json:"hex"`
	Name      string `json:"name"`
}
