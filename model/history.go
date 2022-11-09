package model

import "time"

type History struct {
	Id        int `gorm:"primaryKey"`
	PersonId  int `gorm:"foreignKey"`
	Time      time.Time
	IpAddress string
	Person    Person
}
