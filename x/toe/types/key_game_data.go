package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GameDataKeyPrefix is the prefix to retrieve all GameData
	GameDataKeyPrefix = "GameData/value/"
)

// GameDataKey returns the store key to retrieve a GameData from the index fields
func GameDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
