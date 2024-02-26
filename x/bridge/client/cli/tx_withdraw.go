package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdCreateWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-withdraw [tx-id] [tx-hash-list] [encoded-data]",
		Short: "Create a new withdraw",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxID := args[0]

			// Get value arguments
			argTxHashList := strings.Split(args[1], listSeparator)
			argEncodedData := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxID,
				argTxHashList,
				argEncodedData,
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

func CmdUpdateWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-withdraw [tx-id] [status]",
		Short: "Update a withdraw",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxID := args[0]

			// Get value arguments
			argStatus, err := types.WithdrawStatusFromString(args[1])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxID,
				argStatus,
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

func CmdSignWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-withdraw [tx-id] [signature]",
		Short: "Sign a withdraw",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxID := args[0]

			// Get value arguments
			argSignature := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSignWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxID,
				argSignature,
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

func CmdDeleteWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-withdraw [tx-id]",
		Short: "Delete a withdraw",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexTxID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxID,
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
