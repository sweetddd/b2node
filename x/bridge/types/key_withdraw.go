package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WithdrawKeyPrefix is the prefix to retrieve all Withdraw
	WithdrawKeyPrefix = "Withdraw/value/"
)

// WithdrawKey returns the store key to retrieve a Withdraw from the index fields
func WithdrawKey(
	txHash string,
) []byte {
	var key []byte

	txHashBytes := []byte(txHash)
	key = append(key, txHashBytes...)
	key = append(key, []byte("/")...)

	return key
}
