package db

import (
	"context"
	//"math/big"
	//"github.com/google/uuid"
)

const addAccountBalance = `-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING id, name, cpf, secret, balance, created_at
`

type AddAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, addAccountBalance, arg.Amount, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Secret,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
	name,
	cpf,
	secret,
	balance
) VALUES (
	$1, $2, $3, $4
) RETURNING id, name, cpf, secret, balance, created_at
`

type CreateAccountParams struct {
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int64  `json:"balance"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Name, arg.Cpf, arg.Secret, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Secret,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, cpf, secret, balance, created_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Secret,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountByCpf = `-- name: GetAccountByCpf :one
SELECT id, name, cpf, secret, balance, created_at FROM accounts
WHERE cpf = $1 LIMIT 1
`

func (q *Queries) GetAccountByCpf(ctx context.Context, cpf string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByCpf, cpf)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Secret,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

//this query avoids deadlock
const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, name, cpf, secret, balance, created_at FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Secret,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, name, cpf, secret, balance, created_at FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountsParam struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParam) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts = []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Cpf,
			&i.Secret,
			&i.Balance,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING id, name, cpf, secret, balance, created_at
`

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Secret,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}
