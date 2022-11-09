package model

type Person struct {
	Id           int `gorm:"primaryKey"`
	Nickname     string
	Email        string
	Password     string
	ProfileImage string
}
