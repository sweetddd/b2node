package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateWithdraw() *cobra.Command { //nolint:dupl
	cmd := &cobra.Command{
		Use:   "create-withdraw [tx-hash] [from] [to] [coin-type] [value] [data]",
		Short: "Create a new withdraw",
		Args:  cobra.ExactArgs(6),
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

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxHash,
				argFrom,
				argTo,
				argCoinType,
				argValue,
				argData,
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
		Use:   "update-withdraw [tx-hash] [status]",
		Short: "Update a withdraw",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxHash := args[0]

			// Get value arguments
			argStatus := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxHash,
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
		Use:   "sign-withdraw [tx-hash] [signature]",
		Short: "Sign a withdraw",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTxHash := args[0]

			// Get value arguments
			argSignature := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSignWithdraw(
				clientCtx.GetFromAddress().String(),
				indexTxHash,
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
		Use:   "delete-withdraw [tx-hash]",
		Short: "Delete a withdraw",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexTxHash := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteWithdraw(
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
