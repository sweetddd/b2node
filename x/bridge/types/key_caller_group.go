package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CallerGroupKeyPrefix is the prefix to retrieve all CallerGroup
	CallerGroupKeyPrefix = "CallerGroup/value/"
)

// CallerGroupKey returns the store key to retrieve a CallerGroup from the index fields
func CallerGroupKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
