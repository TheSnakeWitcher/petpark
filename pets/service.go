package pets

import (
	"context"
	"database/sql"
	"math/rand"
)

type Service struct { *Queries }

func NewService(db *sql.DB) *Service {
    return &Service{ New(db) }
}

func (self Service) ListPets(ctx context.Context) ([]Pet , error) {
    pets,err := self.Queries.ListPets(ctx)
    if err != nil {
        return []Pet{} , err
    }

    return pets,nil
}

func (self Service) GetPet(ctx context.Context,id int32) (Pet,error) {
    pet , err := self.Queries.GetPet(ctx,id)
    if err != nil {
        return Pet{},err
    }
    return pet,nil
}   

func (self Service) AddPet(ctx context.Context,name,loc string) (sql.Result,error) {
    arg := AddPetParams{
        ID: rand.Int31(),
        Name: sql.NullString{String: name , Valid: true},
        Location: sql.NullString{String: loc , Valid: true},
    }
    pet, err := self.Queries.AddPet(ctx,arg)
    return pet,err
}

func (self Service) DelPet(ctx context.Context,id int32) (error) {
    err := self.Queries.DelPet(ctx,id)
    return err
}
