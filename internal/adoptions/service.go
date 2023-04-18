package adoptions

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type PetService = Querier

type Service struct { *Queries }

func NewService(db *sql.DB) *Service {
    return &Service{ New(db) }
}

func (svc Service) ListPets(ctx context.Context) ([]Pet , error) {
    pets,err := svc.Queries.ListPets(ctx)
    if err != nil {
        return []Pet{} , err
    }
    return pets,nil
}

func (svc Service) GetPet(ctx context.Context,id uuid.UUID) (Pet,error) {
    pet , err := svc.Queries.GetPet(ctx,id)
    if err != nil {
        return Pet{},err
    }
    return pet,nil
}   

func (svc Service) AddPet(ctx context.Context,arg AddPetParams) (sql.Result,error) {
    outPet, err := svc.Queries.AddPet(ctx,arg)
    return outPet,err
}

func (svc Service) DelPet(ctx context.Context,id uuid.UUID) (error) {
    err := svc.Queries.DelPet(ctx,id)
    return err
}
