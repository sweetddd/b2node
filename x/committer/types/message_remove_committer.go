package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgRemoveCommitter(from string, committer string) *MsgRemoveCommitter {
	return &MsgRemoveCommitter{
		From:      from,
		Committer: committer,
	}
}

func (msg *MsgRemoveCommitter) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCommitter) Type() string {
	return "RemoveCommitter"
}

func (msg *MsgRemoveCommitter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCommitter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRemoveCommitter) ValidateBasic() error {
	return validateFromAndCommitter(msg.From, msg.Committer)
}
