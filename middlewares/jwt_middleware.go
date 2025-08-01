package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/utils"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing token"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := utils.ParseJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		c.Set("user_id", userID)
		return next(c)
	}
}
