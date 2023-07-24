package db

import (
	"context"
	"github.com/Rasulid/Go_Backend/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createTestTransfer(t *testing.T, account1, account2 Account) Transfer {
	args := CreateTransferParams{
		ToAccountID:   account1.ID,
		FromAccountID: account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Amount, args.Amount)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	account1 := createTestAccount(t)
	account2 := createTestAccount(t)

	createTestTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createTestAccount(t)
	account2 := createTestAccount(t)

	transfer1 := createTestTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createTestAccount(t)
	account2 := createTestAccount(t)

	for i := 0; i < 5; i++ {
		createTestTransfer(t, account1, account2)
		createTestTransfer(t, account1, account2)

	}

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}

}
