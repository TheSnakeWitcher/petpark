package adoptions

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type (
    ListPetsRes struct {
        Pets []Pet  `json:"pets"`
        Err  error  `json:"error"`
    }

    GetPetReq struct { Id string `json:"id"`} 
    GetPetRes struct {
        Pet Pet   `json:"pet"`
        Err error `json:"error"`
    }

    AddPetReq struct { Arg AddPetParams `"json:"pet"` }
    AddPetRes struct { 
        Result sql.Result `json:"result"`
        Err    error      `json:"error"`
    }

    DelPetReq struct { Id string `json:"id"`}
    DelPetRes struct { Err error `json:"error"`}
)

type Endpoints struct {
	AddPetEndpoint   endpoint.Endpoint
	DelPetEndpoint   endpoint.Endpoint
	GetPetEndpoint   endpoint.Endpoint
	ListPetsEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(svc PetService) Endpoints {
    return Endpoints {
        ListPetsEndpoint:  MakeListPetsEndpoint(svc),
        GetPetEndpoint:  MakeGetPetEndpoint(svc),
        AddPetEndpoint:  MakeAddPetEndpoint(svc),
        DelPetEndpoint:  MakeDelPetEndpoint(svc),
    }
}

func MakeListPetsEndpoint(svc PetService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		petList,err := svc.ListPets(ctx)
		if err != nil {
			return ListPetsRes{[]Pet{} , err }, nil
		}
		return ListPetsRes{petList , nil}, nil
	}
}

func MakeGetPetEndpoint(svc PetService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(GetPetReq)
		pet, err := svc.GetPet(ctx,uuid.MustParse(req.Id))
		if err != nil {
			return GetPetRes{Pet{}, err}, nil
		}
		return GetPetRes{pet, nil}, nil
	}
}

func MakeAddPetEndpoint(svc PetService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(AddPetReq)
		result,err := svc.AddPet(ctx,req.Arg)
		return AddPetRes{result , err}, nil
	}
}

func MakeDelPetEndpoint(svc PetService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(DelPetReq)
		err := svc.DelPet(ctx,uuid.MustParse(req.Id))
		return DelPetRes{err}, nil
	}
}
