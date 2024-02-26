package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateSignerGroup{}, "bridge/CreateSignerGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateSignerGroup{}, "bridge/UpdateSignerGroup", nil)
	cdc.RegisterConcrete(&MsgDeleteSignerGroup{}, "bridge/DeleteSignerGroup", nil)
	cdc.RegisterConcrete(&MsgCreateCallerGroup{}, "bridge/CreateCallerGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateCallerGroup{}, "bridge/UpdateCallerGroup", nil)
	cdc.RegisterConcrete(&MsgDeleteCallerGroup{}, "bridge/DeleteCallerGroup", nil)
	cdc.RegisterConcrete(&MsgCreateDeposit{}, "bridge/CreateDeposit", nil)
	cdc.RegisterConcrete(&MsgUpdateDeposit{}, "bridge/UpdateDeposit", nil)
	cdc.RegisterConcrete(&MsgDeleteDeposit{}, "bridge/DeleteDeposit", nil)
	cdc.RegisterConcrete(&MsgCreateWithdraw{}, "bridge/CreateWithdraw", nil)
	cdc.RegisterConcrete(&MsgUpdateWithdraw{}, "bridge/UpdateWithdraw", nil)
	cdc.RegisterConcrete(&MsgDeleteWithdraw{}, "bridge/DeleteWithdraw", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignerGroup{},
		&MsgUpdateSignerGroup{},
		&MsgDeleteSignerGroup{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCallerGroup{},
		&MsgUpdateCallerGroup{},
		&MsgDeleteCallerGroup{},
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
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
