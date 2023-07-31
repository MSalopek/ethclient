package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"

	"ethclient/x/ethclient/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		ethClient  *ethclient.Client
		// ethereum RPC node client
		// used to fetch storage values and merkle proofs
		ethExtendedClient *gethclient.Client
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ethNodeURL string,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	k := Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}

	if ethNodeURL != "" {
		k.ethClient, k.ethExtendedClient = k.MustDialEthClient(ethNodeURL)
	}

	return &k
}

// Dial an ethereum RPC node client or panic if the client cannot be dialed
// Returns both an ethereum client and a geth client because geth client has extra functionality
func (k Keeper) MustDialEthClient(apiURL string) (*ethclient.Client, *gethclient.Client) {
	client, err := ethclient.Dial(apiURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial eth client: %v", err))
	}
	return client, gethclient.New(client.Client())
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Fetch ethereum storage value for contract address and storage (key) at block height.
func (k Keeper) EthGetStorage(ctx context.Context, address, storage string, block int) ([]byte, error) {
	addr := common.HexToAddress(address)
	storageHash := common.HexToHash(storage)                            // data at slot 0
	b := big.NewInt(int64(block))                                       // block number
	storageVal, err := k.ethClient.StorageAt(ctx, addr, storageHash, b) // value is hexadecimal number
	if err != nil {
		return []byte{}, err
	}
	return storageVal, nil
}

// Fetch ethereum proof for storage value for contract address and storage (key) at block height.
func (k Keeper) EthGetProof(ctx context.Context, address, storage string, block int) ([]byte, error) {
	padded := common.LeftPadBytes([]byte{0x0}, 32)

	hex := common.Bytes2Hex(padded)
	addr := common.HexToAddress(address)
	b := big.NewInt(int64(block))
	proof, err := k.ethExtendedClient.GetProof(ctx, addr, []string{hex}, b)
	if err != nil {
		return nil, err
	}

	// simply marshal the proof to json for now
	return json.Marshal(proof)
}
