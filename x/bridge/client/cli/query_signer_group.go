package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdListSignerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-signer-group",
		Short: "list all signerGroup",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSignerGroupRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SignerGroupAll(context.Background(), params)
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

func CmdShowSignerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-signer-group [name]",
		Short: "shows a signerGroup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]

			params := &types.QueryGetSignerGroupRequest{
				Name: argName,
			}

			res, err := queryClient.SignerGroup(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
