// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package adoptions

import (
	"time"

	"github.com/google/uuid"
)

type Pet struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Picked  time.Time `db:"picked" json:"picked"`
	Address string    `db:"address" json:"address"`
	Contact string    `db:"contact" json:"contact"`
	Details string    `db:"details" json:"details"`
}
