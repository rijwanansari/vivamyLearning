package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/dto"
	"github.com/rijwanansari/vivaLearning/services"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (a *AuthController) RegisterUser(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	user, err := a.userService.RegisterUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"user":    user.Email,
	})

}
