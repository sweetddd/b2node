//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/evmos/ethermint/x/bridge/types"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/evmos/ethermint/testutil/network"
	"github.com/evmos/ethermint/x/bridge/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithDeferentWithdrawState(t *testing.T, withdraws ...types.Withdraw) (*network.Network, error) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))
	if len(withdraws) != 0 {
		state.WithdrawList = append(state.WithdrawList, withdraws...)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, t.TempDir(), cfg)
}

func prepareCallerAndSignerGroup(ctx client.Context, flags []string, msgCreator string) {
	callerArgs := []string{"caller group", msgCreator, msgCreator}
	callerArgs = append(callerArgs, flags...)
	clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCallerGroup(), callerArgs)

	signerArgs := []string{"signer group", msgCreator, "3", msgCreator}
	signerArgs = append(signerArgs, flags...)
	clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateSignerGroup(), signerArgs)
}

func TestCreateWithdraw(t *testing.T) {
	net, err := networkWithDeferentWithdrawState(t)
	require.NoError(t, err)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "COIN_TYPE_BTC", "111", "xyz"}
	for _, tc := range []struct {
		desc     string
		idTxHash string

		args []string
		err  error
		code uint32
	}{
		{
			idTxHash: strconv.Itoa(0),

			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagGasPrices, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10000000000))).String()),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idTxHash,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			prepareCallerAndSignerGroup(ctx, tc.args, val.Address.String())
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateWithdraw(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdateWithdraw(t *testing.T) {
	withdraw := types.Withdraw{
		TxHash:     "0",
		From:       "xyz",
		To:         "xyz",
		CoinType:   types.CoinType_COIN_TYPE_BTC,
		Value:      111,
		Data:       "xyz",
		Status:     types.WithdrawStatus_WITHDRAW_STATUS_SIGNED,
		Signatures: map[string]string{"A": "A", "B": "B", "C": "C"},
		Creator:    "xyz",
	}
	net, err := networkWithDeferentWithdrawState(t, withdraw)
	require.NoError(t, err)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"WITHDRAW_STATUS_COMPLETED"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagGasPrices, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10000000000))).String()),
	}

	prepareCallerAndSignerGroup(ctx, common, val.Address.String())

	for _, tc := range []struct {
		desc     string
		idTxHash string

		args []string
		code uint32
		err  error
	}{
		{
			desc:     "valid",
			idTxHash: strconv.Itoa(0),

			args: common,
		},
		{
			desc:     "key not found",
			idTxHash: strconv.Itoa(100000),

			args: common,
			code: types.ErrIndexNotExist.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idTxHash,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateWithdraw(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestSignWithdraw(t *testing.T) {
	withdraw := types.Withdraw{
		TxHash:     "0",
		From:       "xyz",
		To:         "xyz",
		CoinType:   types.CoinType_COIN_TYPE_BTC,
		Value:      111,
		Data:       "xyz",
		Status:     types.WithdrawStatus_WITHDRAW_STATUS_PENDING,
		Signatures: map[string]string{"A": "A", "B": "B"},
		Creator:    "xyz",
	}
	net, err := networkWithDeferentWithdrawState(t, withdraw)
	require.NoError(t, err)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"ccc"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagGasPrices, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10000000000))).String()),
	}

	prepareCallerAndSignerGroup(ctx, common, val.Address.String())

	for _, tc := range []struct {
		desc     string
		idTxHash string

		args []string
		code uint32
		err  error
	}{
		{
			desc:     "valid",
			idTxHash: strconv.Itoa(0),

			args: common,
		},
		{
			desc:     "key not found",
			idTxHash: strconv.Itoa(100000),

			args: common,
			code: types.ErrIndexNotExist.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idTxHash,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdSignWithdraw(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestDeleteWithdraw(t *testing.T) {
	net, err := network.New(t, t.TempDir(), network.DefaultConfig())
	require.NoError(t, err)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "COIN_TYPE_BTC", "111", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagGasPrices, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10000000000))).String()),
	}
	args := []string{
		"0",
	}
	args = append(args, fields...)
	args = append(args, common...)
	prepareCallerAndSignerGroup(ctx, common, val.Address.String())
	_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateWithdraw(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc     string
		idTxHash string

		args []string
		code uint32
		err  error
	}{
		{
			desc:     "valid",
			idTxHash: strconv.Itoa(0),

			args: common,
		},
		{
			desc:     "key not found",
			idTxHash: strconv.Itoa(100000),

			args: common,
			code: types.ErrIndexNotExist.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idTxHash,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteWithdraw(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}
