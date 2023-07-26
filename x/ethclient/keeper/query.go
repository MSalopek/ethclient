package keeper

import (
	"ethclient/x/ethclient/types"
)

var _ types.QueryServer = Keeper{}
