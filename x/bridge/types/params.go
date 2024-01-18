package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyCallerGroupName = []byte("CallerGroupName")
	// TODO: Determine the default value
	DefaultCallerGroupName string = "caller group"
)

var (
	KeySignerGroupName = []byte("SignerGroupName")
	// TODO: Determine the default value
	DefaultSignerGroupName string = "signer group"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	callerGroupName string,
	signerGroupName string,
) Params {
	return Params{
		CallerGroupName: callerGroupName,
		SignerGroupName: signerGroupName,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultCallerGroupName,
		DefaultSignerGroupName,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCallerGroupName, &p.CallerGroupName, validateCallerGroupName),
		paramtypes.NewParamSetPair(KeySignerGroupName, &p.SignerGroupName, validateSignerGroupName),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateCallerGroupName(p.CallerGroupName); err != nil {
		return err
	}

	if err := validateSignerGroupName(p.SignerGroupName); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateCallerGroupName validates the CallerGroupName param
func validateCallerGroupName(v interface{}) error {
	callerGroupName, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = callerGroupName

	return nil
}

// validateSignerGroupName validates the SignerGroupName param
func validateSignerGroupName(v interface{}) error {
	signerGroupName, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = signerGroupName

	return nil
}
