package models


type FacilityRoom struct {
	IdFacilityRoom	uint 	`gorm:"primary_key" json:"idFacilityRoom"`
	Facility		string	`json:"facilities" validate:"required"`
}

type BuildingFacility struct {
	IdFacilityBuilding	uint	`gorm:"primary_key" json:"idFacilityBuilding"`
	Facility			string	`json:"facilities" validate:"required"`
}