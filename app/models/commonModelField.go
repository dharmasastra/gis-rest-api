package models

import "time"

type CommonModelFields struct {
	Id			uint		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time  	`json:"createdAt"`
	UpdatedAt 	time.Time  	`json:"updatedAt"`
	DeleteAt	*time.Time	`json:"deleteAt"`
} 