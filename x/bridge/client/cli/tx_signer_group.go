package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateSignerGroup() *cobra.Command { //nolint:dupl
	cmd := &cobra.Command{
		Use:   "create-signer-group [name] [admin] [threshold] [members]",
		Short: "Create a new signerGroup",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]

			// Get value arguments
			argAdmin := args[1]
			argThreshold, err := cast.ToUint32E(args[2])
			if err != nil {
				return err
			}
			argMembers := strings.Split(args[3], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignerGroup(
				clientCtx.GetFromAddress().String(),
				indexName,
				argAdmin,
				argThreshold,
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

func CmdUpdateSignerGroup() *cobra.Command { //nolint:dupl
	cmd := &cobra.Command{
		Use:   "update-signer-group [name] [admin] [threshold] [members]",
		Short: "Update a signerGroup",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]

			// Get value arguments
			argAdmin := args[1]
			argThreshold, err := cast.ToUint32E(args[2])
			if err != nil {
				return err
			}
			argMembers := strings.Split(args[3], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSignerGroup(
				clientCtx.GetFromAddress().String(),
				indexName,
				argAdmin,
				argThreshold,
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
