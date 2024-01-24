package bridge

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/evmos/ethermint/testutil/bridge/sample"
	bridgesimulation "github.com/evmos/ethermint/x/bridge/simulation"
	"github.com/evmos/ethermint/x/bridge/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bridgesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateSignerGroup = "op_weight_msg_signer_group" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSignerGroup int = 100

	opWeightMsgUpdateSignerGroup = "op_weight_msg_signer_group" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSignerGroup int = 100

	opWeightMsgDeleteSignerGroup = "op_weight_msg_signer_group" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSignerGroup int = 100

	opWeightMsgCreateCallerGroup = "op_weight_msg_caller_group" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCallerGroup int = 100

	opWeightMsgUpdateCallerGroup = "op_weight_msg_caller_group" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCallerGroup int = 100

	opWeightMsgDeleteCallerGroup = "op_weight_msg_caller_group" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCallerGroup int = 100

	opWeightMsgCreateDeposit = "op_weight_msg_deposit" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeposit int = 100

	opWeightMsgUpdateDeposit = "op_weight_msg_deposit" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeposit int = 100

	opWeightMsgDeleteDeposit = "op_weight_msg_deposit" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDeposit int = 100

	opWeightMsgCreateWithdraw = "op_weight_msg_withdraw" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWithdraw int = 100

	opWeightMsgUpdateWithdraw = "op_weight_msg_withdraw" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWithdraw int = 100

	opWeightMsgDeleteWithdraw = "op_weight_msg_withdraw" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWithdraw int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bridgeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		SignerGroupList: []types.SignerGroup{
			{
				Creator: sample.AccAddress(),
				Name:    "0",
			},
			{
				Creator: sample.AccAddress(),
				Name:    "1",
			},
		},
		CallerGroupList: []types.CallerGroup{
			{
				Creator: sample.AccAddress(),
				Name:    "0",
			},
			{
				Creator: sample.AccAddress(),
				Name:    "1",
			},
		},
		DepositList: []types.Deposit{
			{
				Creator: sample.AccAddress(),
				TxHash:  "0",
			},
			{
				Creator: sample.AccAddress(),
				TxHash:  "1",
			},
		},
		WithdrawList: []types.Withdraw{
			{
				Creator: sample.AccAddress(),
				TxHash:  "0",
			},
			{
				Creator: sample.AccAddress(),
				TxHash:  "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bridgeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateSignerGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSignerGroup, &weightMsgCreateSignerGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSignerGroup = defaultWeightMsgCreateSignerGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSignerGroup,
		bridgesimulation.SimulateMsgCreateSignerGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSignerGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSignerGroup, &weightMsgUpdateSignerGroup, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSignerGroup = defaultWeightMsgUpdateSignerGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSignerGroup,
		bridgesimulation.SimulateMsgUpdateSignerGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSignerGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSignerGroup, &weightMsgDeleteSignerGroup, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSignerGroup = defaultWeightMsgDeleteSignerGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSignerGroup,
		bridgesimulation.SimulateMsgDeleteSignerGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCallerGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCallerGroup, &weightMsgCreateCallerGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCallerGroup = defaultWeightMsgCreateCallerGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCallerGroup,
		bridgesimulation.SimulateMsgCreateCallerGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCallerGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCallerGroup, &weightMsgUpdateCallerGroup, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCallerGroup = defaultWeightMsgUpdateCallerGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCallerGroup,
		bridgesimulation.SimulateMsgUpdateCallerGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCallerGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCallerGroup, &weightMsgDeleteCallerGroup, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCallerGroup = defaultWeightMsgDeleteCallerGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCallerGroup,
		bridgesimulation.SimulateMsgDeleteCallerGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDeposit, &weightMsgCreateDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDeposit = defaultWeightMsgCreateDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDeposit,
		bridgesimulation.SimulateMsgCreateDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDeposit, &weightMsgUpdateDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDeposit = defaultWeightMsgUpdateDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDeposit,
		bridgesimulation.SimulateMsgUpdateDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDeposit, &weightMsgDeleteDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDeposit = defaultWeightMsgDeleteDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDeposit,
		bridgesimulation.SimulateMsgDeleteDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWithdraw, &weightMsgCreateWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWithdraw = defaultWeightMsgCreateWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWithdraw,
		bridgesimulation.SimulateMsgCreateWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWithdraw, &weightMsgUpdateWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWithdraw = defaultWeightMsgUpdateWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWithdraw,
		bridgesimulation.SimulateMsgUpdateWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWithdraw, &weightMsgDeleteWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWithdraw = defaultWeightMsgDeleteWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWithdraw,
		bridgesimulation.SimulateMsgDeleteWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
