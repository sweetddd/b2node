package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgBatchProofTx{}, "committer/BatchProofTx", nil)
	cdc.RegisterConcrete(&MsgTapRootTx{}, "committer/TapRootTx", nil)
	cdc.RegisterConcrete(&MsgTimeoutProposalTx{}, "committer/TimeoutProposalTx", nil)
} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil), 
	&MsgBatchProofTx{},
	&MsgTapRootTx{},
	&MsgTimeoutProposalTx{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
