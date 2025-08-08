package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/dto"
	"github.com/rijwanansari/vivaLearning/services"
)

type AuthController struct {
	userService services.UserService
	authService services.AuthService
}

func NewAuthController(userService services.UserService, authService services.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

func (a *AuthController) RegisterUser(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	user, err := a.authService.Register(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"user":    user.Email,
	})

}

func (a *AuthController) LoginUser(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	token, err := a.authService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
