package db

import (
	"context"
	//"math/big"
	//"github.com/google/uuid"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
    account_origin_id,
    account_destination_id,
    amount
) VALUES (
    $1,$2,$3
) RETURNING id, account_origin_id, account_destination_id, amount, created_at;
`

type CreateTransferParam struct {
	AccountOriginID      int64 `json:"account_origin_id"`
	AccountDestinationID int64 `json:"account_destination_id"`
	Amount               int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParam) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.AccountOriginID, arg.AccountDestinationID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.AccountOriginID,
		&i.AccountDestinationID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.AccountOriginID,
		&i.AccountDestinationID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `--name: ListTransfers :many
SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfers
WHERE
	account_origin_id = $1 OR
	account_destination_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfersParam struct {
	AccountOriginId      int64 `json:"account_origin_id"`
	AccountDestinationId int64 `json:"account_destination_id"`
	Limit                int32 `json:"limit"`
	Offset               int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParam) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers,
		arg.AccountOriginId,
		arg.AccountDestinationId,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transfers := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.AccountOriginID,
			&i.AccountDestinationID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		transfers = append(transfers, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transfers, nil
}
