package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)

func NewMsgAddCommitter(from string, committer string) *MsgAddCommitter {
	return &MsgAddCommitter{
		From: from,
		Committer: committer,
	}
}

func (msg *MsgAddCommitter) Route() string {
	return RouterKey
}

func (msg *MsgAddCommitter) Type() string {
	return "AddCommitter"
}

func (msg *MsgAddCommitter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddCommitter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgAddCommitter) ValidateBasic() error {
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
