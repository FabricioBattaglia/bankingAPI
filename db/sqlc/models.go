package db

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   big.Float `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	AccountOriginID      uuid.UUID `json:"account_origin_id"`
	AccountDestinationID uuid.UUID `json:"account_destination_id"`
	Amount               big.Float `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}
