package cli_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/evmos/ethermint/testutil/network"
	"github.com/evmos/ethermint/x/bridge/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func perpareCallerGroup(ctx client.Context, flags []string, msgCreator string) {
	args := []string{"caller group", msgCreator, msgCreator}
	args = append(args, flags...)
	clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCallerGroup(), args)
}

func TestCreateDeposit(t *testing.T) {
	net, err := network.New(t, t.TempDir(), network.DefaultConfig())
	require.NoError(t, err)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "111", "xyz"}
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
			perpareCallerGroup(ctx, tc.args, val.Address.String())
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateDeposit(), args)
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

func TestUpdateDeposit(t *testing.T) {
	net, err := network.New(t, t.TempDir(), network.DefaultConfig())
	require.NoError(t, err)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "111", "xyz"}
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
	perpareCallerGroup(ctx, common, val.Address.String())
	_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateDeposit(), args)
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
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idTxHash,
			}
			args = append(args, "finished")
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateDeposit(), args)
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

func TestDeleteDeposit(t *testing.T) {
	net, err := network.New(t, t.TempDir(), network.DefaultConfig())
	require.NoError(t, err)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "111", "xyz"}
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
	perpareCallerGroup(ctx, common, val.Address.String())
	_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateDeposit(), args)
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
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idTxHash,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteDeposit(), args)
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
