package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strings"
)

func CmdCreateWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-withdraw [index] [tx-hash] [from] [to] [coin-type] [value] [data] [status] [signatures]",
		Short: "Create a new withdraw",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argTxHash := args[1]
			argFrom := args[2]
			argTo := args[3]
			argCoinType := args[4]
			argValue, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argData := args[6]
			argStatus := args[7]
			argSignatures := strings.Split(args[8], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateWithdraw(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argTxHash,
				argFrom,
				argTo,
				argCoinType,
				argValue,
				argData,
				argStatus,
				argSignatures,
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
		Use:   "update-withdraw [index] [tx-hash] [from] [to] [coin-type] [value] [data] [status] [signatures]",
		Short: "Update a withdraw",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argTxHash := args[1]
			argFrom := args[2]
			argTo := args[3]
			argCoinType := args[4]
			argValue, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argData := args[6]
			argStatus := args[7]
			argSignatures := strings.Split(args[8], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateWithdraw(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argTxHash,
				argFrom,
				argTo,
				argCoinType,
				argValue,
				argData,
				argStatus,
				argSignatures,
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
		Use:   "delete-withdraw [index]",
		Short: "Delete a withdraw",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteWithdraw(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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
