package cli //nolint:dupl

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdListWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-withdraw",
		Short: "list all withdraw",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWithdrawRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WithdrawAll(context.Background(), params)
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

func CmdShowWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-withdraw [tx-hash]",
		Short: "shows a withdraw",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTxHash := args[0]

			params := &types.QueryGetWithdrawRequest{
				TxHash: argTxHash,
			}

			res, err := queryClient.Withdraw(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
