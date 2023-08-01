package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"

	"ethclient/x/ethclient/types"
)

// ETHStorageHandler defines the expected methods for fetching ethereum storage values and proofs
type ETHStorageHandler interface {
	EthGetStorage(ctx context.Context, address, storage string, block int64) ([]byte, error)
	EthGetProof(ctx context.Context, address, storage string, block int64) ([]byte, error)
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

	meta := types.MetaData{
		Address:   address,
		Storage:   storageKey,
		Block:     block,
		Timestamp: time.Now().UnixNano(),
	}
	s := types.Storage{
		Value: storageValRes,
		Meta:  &meta,
	}
	p := types.StorageProof{
		Proof: proofRes,
	}

	k.SetStorage(ctx, s, address, storageKey, block)
	k.SetProof(ctx, address, storageKey, block, p)

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
