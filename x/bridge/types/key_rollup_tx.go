package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RollupTxKeyPrefix is the prefix to retrieve all RollupTx
	RollupTxKeyPrefix = "RollupTx/value/"
)

// RollupTxKey returns the store key to retrieve a RollupTx from the index fields
func RollupTxKey(
	txHash string,
) []byte {
	var key []byte

	txHashBytes := []byte(txHash)
	key = append(key, txHashBytes...)
	key = append(key, []byte("/")...)

	return key
}
