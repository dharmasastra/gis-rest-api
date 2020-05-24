package controller

import (
	"fmt"
	"github.com/dharmasastra/kost-gis-echo/app/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func GetAllKost(c echo.Context) (err error){
	var kost []models.Kost
	if err = db.Preload("FacilitiesRooms").Preload("BuildingFacilities").Preload("Locations").Find(&kost).Error; err != nil {
		fmt.Println(err)
		return
	}
	return c.JSON(http.StatusOK, kost)
}

func GetKost(c echo.Context) (err error) {
	name := c.Param("name")
	kost := models.Kost{}
	if err = db.Preload("FacilitiesRooms").Preload("BuildingFacilities").Preload("Locations").First(&kost, models.Kost{NameOwnerKost: name}).Error; err != nil {
		fmt.Println(err)
		return
	}
	return c.JSON(http.StatusOK, kost)
}

func CreateKost(c echo.Context) (err error) {
	var kost models.Kost
	if err = c.Bind(&kost); err != nil {
		return
	}
	if err = c.Validate(&kost); err != nil{
		return
	}
	db.Create(&kost)
	return c.JSON(http.StatusCreated, kost)
}

func UpdateKost(c echo.Context) (err error) {
	name := c.Param("name")
	kost := models.Kost{}
	if err = db.First(&kost, models.Kost{NameOwnerKost: name}).Error; err != nil {
		fmt.Println(err)
		return
	}
	if err = c.Bind(&kost); err != nil {
		return
	}
	if err = c.Validate(&kost); err != nil{
		return
	}
	if err = db.Save(&kost).Error; err != nil {
		return
	}
	return c.JSON(http.StatusOK, kost)
}

func DeleteKost(c echo.Context) (err error) {
	name := c.Param("name")
	kost := models.Kost{}
	if err = db.First(&kost, models.Kost{NameOwnerKost: name}).Error; err != nil {
		fmt.Println(err)
		return
	}
	if err = db.Delete(&kost).Error; err != nil {
		return
	}
	s := map[string]string{"message": name + " delete"}
	return c.JSON(http.StatusOK, s)
}