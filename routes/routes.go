package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/controllers"
)

type Routes struct {
	echo *echo.Echo
	auth *controllers.AuthController
}

func New(e *echo.Echo, auth *controllers.AuthController) *Routes {
	return &Routes{
		echo: e,
		auth: auth,
	}
}

func (r *Routes) Init() {
	e := r.echo

	// Define your routes here\
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	g := e.Group("/api/v1")
	// auth routes
	g.POST("/register", r.auth.RegisterUser)
}
