package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgAddCommitter(from string, committer string) *MsgAddCommitter {
	return &MsgAddCommitter{
		From:      from,
		Committer: committer,
	}
}

func (msg *MsgAddCommitter) Route() string {
	return RouterKey
}

func (msg *MsgAddCommitter) Type() string {
	return "AddCommitter"
}

func (msg *MsgAddCommitter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddCommitter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgAddCommitter) ValidateBasic() error {
	return validateFromAndCommitter(msg.From, msg.Committer)
}
