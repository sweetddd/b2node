package types

const (
	// Proposal Status
	Voting_Status = 0
	Pending_Status = 1
	Succeed_Status = 2
	Timeout_Status = 3
	
	// Proposal Params
	DefaultTimeoutBlockPeriod = 10000
)

// type Proposal struct {
// 		// Id is the unique id of the proposal
// 		Id uint64 `json:"id"`
// 		// Proposer is the address of the proposer
// 		Proposer string `json:"proposer"`
// 		// ProofHash is the hash of the proof
// 		ProofHash string `json:"proof_hash"`
// 		// StateRootHash is the hash of the state root
// 		StateRootHash string `json:"state_root_hash"`
// 		// StartIndex is the start index of the batch
// 		StartIndex uint64 `json:"start_index"`
// 		// EndIndex is the end index of the batch
// 		EndIndex uint64 `json:"end_index"`
// 		// BlockHeight is the block height of the proposal
// 		BlockHight uint64 `json:"block_hight"`
// 		// Status is the status of the proposal
// 		Status uint64 `json:"status"`
// 		// BitcoinTxHash is the hash of the bitcoin tx
// 		BitcoinTxHash string `json:"bitcoin_tx"`
// 		// Winner is the winner of the proposal
// 		Winner string `json:"winner"`
// 		// VotedListPhaseCommit is the list of committers who voted for the proposal in the commit phase
// 		VotedListPhaseCommit []string `json:"voted_list_phase_commit"`
// 		// VotedListPhaseTimeout is the list of committers who voted for the proposal in the timeout phase
// 		VotedListPhaseTimeout []string `json:"voted_list_phase_timeout"`
// }

// type LastProposal struct {
// 	ID uint64
// }