package bitcoin_test

import (
	"testing"

	"github.com/evmos/ethermint/bitcoin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestGetStateRoot(t *testing.T) {
	cfg := bitcoin.StateConfig{
		Host:   "localhost",
		Port:   5432,
		User:   "state_user",
		Pass:   "state_password",
		DBName: "state_db",
	}
	items, err := bitcoin.GetStateRoot(cfg, 1)
	require.NoError(t, err)
	for _, item := range items {
		require.Equal(t, int64(1299), item.BlockNum)
		require.Equal(t, "0x1cc9e812fdad14a03f6e3c8563393d0e0c155dbd1d54361f41b33da46b087294", item.StateRoot)
	}
}
