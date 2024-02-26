package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DepositKeyPrefix is the prefix to retrieve all Deposit
	DepositKeyPrefix = "Deposit/value/"
)

// DepositKey returns the store key to retrieve a Deposit from the index fields
func DepositKey(
	txHash string,
) []byte {
	var key []byte

	txHashBytes := []byte(txHash)
	key = append(key, txHashBytes...)
	key = append(key, []byte("/")...)

	return key
}
