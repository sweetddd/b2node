package cli

import (
	"fmt"
	"time"

	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/ethermint/x/committer/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

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
	cmd.AddCommand(CmdBatchProof())
	cmd.AddCommand(CmdBatchProof())
	cmd.AddCommand(CmdTapRoot())

	return cmd 
}

func CmdBatchProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-proof [proposal-id] [proof-hash] [state-root-hash] [start-index] [end-index]",
		Short: "Broadcast BatchProof transaction",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			
			proposalId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			proofHash := args[1]
			stateRootHash := args[2]

			startIndex, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}
			endIndex, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewBatchProofMsg(
				uint64(proposalId),
				clientCtx.GetFromAddress().String(), 
				proofHash,
				stateRootHash,
				uint64(startIndex),
				uint64(endIndex),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}


func CmdTapRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tap-root [proposal-id] [bitcoin-tx-hash]",
		Short: "Broadcast TapRoot transaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			
			proposalId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			bitcoinTxHash := args[1]

			msg := types.NewTapRootMsg(
				uint64(proposalId),
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
		Short: "Broadcast TimeoutProposal transaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			
			proposalId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewTimeoutProposalMsg(
				uint64(proposalId),
				clientCtx.GetFromAddress().String(), 
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}