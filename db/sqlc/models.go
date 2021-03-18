package db

import (
	//"math/big"
	"time"
	//"github.com/google/uuid"
)

type Account struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID                   int64     `json:"id"`
	AccountOriginID      int64     `json:"account_origin_id"`
	AccountDestinationID int64     `json:"account_destination_id"`
	Amount               int64     `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
