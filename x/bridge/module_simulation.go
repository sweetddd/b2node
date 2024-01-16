package bridge

import (
	"github.com/evmos/ethermint/testutil/bridge/sample"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
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
	opWeightMsgCreateCaller = "op_weight_msg_caller"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCaller int = 100

	opWeightMsgUpdateCaller = "op_weight_msg_caller"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCaller int = 100

	opWeightMsgDeleteCaller = "op_weight_msg_caller"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCaller int = 100

	opWeightMsgCreateDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeposit int = 100

	opWeightMsgUpdateDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeposit int = 100

	opWeightMsgDeleteDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDeposit int = 100

	opWeightMsgCreateWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWithdraw int = 100

	opWeightMsgUpdateWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWithdraw int = 100

	opWeightMsgDeleteWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWithdraw int = 100

	opWeightMsgCreateSigner = "op_weight_msg_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSigner int = 100

	opWeightMsgUpdateSigner = "op_weight_msg_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSigner int = 100

	opWeightMsgDeleteSigner = "op_weight_msg_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSigner int = 100

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
		CallerList: []types.Caller{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CallerCount: 2,
		DepositList: []types.Deposit{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		WithdrawList: []types.Withdraw{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		SignerList: []types.Signer{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SignerCount: 2,
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

	var weightMsgCreateCaller int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCaller, &weightMsgCreateCaller, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCaller = defaultWeightMsgCreateCaller
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCaller,
		bridgesimulation.SimulateMsgCreateCaller(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCaller int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCaller, &weightMsgUpdateCaller, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCaller = defaultWeightMsgUpdateCaller
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCaller,
		bridgesimulation.SimulateMsgUpdateCaller(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCaller int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCaller, &weightMsgDeleteCaller, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCaller = defaultWeightMsgDeleteCaller
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCaller,
		bridgesimulation.SimulateMsgDeleteCaller(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgCreateSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSigner, &weightMsgCreateSigner, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSigner = defaultWeightMsgCreateSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSigner,
		bridgesimulation.SimulateMsgCreateSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSigner, &weightMsgUpdateSigner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSigner = defaultWeightMsgUpdateSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSigner,
		bridgesimulation.SimulateMsgUpdateSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSigner, &weightMsgDeleteSigner, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSigner = defaultWeightMsgDeleteSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSigner,
		bridgesimulation.SimulateMsgDeleteSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
