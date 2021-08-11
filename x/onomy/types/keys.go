package types

const (
	// ModuleName defines the module name.
	ModuleName = "onomy"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_onomy"

	// this line is used by starport scaffolding # ibc/keys/name.
)

// this line is used by starport scaffolding # ibc/keys/port.

// KeyPrefix return []byte representation of the prefix.
func KeyPrefix(p string) []byte {
	return []byte(p)
}