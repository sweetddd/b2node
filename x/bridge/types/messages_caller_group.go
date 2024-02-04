package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCallerGroup = "create_caller_group"
	TypeMsgUpdateCallerGroup = "update_caller_group"
	TypeMsgDeleteCallerGroup = "delete_caller_group"
)

var _ sdk.Msg = &MsgCreateCallerGroup{}

func NewMsgCreateCallerGroup(
	creator string,
	name string,
	admin string,
	members []string,
) *MsgCreateCallerGroup {
	return &MsgCreateCallerGroup{
		Creator: creator,
		Name:    name,
		Admin:   admin,
		Members: members,
	}
}

func (msg *MsgCreateCallerGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateCallerGroup) Type() string {
	return TypeMsgCreateCallerGroup
}

func (msg *MsgCreateCallerGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCallerGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCallerGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCallerGroup{}

func NewMsgUpdateCallerGroup(
	creator string,
	name string,
	admin string,
	members []string,
) *MsgUpdateCallerGroup {
	return &MsgUpdateCallerGroup{
		Creator: creator,
		Name:    name,
		Admin:   admin,
		Members: members,
	}
}

func (msg *MsgUpdateCallerGroup) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCallerGroup) Type() string {
	return TypeMsgUpdateCallerGroup
}

func (msg *MsgUpdateCallerGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCallerGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCallerGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCallerGroup{}

func NewMsgDeleteCallerGroup(
	creator string,
	name string,
) *MsgDeleteCallerGroup {
	return &MsgDeleteCallerGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDeleteCallerGroup) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCallerGroup) Type() string {
	return TypeMsgDeleteCallerGroup
}

func (msg *MsgDeleteCallerGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCallerGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCallerGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
