package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CallerList:   []Caller{},
		DepositList:  []Deposit{},
		WithdrawList: []Withdraw{},
		SignerList:   []Signer{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in caller
	callerIdMap := make(map[uint64]bool)
	callerCount := gs.GetCallerCount()
	for _, elem := range gs.CallerList {
		if _, ok := callerIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for caller")
		}
		if elem.Id >= callerCount {
			return fmt.Errorf("caller id should be lower or equal than the last id")
		}
		callerIdMap[elem.Id] = true
	}
	// Check for duplicated index in deposit
	depositIndexMap := make(map[string]struct{})

	for _, elem := range gs.DepositList {
		index := string(DepositKey(elem.Index))
		if _, ok := depositIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for deposit")
		}
		depositIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in withdraw
	withdrawIndexMap := make(map[string]struct{})

	for _, elem := range gs.WithdrawList {
		index := string(WithdrawKey(elem.Index))
		if _, ok := withdrawIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for withdraw")
		}
		withdrawIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in signer
	signerIdMap := make(map[uint64]bool)
	signerCount := gs.GetSignerCount()
	for _, elem := range gs.SignerList {
		if _, ok := signerIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for signer")
		}
		if elem.Id >= signerCount {
			return fmt.Errorf("signer id should be lower or equal than the last id")
		}
		signerIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
