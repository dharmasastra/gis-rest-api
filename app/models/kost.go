package models

type Kost struct {

	CommonModelFields
	NameOwnerKost		string				`gorm:"unique; not null" json:"nameOwnerKost" validate:"required"`
	AddressOfOwner		string				`gorm:"unique" json:"addressOfOwner" validate:"required"`
	AdressOfKost		string				`gorm:"unique" json:"addressOfKost" validate:"required"`
	FacilitiesRooms		[]FacilityRoom		`gorm:"ForeignKey:IdFacilityRoom" json:"facilitiesRoom"`
	BuildingFacilities	[]BuildingFacility	`gorm:"ForeignKey:IdFacilityBuilding" json:"buildingFacilities"`
	Locations			[]Location			`gorm:"ForeignKey:IdLocation" json:"locations"`
}

