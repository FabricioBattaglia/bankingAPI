package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Name:    "fabricio",
		Cpf:     "123.456.789-10",
		Secret:  "123456",
		Balance: 1000,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	//check if error is nil
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.Cpf, account.Cpf)
	require.Equal(t, arg.Secret, account.Secret)
	require.Equal(t, arg.Balance, account.Balance)

	//was: NotZero
	require.NotEmpty(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createTestAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Name, account2.Name)
	require.Equal(t, account.Cpf, account2.Cpf)
	require.Equal(t, account.Secret, account2.Secret)
	require.Equal(t, account.Balance, account2.Balance)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createTestAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: 2000,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Name, account2.Name)
	require.Equal(t, account.Cpf, account2.Cpf)
	require.Equal(t, account.Secret, account2.Secret)
	require.Equal(t, arg.Balance, account2.Balance)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createTestAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestAccount(t)
	}

	arg := ListAccountsParam{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
