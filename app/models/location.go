package models

type Location struct {
	IdLocation	uint	`gorm:"primary_key" json:"idLocation"`
	Latitude	float64	`json:"latitude" validate:"required"`
	Longitude	float64	`json:"longitude" validate:"required"`
}
