package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WithdrawKeyPrefix is the prefix to retrieve all Withdraw
	WithdrawKeyPrefix = "Withdraw/value/"
	// WithdrawStatusKeyPrefix is the prefix to retrieve all Withdraw by status
	WithdrawStatusKeyPrefix = "Withdraw/status/"
)

// WithdrawKey returns the store key to retrieve a Withdraw from the index fields
func WithdrawKey(
	txID string,
) []byte {
	var key []byte

	txIDBytes := []byte(txID)
	key = append(key, txIDBytes...)
	key = append(key, []byte("/")...)

	return key
}

// WithdrawStatusKey returns the store key to retrieve a Withdraw from it's status and txId
func WithdrawStatusKey(status string, txID string) []byte {
	var key []byte
	key = append(key, []byte(status)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(txID)...)
	key = append(key, []byte("/")...)
	return key
}
