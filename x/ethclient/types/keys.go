package types

import fmt "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "ethclient"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ethclient"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	StorageKeyPrefix = "Storage"
	ProofKeyPrefix   = "Proof"
)

// Returns storage key in the form "Storage/{address}/{key}/{block}" as bytes
// ethereum block is saved as hexencoded string of block height
// NOTE: this can be made more efficient
func GetStorageKey(address, key string, block int64) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s/%d", StorageKeyPrefix, address, key, block))
}

// Returns storage key in the form "Proof/{address}/{key}/{block}" as bytes
// NOTE: this can be made more efficient
func GetProofKey(address, key string, block int64) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s/%d", ProofKeyPrefix, address, key, block))
}
