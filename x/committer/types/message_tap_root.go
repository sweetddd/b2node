package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)	

func NewTapRootMsg(
	id uint64,
	from string,
	txHash string,
) *MsgTapRootTx {
	return &MsgTapRootTx{
		Id: id,
		From: from,
		BitcoinTxHash: txHash,
	}
}

func (msg *MsgTapRootTx) Route() string {
	return RouterKey
}

func (msg *MsgTapRootTx) Type() string {
	return "TapRootTx"
}

func (msg *MsgTapRootTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTapRootTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgTapRootTx) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if msg.BitcoinTxHash == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing bitcoin tx hash")
	}
	return nil
}
	