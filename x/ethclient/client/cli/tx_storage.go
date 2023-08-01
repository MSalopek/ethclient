package cli

import (
	"ethclient/x/ethclient/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-storage",
		Short: "Create storage",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqStorage := args[1]
			reqBlock := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			parsedBlock, err := strconv.Atoi(reqBlock)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateStorage(clientCtx.GetFromAddress().String(), reqAddress, reqStorage, int64(parsedBlock))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
