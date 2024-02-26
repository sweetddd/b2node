package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgBitcoinTx(
	id uint64,
	from string,
	txHash string,
) *MsgBitcoinTx {
	return &MsgBitcoinTx{
		Id:            id,
		From:          from,
		BitcoinTxHash: txHash,
	}
}

func (msg *MsgBitcoinTx) Route() string {
	return RouterKey
}

func (msg *MsgBitcoinTx) Type() string {
	return "BitcoinTx"
}

func (msg *MsgBitcoinTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBitcoinTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgBitcoinTx) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}

	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid from address")
	}

	if msg.BitcoinTxHash == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing bitcoin tx hash")
	}
	return nil
}
