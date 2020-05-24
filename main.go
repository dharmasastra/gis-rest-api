package main

import (
	"github.com/dharmasastra/gis-rest-api/app/controllers"
	"github.com/dharmasastra/gis-rest-api/config"
	"github.com/go-playground/validator/v10"
)

func main() {
	e := config.NewRouter()
	e.Validator = &controller.CustomValidator{Validator: validator.New()}


	e.Logger.Fatal(e.Start(":8080"))
}