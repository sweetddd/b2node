package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCaller{}, "bridge/CreateCaller", nil)
	cdc.RegisterConcrete(&MsgUpdateCaller{}, "bridge/UpdateCaller", nil)
	cdc.RegisterConcrete(&MsgDeleteCaller{}, "bridge/DeleteCaller", nil)
	cdc.RegisterConcrete(&MsgCreateDeposit{}, "bridge/CreateDeposit", nil)
	cdc.RegisterConcrete(&MsgUpdateDeposit{}, "bridge/UpdateDeposit", nil)
	cdc.RegisterConcrete(&MsgDeleteDeposit{}, "bridge/DeleteDeposit", nil)
	cdc.RegisterConcrete(&MsgCreateWithdraw{}, "bridge/CreateWithdraw", nil)
	cdc.RegisterConcrete(&MsgUpdateWithdraw{}, "bridge/UpdateWithdraw", nil)
	cdc.RegisterConcrete(&MsgDeleteWithdraw{}, "bridge/DeleteWithdraw", nil)
	cdc.RegisterConcrete(&MsgCreateSigner{}, "bridge/CreateSigner", nil)
	cdc.RegisterConcrete(&MsgUpdateSigner{}, "bridge/UpdateSigner", nil)
	cdc.RegisterConcrete(&MsgDeleteSigner{}, "bridge/DeleteSigner", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCaller{},
		&MsgUpdateCaller{},
		&MsgDeleteCaller{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeposit{},
		&MsgUpdateDeposit{},
		&MsgDeleteDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWithdraw{},
		&MsgUpdateWithdraw{},
		&MsgDeleteWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSigner{},
		&MsgUpdateSigner{},
		&MsgDeleteSigner{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
