package controllers

import (
	"log"
	"net/http"

	"go_echo_api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) CreateUserHandler(c echo.Context) error {
	req := CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	user := models.NewUser(req.Name)
	if err := user.Create(h.DB); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "failed to create user")
	}

	return c.JSON(http.StatusCreated, "user created")
}
