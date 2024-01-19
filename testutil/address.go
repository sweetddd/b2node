package testutil

import (
	ed25519 "github.com/tendermint/tendermint/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}