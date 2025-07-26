package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func New(echo *echo.Echo) *Server {
	return &Server{
		echo: echo,
	}
}

func (s *Server) Start(port int) {
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%d", port)))
}
