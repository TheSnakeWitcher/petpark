package http

import "github.com/TheSnakeWitcher/PetPark/pets"

func (srv *Server) InitRoutes(svc pets.Service) {
    srv.GET("/",Root)
    srv.GET("/pets",ListPets(svc))
    srv.GET("/pets/:id",GetPet(svc))
    srv.POST("/pets",AddPet(svc))
    srv.DELETE("/pets/:id",DelPet(svc))
}
