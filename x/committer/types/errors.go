package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/committer module sentinel errors
var (
	ErrSample            = errors.Register(ModuleName, 1100, "sample error")
	ErrNotExistProposal  = errors.Register(ModuleName, 1101, "proposal does not exist")
	ErrAccountPermission = errors.Register(ModuleName, 1102, "account not authorized")
	ErrProposalStatus    = errors.Register(ModuleName, 1103, "proposal status error")
	ErrInvalidProposal   = errors.Register(ModuleName, 1104, "invalid proposal")
	ErrProposalTimeout   = errors.Register(ModuleName, 1105, "proposal timeout")
	ErrExistCommitter    = errors.Register(ModuleName, 1106, "committer already exists")
	ErrNotExistCommitter = errors.Register(ModuleName, 1107, "committer does not exist")
	ErrAlreadyVoted      = errors.Register(ModuleName, 1108, "already voted")
)

func (p Params) GetAdminPolicyAccount(policyType PolicyType) string {
	for _, admin := range p.AdminPolicy {
		if admin.PolicyType == policyType {
			return admin.Address
		}
	}
	return ""
}

