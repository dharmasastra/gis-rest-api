package config

import (
	controller "github.com/dharmasastra/gis-rest-api/app/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func NewRouter() *echo.Echo{
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} | ${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/kosts", controller.GetAllKost)
	e.GET("/kost/:name", controller.GetKost)
	e.POST("/kost", controller.CreateKost)
	e.POST("/kost/:name", controller.UpdateKost)
	e.DELETE("/kost/:name", controller.DeleteKost)

	return e
}