package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCaller = "create_caller"
	TypeMsgUpdateCaller = "update_caller"
	TypeMsgDeleteCaller = "delete_caller"
)

var _ sdk.Msg = &MsgCreateCaller{}

func NewMsgCreateCaller(creator string, address string) *MsgCreateCaller {
	return &MsgCreateCaller{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateCaller) Route() string {
	return RouterKey
}

func (msg *MsgCreateCaller) Type() string {
	return TypeMsgCreateCaller
}

func (msg *MsgCreateCaller) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCaller) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCaller) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCaller{}

func NewMsgUpdateCaller(creator string, id uint64, address string) *MsgUpdateCaller {
	return &MsgUpdateCaller{
		Id:      id,
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateCaller) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCaller) Type() string {
	return TypeMsgUpdateCaller
}

func (msg *MsgUpdateCaller) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCaller) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCaller) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCaller{}

func NewMsgDeleteCaller(creator string, id uint64) *MsgDeleteCaller {
	return &MsgDeleteCaller{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteCaller) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCaller) Type() string {
	return TypeMsgDeleteCaller
}

func (msg *MsgDeleteCaller) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCaller) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCaller) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
