package controllers

import (
	"net/http"

	_ "github.com/krhone/go-quest/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func New() http.Handler {
	router := echo.New()

	config := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(config))

	router.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/quests", GetAllQuests)
		v1.GET("/quest/:id", GetQuest)
		v1.POST("/quest", CreateQuest)
		v1.PUT("/quest/:id", UpdateQuest)
		v1.DELETE("/quest/:id", DeleteQuest)
	}
	v1.GET("/", HealthCheck)
	return router
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
