package routes

import (
	"go_echo_api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(h *controllers.Handler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", controllers.HelloHandler)

	users := e.Group("/users")
	users.POST("/", h.CreateUserHandler)
	users.GET("/", h.GetUsersHandler)
	users.PUT("/:id", h.UpdateUserHandler)
	users.DELETE("/:id", h.DeleteUserHandler)

	return e
}
