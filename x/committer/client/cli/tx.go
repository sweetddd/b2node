package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/ethermint/x/committer/types"
)

// var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdSubmitProof())
	cmd.AddCommand(CmdTimeoutProposal())
	cmd.AddCommand(CmdBitcoinTx())
	cmd.AddCommand(CmdAddCommitter())
	cmd.AddCommand(CmdRemoveCommitter())

	return cmd
}

func CmdSubmitProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proof [proposal-id] [proof-hash] [state-root-hash] [start-index] [end-index]",
		Short: "Submit a proof proposal",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			proofHash := args[1]
			stateRootHash := args[2]

			startIndex, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}
			endIndex, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitProof(
				proposalID,
				clientCtx.GetFromAddress().String(),
				proofHash,
				stateRootHash,
				startIndex,
				endIndex,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdBitcoinTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bitcoin-tx [proposal-id] [bitcoin-tx-hash]",
		Short: "Submit the Bitcoin tx hash to update the proposal's status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			bitcoinTxHash := args[1]

			msg := types.NewMsgBitcoinTx(
				proposalID,
				clientCtx.GetFromAddress().String(),
				bitcoinTxHash,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdTimeoutProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "timeout-proposal [proposal-id]",
		Short: "Submit a tx to trigger proposal timeout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgTimeoutProposal(
				proposalID,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdAddCommitter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-committer [address]",
		Short: "Add a committer to the list of committers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddCommitter(clientCtx.GetFromAddress().String(), address)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveCommitter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-committer [address]",
		Short: "Remove a committer from the list of committers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveCommitter(clientCtx.GetFromAddress().String(), address)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
