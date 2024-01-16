package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WithdrawKeyPrefix is the prefix to retrieve all Withdraw
	WithdrawKeyPrefix = "Withdraw/value/"
)

// WithdrawKey returns the store key to retrieve a Withdraw from the index fields
func WithdrawKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
