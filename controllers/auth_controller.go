package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/dto"
	"github.com/rijwanansari/vivaLearning/services"
	"github.com/rijwanansari/vivaLearning/utils"
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": "Invalid input",
		})
	}

	user, err := a.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Generate token for the new user
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"message": "Failed to generate token",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true,
		"message": "User registered successfully",
		"data": echo.Map{
			"user": echo.Map{
				"id":         user.ID,
				"name":       user.Name,
				"email":      user.Email,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			},
			"access_token":  token,
			"refresh_token": token,
			"expires_at":    "7d",
		},
	})
}

func (a *AuthController) LoginUser(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": "Invalid input",
		})
	}

	user, token, err := a.authService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Login successful",
		"data": echo.Map{
			"user": echo.Map{
				"id":         user.ID,
				"name":       user.Name,
				"email":      user.Email,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			},
			"access_token":  token,
			"refresh_token": token, // For now, using same token for both
			"expires_at":    "7d",  // Simplified expiry info
		},
	})
}
