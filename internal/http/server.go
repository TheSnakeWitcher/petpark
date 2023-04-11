package http

import (
	"github.com/TheSnakeWitcher/PetPark/pets"
	"github.com/labstack/echo/v4"
)

type Server struct {
    *echo.Echo
}

func NewServer(svc pets.Service) (srv *Server) {
    e := echo.New()
    srv = &Server { e }
    srv.InitRoutes(svc)
    return srv
}
