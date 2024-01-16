package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDeposit = "create_deposit"
	TypeMsgUpdateDeposit = "update_deposit"
	TypeMsgDeleteDeposit = "delete_deposit"
)

var _ sdk.Msg = &MsgCreateDeposit{}

func NewMsgCreateDeposit(
	creator string,
	index string,
	txHash string,
	from string,
	to string,
	coinType string,
	value uint64,
	data string,

) *MsgCreateDeposit {
	return &MsgCreateDeposit{
		Creator:  creator,
		Index:    index,
		TxHash:   txHash,
		From:     from,
		To:       to,
		CoinType: coinType,
		Value:    value,
		Data:     data,
	}
}

func (msg *MsgCreateDeposit) Route() string {
	return RouterKey
}

func (msg *MsgCreateDeposit) Type() string {
	return TypeMsgCreateDeposit
}

func (msg *MsgCreateDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDeposit{}

func NewMsgUpdateDeposit(
	creator string,
	index string,
	txHash string,
	from string,
	to string,
	coinType string,
	value uint64,
	data string,

) *MsgUpdateDeposit {
	return &MsgUpdateDeposit{
		Creator:  creator,
		Index:    index,
		TxHash:   txHash,
		From:     from,
		To:       to,
		CoinType: coinType,
		Value:    value,
		Data:     data,
	}
}

func (msg *MsgUpdateDeposit) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDeposit) Type() string {
	return TypeMsgUpdateDeposit
}

func (msg *MsgUpdateDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDeposit{}

func NewMsgDeleteDeposit(
	creator string,
	index string,

) *MsgDeleteDeposit {
	return &MsgDeleteDeposit{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDeposit) Type() string {
	return TypeMsgDeleteDeposit
}

func (msg *MsgDeleteDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
