package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)

func NewMsgRemoveCommitter(from string, committer string) *MsgRemoveCommitterTx {
	return &MsgRemoveCommitterTx{
		From: from,
		Committer: committer,
	}
}

func (msg *MsgRemoveCommitterTx) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCommitterTx) Type() string {
	return "RemoveCommitterTx"
}

func (msg *MsgRemoveCommitterTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCommitterTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRemoveCommitterTx) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if msg.Committer == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing committer")
	}
	return nil
}
