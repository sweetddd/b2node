package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)	

func NewMsgSubmitProof(
	id uint64,
	from string,
	proofHash string,
	stateRootHash string,
	startIndex uint64,
	endIndex uint64,
) *MsgSubmitProof {
	return &MsgSubmitProof{
		Id: id,
		From: from,
		ProofHash: proofHash,
		StateRootHash: stateRootHash,
		StartIndex: startIndex,
		EndIndex: endIndex,
	}
}

func (msg *MsgSubmitProof) Route() string {
	return RouterKey
}

func (msg *MsgSubmitProof) Type() string {
	return "SubmitProof"
}

func (msg *MsgSubmitProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitProof) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgSubmitProof) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid from address")
	}

	if msg.ProofHash == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing proof hash")
	}

	if msg.StateRootHash == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing state root hash")
	}

	if msg.StartIndex > msg.EndIndex {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "start index must be less than end index")
	}

	return nil
}