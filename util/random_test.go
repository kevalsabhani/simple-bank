package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.Len(t, owner, 6)
}

func TestRandomBalance(t *testing.T) {
	balance := RandomBalance()
	require.GreaterOrEqual(t, balance, int64(0))
	require.LessOrEqual(t, balance, int64(10000))
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	require.Contains(t, []string{"INR", "USD", "CAD", "EUR"}, currency)
}

func TestRandomAccountID(t *testing.T) {
	accId := RandomAccountID()
	require.GreaterOrEqual(t, accId, int64(1))
	require.LessOrEqual(t, accId, int64(100000))
}
