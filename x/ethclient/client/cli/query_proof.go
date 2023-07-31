package cli

import (
	"strconv"

	"ethclient/x/ethclient/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof [address] [storage] [block]",
		Short: "Query proof",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqStorage := args[1]
			reqBlock := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProofRequest{

				Address: reqAddress,
				Storage: reqStorage,
				Block:   reqBlock,
			}

			res, err := queryClient.Proof(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
