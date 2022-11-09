package model

import "time"

type Question struct {
	Id          int `gorm:"primaryKey"`
	Title       string
	PersonId    int `gorm:"foreignKey"`
	CreatedDate time.Time
	Image       string
	Person      Person
}
