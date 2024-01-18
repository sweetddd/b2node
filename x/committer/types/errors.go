package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/committer module sentinel errors
var (
	ErrSample = errors.Register(ModuleName, 1100, "sample error")
	ErrNotExistProposal = errors.Register(ModuleName, 1101, "proposal does not exist")
	ErrAccountPermission = errors.Register(ModuleName, 1102, "account not authorized")	
	ErrProposalStatus = errors.Register(ModuleName, 1103, "proposal status error")
	ErrInvalidProposal = errors.Register(ModuleName, 1104, "invalid proposal")
)
