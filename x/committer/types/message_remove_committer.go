package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)

func NewMsgRemoveCommitter(from string, committer string) *MsgRemoveCommitter {
	return &MsgRemoveCommitter{
		From: from,
		Committer: committer,
	}
}

func (msg *MsgRemoveCommitter) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCommitter) Type() string {
	return "RemoveCommitter"
}

func (msg *MsgRemoveCommitter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCommitter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRemoveCommitter) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}

	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid from address")
	}

	if msg.Committer == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing committer")
	}

	_, err = sdk.AccAddressFromBech32(msg.Committer)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid committer address")
	}

	return nil
}
