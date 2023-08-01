package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"ethclient/x/ethclient/types"
)

func CmdShowStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "storage",
		Short: "Query ethereum storage for account, storage (key) and block combination",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqAddress := args[0]
			reqStorage := args[1]
			reqBlock := args[2]

			parsedBlock, err := strconv.Atoi(reqBlock)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStorageRequest{
				Address: reqAddress,
				Storage: reqStorage,
				Block:   int64(parsedBlock),
			}

			res, err := queryClient.QueryStorage(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
