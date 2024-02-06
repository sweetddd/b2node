package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWithdraw = "create_withdraw"
	TypeMsgUpdateWithdraw = "update_withdraw"
	TypeMsgSignWithdraw   = "sign_withdraw"
	TypeMsgDeleteWithdraw = "delete_withdraw"
)

var _ sdk.Msg = &MsgCreateWithdraw{}

func NewMsgCreateWithdraw(
	creator string,
	txID string,
	txHashList []string,
	encodedData string,
) *MsgCreateWithdraw {
	return &MsgCreateWithdraw{
		Creator:     creator,
		TxId:        txID,
		TxHashList:  txHashList,
		EncodedData: encodedData,
	}
}

func (msg *MsgCreateWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgCreateWithdraw) Type() string {
	return TypeMsgCreateWithdraw
}

func (msg *MsgCreateWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWithdraw{}

func NewMsgUpdateWithdraw(
	creator string,
	txId string,
	status WithdrawStatus,
) *MsgUpdateWithdraw {
	return &MsgUpdateWithdraw{
		Creator: creator,
		TxId:    txId,
		Status:  status,
	}
}

func (msg *MsgUpdateWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWithdraw) Type() string {
	return TypeMsgUpdateWithdraw
}

func (msg *MsgUpdateWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgSignWithdraw{}

func NewMsgSignWithdraw(
	creator string,
	txID string,
	signature string,
) *MsgSignWithdraw {
	return &MsgSignWithdraw{
		Creator:   creator,
		TxId:      txID,
		Signature: signature,
	}
}

func (msg *MsgSignWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgSignWithdraw) Type() string {
	return TypeMsgSignWithdraw
}

func (msg *MsgSignWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWithdraw{}

func NewMsgDeleteWithdraw(
	creator string,
	txID string,
) *MsgDeleteWithdraw {
	return &MsgDeleteWithdraw{
		Creator: creator,
		TxId:    txID,
	}
}

func (msg *MsgDeleteWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWithdraw) Type() string {
	return TypeMsgDeleteWithdraw
}

func (msg *MsgDeleteWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
