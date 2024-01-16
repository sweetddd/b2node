package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSigner = "create_signer"
	TypeMsgUpdateSigner = "update_signer"
	TypeMsgDeleteSigner = "delete_signer"
)

var _ sdk.Msg = &MsgCreateSigner{}

func NewMsgCreateSigner(creator string, address string, name string) *MsgCreateSigner {
	return &MsgCreateSigner{
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgCreateSigner) Route() string {
	return RouterKey
}

func (msg *MsgCreateSigner) Type() string {
	return TypeMsgCreateSigner
}

func (msg *MsgCreateSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSigner{}

func NewMsgUpdateSigner(creator string, id uint64, address string, name string) *MsgUpdateSigner {
	return &MsgUpdateSigner{
		Id:      id,
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgUpdateSigner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSigner) Type() string {
	return TypeMsgUpdateSigner
}

func (msg *MsgUpdateSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSigner{}

func NewMsgDeleteSigner(creator string, id uint64) *MsgDeleteSigner {
	return &MsgDeleteSigner{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSigner) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSigner) Type() string {
	return TypeMsgDeleteSigner
}

func (msg *MsgDeleteSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
