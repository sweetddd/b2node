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
	txHash string,
) []byte {
	var key []byte

	txHashBytes := []byte(txHash)
	key = append(key, txHashBytes...)
	key = append(key, []byte("/")...)

	return key
}

// WithdrawStatusKey returns the store key to retrieve a Withdraw from it's status and txHash
func WithdrawStatusKey(status string, txHash string) []byte {
	var key []byte
	key = append(key, []byte(status)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(txHash)...)
	key = append(key, []byte("/")...)
	return key
}
