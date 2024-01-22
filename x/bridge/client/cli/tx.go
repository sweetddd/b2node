package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/ethermint/x/bridge/types"
)

const (
	listSeparator = ","
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

	cmd.AddCommand(CmdCreateSignerGroup())
	cmd.AddCommand(CmdUpdateSignerGroup())
	cmd.AddCommand(CmdDeleteSignerGroup())
	cmd.AddCommand(CmdCreateCallerGroup())
	cmd.AddCommand(CmdUpdateCallerGroup())
	cmd.AddCommand(CmdDeleteCallerGroup())
	cmd.AddCommand(CmdCreateDeposit())
	cmd.AddCommand(CmdUpdateDeposit())
	cmd.AddCommand(CmdDeleteDeposit())
	cmd.AddCommand(CmdCreateWithdraw())
	cmd.AddCommand(CmdUpdateWithdraw())
	cmd.AddCommand(CmdSignWithdraw())
	cmd.AddCommand(CmdDeleteWithdraw())
	// this line is used by starport scaffolding # 1

	return cmd
}
