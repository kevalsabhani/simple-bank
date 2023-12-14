package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transferTxResult, err := testStore.TransferTx(context.Background(), TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        int64(10),
	})

	require.NoError(t, err)
	require.NotEmpty(t, transferTxResult)
	require.NotEmpty(t, transferTxResult.FromAccount)
	require.NotEmpty(t, transferTxResult.ToAccount)
	require.NotEmpty(t, transferTxResult.FromEntry)
	require.NotEmpty(t, transferTxResult.ToEntry)
	require.Equal(t, account1.ID, transferTxResult.FromAccount.ID)
	require.Equal(t, account2.ID, transferTxResult.ToAccount.ID)
	diff1 := account1.Balance - transferTxResult.FromAccount.Balance
	diff2 := transferTxResult.ToAccount.Balance - account2.Balance
	require.Equal(t, diff1, diff2)
}
