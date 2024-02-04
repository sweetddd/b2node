package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func validateFromAndCommitter(from string, committer string) error {
	if from == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}

	_, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid from address")
	}

	if committer == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing committer")
	}

	_, err = sdk.AccAddressFromBech32(committer)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid committer address")
	}

	return nil
}
