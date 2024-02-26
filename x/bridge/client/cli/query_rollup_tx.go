package cli //nolint:dupl

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdListRollupTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-rollup-tx",
		Short: "list all rollupTx",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRollupTxRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RollupTxAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowRollupTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-rollup-tx [tx-hash]",
		Short: "shows a rollupTx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTxHash := args[0]

			params := &types.QueryGetRollupTxRequest{
				TxHash: argTxHash,
			}

			res, err := queryClient.RollupTx(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
