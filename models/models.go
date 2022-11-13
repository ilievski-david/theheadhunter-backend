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

type ColorPost struct {
	UserToken string `json:"userToken"`
	Hex       string `json:"hex"`
	Name      string `json:"name"`
}

type ColorRemove struct {
	UserToken string `json:"userToken"`
	ID        int    `json:"id"`
}
