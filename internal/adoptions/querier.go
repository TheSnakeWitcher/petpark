// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package adoptions

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	AddPet(ctx context.Context, arg AddPetParams) (sql.Result, error)
	DelPet(ctx context.Context, id uuid.UUID) error
	GetPet(ctx context.Context, id uuid.UUID) (Pet, error)
	ListPets(ctx context.Context) ([]Pet, error)
}

var _ Querier = (*Queries)(nil)