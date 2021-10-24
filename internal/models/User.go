package models

type User struct {
	Id         uint      `json:"user_id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	LocationId uint      `json:"location_id"`
	Location   *Location `json:"location" gorm:"foreignKey:LocationId"`
	Age        uint      `json:"age"`
}
