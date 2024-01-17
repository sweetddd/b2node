package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cobra"
	"strings"
)

func CmdCreateSignerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-signer-group [name] [admin] [members]",
		Short: "Create a new signerGroup",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]

			// Get value arguments
			argAdmin := args[1]
			argMembers := strings.Split(args[2], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignerGroup(
				clientCtx.GetFromAddress().String(),
				indexName,
				argAdmin,
				argMembers,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateSignerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-signer-group [name] [admin] [members]",
		Short: "Update a signerGroup",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]

			// Get value arguments
			argAdmin := args[1]
			argMembers := strings.Split(args[2], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSignerGroup(
				clientCtx.GetFromAddress().String(),
				indexName,
				argAdmin,
				argMembers,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteSignerGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-signer-group [name]",
		Short: "Delete a signerGroup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSignerGroup(
				clientCtx.GetFromAddress().String(),
				indexName,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
