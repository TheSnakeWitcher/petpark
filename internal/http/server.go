package http

import (
	"github.com/TheSnakeWitcher/petplanet/pets"
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

func (self *Server) InitRoutes(svc pets.Service) {
    self.GET("/hello",Home)
    self.GET("/pets/",ListPets(svc))
    self.GET("/pets/:id",GetPet(svc))
    self.POST("/pets",AddPet(svc))
    self.DELETE("/pets/:id",DelPet(svc))
}
