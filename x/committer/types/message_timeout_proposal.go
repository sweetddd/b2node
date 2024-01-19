package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)	

func NewMsgTimeoutProposal(
	id uint64,
	from string,
) *MsgTimeoutProposal {
	return &MsgTimeoutProposal{
		Id: id,
		From: from,
	}
}

func (msg *MsgTimeoutProposal) Route() string {
	return RouterKey
}

func (msg *MsgTimeoutProposal) Type() string {
	return "TimeoutProposal"
}

func (msg *MsgTimeoutProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTimeoutProposal) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgTimeoutProposal) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}

	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid from address")
	}

	return nil
}