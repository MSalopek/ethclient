package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

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

// TODO: add proof verification by building a merkle tree from the proof and verifying the hash
// Fetch ethereum storage value & proofs for contract address and storage (key) at block height.
// In case the call fails, no data will be stored.
//
// NOTE: both storage value & proofs are inserted to DB using this method
func (k Keeper) EthCreateStorage(ctx sdk.Context, address, storageKey string, block int64) (types.Storage, error) {
	storageValRes, err := k.EthGetStorage(ctx, address, storageKey, block)
	if err != nil {
		return types.Storage{}, err
	}

	proofRes, err := k.EthGetProof(ctx, address, storageKey, block)
	if err != nil {
		return types.Storage{}, err
	}

	bs := strconv.Itoa(int(block))
	args := types.Args{
		Address: address,
		Storage: storageKey,
		Block:   bs,
	}
	s := types.Storage{
		Address:   address,
		Storage:   storageKey,
		Block:     bs,
		Value:     string(storageValRes),
		Args:      &args,
		Timestamp: time.Now().UnixNano(),
	}
	p := types.StorageProof{
		Proof: proofRes,
	}

	k.SetStorage(ctx, s, address, storageKey, bs)
	k.SetProof(ctx, address, storageKey, bs, p)

	return s, nil
}

// Fetch ethereum storage value for contract address and storage (key) at block height.
func (k Keeper) EthGetStorage(ctx context.Context, address, storage string, block int64) ([]byte, error) {
	addr := common.HexToAddress(address)
	storageHash := common.HexToHash(storage)                            // data at slot 0
	b := big.NewInt(block)                                              // block number
	storageVal, err := k.ethClient.StorageAt(ctx, addr, storageHash, b) // value is hexadecimal number
	if err != nil {
		return []byte{}, err
	}
	return storageVal, nil
}

// Fetch ethereum proof for storage value for contract address and storage (key) at block height.
func (k Keeper) EthGetProof(ctx context.Context, address, storage string, block int64) ([]byte, error) {
	padded := common.LeftPadBytes([]byte{0x0}, 32)

	hex := common.Bytes2Hex(padded)
	addr := common.HexToAddress(address)
	b := big.NewInt(block)
	proof, err := k.ethExtendedClient.GetProof(ctx, addr, []string{hex}, b)
	if err != nil {
		return nil, err
	}

	// simply marshal the proof to json for now
	return json.Marshal(proof)
}
