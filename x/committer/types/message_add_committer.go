package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)

func NewMsgAddCommitter(from string, committer string) *MsgAddCommitterTx {
	return &MsgAddCommitterTx{
		From: from,
		Committer: committer,
	}
}

func (msg *MsgAddCommitterTx) Route() string {
	return RouterKey
}

func (msg *MsgAddCommitterTx) Type() string {
	return "AddCommitterTx"
}

func (msg *MsgAddCommitterTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddCommitterTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgAddCommitterTx) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if msg.Committer == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing committer")
	}
	return nil
}
