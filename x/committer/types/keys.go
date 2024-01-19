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

  // ProposalKeyPrefix is the prefix to retrieve all Proposal
	ProposalKeyPrefix = "proposal/"

	// LastProposalIdKey defines the key to store the last proposal id
	LastProposalId = "last_proposal_id"

	// CommitterKeyPrefix is the prefix to retrieve all Committer			
	CommitterKeyPrefix = "committer/"

	KeyPrefixParams = "params/"
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
