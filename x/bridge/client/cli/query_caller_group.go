package cli //nolint:dupl

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdListCallerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-caller-group",
		Short: "list all callerGroup",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCallerGroupRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CallerGroupAll(context.Background(), params)
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

func CmdShowCallerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-caller-group [name]",
		Short: "shows a callerGroup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]

			params := &types.QueryGetCallerGroupRequest{
				Name: argName,
			}

			res, err := queryClient.CallerGroup(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
