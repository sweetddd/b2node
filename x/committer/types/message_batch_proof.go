package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/errors"
)	

func NewBatchProofMsg(
	id uint64,
	from string,
	proofHash string,
	stateRootHash string,
	startIndex uint64,
	endIndex uint64,
) *MsgBatchProofTx {
	return &MsgBatchProofTx{
		Id: id,
		From: from,
		ProofHash: proofHash,
		StateRootHash: stateRootHash,
		StartIndex: startIndex,
		EndIndex: endIndex,
	}
}

func (msg *MsgBatchProofTx) Route() string {
	return RouterKey
}

func (msg *MsgBatchProofTx) Type() string {
	return "BatchProofTx"
}

func (msg *MsgBatchProofTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBatchProofTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgBatchProofTx) ValidateBasic() error {
	if msg.From == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if msg.ProofHash == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing proof hash")
	}
	if msg.StateRootHash == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "missing state root hash")
	}
	return nil
}