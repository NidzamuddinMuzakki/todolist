package app

import (
	"github.com/NidzamuddinMuzakki/todolist/controller"
	"github.com/NidzamuddinMuzakki/todolist/middlewares"
	"github.com/ReneKroon/ttlcache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(ActivityController controller.ActivityController, TodolistController controller.TodolistController, cache *ttlcache.Cache) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())

	// r.Use(middlewares.Auth)
	r.Use(middlewares.Recover)
	Activity := r.Group("activity-groups")
	{
		Activity.GET("", func(c echo.Context) error {
			ActivityController.FindAll(c, cache)
			return nil
		})
		Activity.GET("/:id", func(c echo.Context) error {
			ActivityController.FindById(c, cache)
			return nil
		})
		Activity.POST("", func(c echo.Context) error {
			ActivityController.Insert(c, cache)
			return nil
		})
		Activity.PATCH("/:id", func(c echo.Context) error {
			ActivityController.Update(c, cache)
			return nil
		})
		Activity.DELETE("/:id", func(c echo.Context) error {
			ActivityController.Delete(c, cache)
			return nil
		})

	}

	Todolsit := r.Group("todo-items")
	{
		Todolsit.GET("", func(c echo.Context) error {
			TodolistController.FindAll(c, cache)
			return nil
		})
		Todolsit.GET("/:id", func(c echo.Context) error {
			TodolistController.FindById(c, cache)
			return nil
		})
		Todolsit.POST("", func(c echo.Context) error {
			TodolistController.Insert(c, cache)
			return nil
		})
		Todolsit.PATCH("/:id", func(c echo.Context) error {
			TodolistController.Update(c, cache)
			return nil
		})
		Todolsit.DELETE("/:id", func(c echo.Context) error {
			TodolistController.Delete(c, cache)
			return nil
		})

	}
	return r
}
