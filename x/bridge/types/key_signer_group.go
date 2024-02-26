package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SignerGroupKeyPrefix is the prefix to retrieve all SignerGroup
	SignerGroupKeyPrefix = "SignerGroup/value/"
)

// SignerGroupKey returns the store key to retrieve a SignerGroup from the index fields
func SignerGroupKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
