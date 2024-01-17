package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deposit [tx-hash] [from] [to] [coin-type] [value] [data] [status]",
		Short: "Create a new deposit",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxHash := args[0]

			// Get value arguments
			argFrom := args[1]
			argTo := args[2]
			argCoinType := args[3]
			argValue, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argData := args[5]
			argStatus := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDeposit(
				clientCtx.GetFromAddress().String(),
				indexTxHash,
				argFrom,
				argTo,
				argCoinType,
				argValue,
				argData,
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

func CmdUpdateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-deposit [tx-hash] [from] [to] [coin-type] [value] [data] [status]",
		Short: "Update a deposit",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxHash := args[0]

			// Get value arguments
			argFrom := args[1]
			argTo := args[2]
			argCoinType := args[3]
			argValue, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argData := args[5]
			argStatus := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDeposit(
				clientCtx.GetFromAddress().String(),
				indexTxHash,
				argFrom,
				argTo,
				argCoinType,
				argValue,
				argData,
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

func CmdDeleteDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-deposit [tx-hash]",
		Short: "Delete a deposit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexTxHash := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDeposit(
				clientCtx.GetFromAddress().String(),
				indexTxHash,
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
