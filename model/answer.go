package model

import "time"

type Answer struct {
	Id         int `gorm:"primaryKey"`
	QuestionId int `gorm:"foreignKey"`
	Content    string
	UpdateDate time.Time
	Question   Question
}
