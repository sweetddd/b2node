package types

const (
	// ModuleName defines the module name
	ModuleName = "committer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_committer"
)

// prefix bytes for the committer store
const (
	prefixLastProposalID = iota + 1
	prefixCommitter
	prefixProposal
	prefixParams
)

// KVStore key prefixes
var (
	KeyPrefixLastProposalID = []byte{prefixLastProposalID}
	KeyPrefixCommitter      = []byte{prefixCommitter}
	KeyPrefixProposal       = []byte{prefixProposal}
	KeyPrefixParams         = []byte{prefixParams}
)

func KeyPrefix(prefix []byte, p string) []byte {
	return append(prefix, []byte(p)...)
}
