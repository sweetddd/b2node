package types //nolint:dupl

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSignerGroup = "create_signer_group"
	TypeMsgUpdateSignerGroup = "update_signer_group"
	TypeMsgDeleteSignerGroup = "delete_signer_group"
)

var _ sdk.Msg = &MsgCreateSignerGroup{}

func NewMsgCreateSignerGroup(
	creator string,
	name string,
	admin string,
	threshold uint32,
	members []string,
) *MsgCreateSignerGroup {
	return &MsgCreateSignerGroup{
		Creator:   creator,
		Name:      name,
		Admin:     admin,
		Threshold: threshold,
		Members:   members,
	}
}

func (msg *MsgCreateSignerGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignerGroup) Type() string {
	return TypeMsgCreateSignerGroup
}

func (msg *MsgCreateSignerGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignerGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignerGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSignerGroup{}

func NewMsgUpdateSignerGroup(
	creator string,
	name string,
	admin string,
	threshold uint32,
	members []string,
) *MsgUpdateSignerGroup {
	return &MsgUpdateSignerGroup{
		Creator:   creator,
		Name:      name,
		Admin:     admin,
		Threshold: threshold,
		Members:   members,
	}
}

func (msg *MsgUpdateSignerGroup) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSignerGroup) Type() string {
	return TypeMsgUpdateSignerGroup
}

func (msg *MsgUpdateSignerGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSignerGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSignerGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSignerGroup{}

func NewMsgDeleteSignerGroup(
	creator string,
	name string,
) *MsgDeleteSignerGroup {
	return &MsgDeleteSignerGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDeleteSignerGroup) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSignerGroup) Type() string {
	return TypeMsgDeleteSignerGroup
}

func (msg *MsgDeleteSignerGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSignerGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSignerGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
