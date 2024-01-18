package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)	

func NewTimeoutProposalMsg(
	id uint64,
	from string,
) *MsgTimeoutProposalTx {
	return &MsgTimeoutProposalTx{
		Id: id,
		From: from,
	}
}

func (msg *MsgTimeoutProposalTx) Route() string {
	return RouterKey
}

func (msg *MsgTimeoutProposalTx) Type() string {
	return "TimeoutProposalTx"
}

func (msg *MsgTimeoutProposalTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTimeoutProposalTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgTimeoutProposalTx) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	return nil
}