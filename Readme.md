# ethclient

ethclient module is an ethereum lightclient stores ethereum data on chain and can be used to retrieve and validate the stored data.

**ethclient** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

# Ethereum storage
Ethereum uses Merkle Patricia tries for storage. Merkle tries allow state verification without having to build the entire storage (database) from scratch. Instead, we can use Merkle proofs (a propery of every Merkle trie) to prove assertions about some state. Under the hood, ethereum storage can be thought of as a key-value store where every contract has read/write access to its own state.

Storage objects can be retrieved using `eth_getStorageAt` calls to eth RPC clients. This method simply returns a value stored under a key hash.
* the method is useful for smart contract accounts
* the method can be called for externally owned acounts, but all storage values are zero
* it takes `address` , `[storageeKeys]` and `blockParamerer` as arguments

It is the caller's responsibility to correctly index the KV store, which complicates accessing non-primitive types (arrays, maps etc.)


# Ethereum storage proofs

Ethereum storage is the `data availability` layer of the ethereum blockchain.
When dealing with ethereum storage off-chain, application developers need to deal with the concept of `trust` - can the data being stored for off-chain use be trusted, for how long and under which conditions is the trust assumption invalidated (broken).

Instead of relying on trust, the validity of storage values obrained from `eth_getStorageAt` can be verified using `eth_getProof`. Using this method, Merkle proofs for some storage state can be obtained and used to verify data validity in a trustless manner.


# Storing ethereum data on cosmos-sdk based chains
All operations for storing ethereum data on a different chain must be initiated by a user agent that is able to pay transaction fees.

Ethereum data can be used for various purposes, so methods to store, access and verify states must be made available to the user agent.

The methods are:
**Transactions**
* `CreateStorage`
    * args: `account`, `height` (block height), `storageKey`
    * e.g. `ethclientd q ethclient storage create-storage`
    * **users must calculate the storage keys indexes on their own - check this [guide](https://docs.infura.io/networks/ethereum/json-rpc-methods/eth_getstorageat)**

**Queries**
* `QueryStorage`
    * args: `account`, `height` (block height ), `storageKey`
    * return storage value with metadata (arguments used to retrieve it from ethereum)
    * e.g. `ethclientd q ethclient storage`
    * **users must calculate the storage indexes on their own - check this [guide](https://docs.infura.io/networks/ethereum/json-rpc-methods/eth_getstorageat)**


* `QueryProof`
    * args `account`, `height`, `storageKey`
    * return proofs for the storage value so they can be independently verified
    * e.g. `ethclientd q ethclient proof`
    * **users must calculate the storage indexes on their own - check this [guide](https://docs.infura.io/networks/ethereum/json-rpc-methods/eth_getstorageat)**


All methods have a corresponding gRPC method implemented and can be accessed using the CLI commands or RPC calls.

All methods for interacting with an ethereum RPC node are in `/x/ethclient/keeper/eth_client`:

```golang
func (k Keeper) MustDialEthClient(apiURL string) (*ethclient.Client, *gethclient.Client) {..}
func (k Keeper) EthCreateStorage(ctx sdk.Context, address, storageKey string, block int64) (types.Storage, error) {...}
func (k Keeper) EthGetStorage(ctx context.Context, address, storage string, block int64) ([]byte, error) {...}
func (k Keeper) EthGetProof(ctx context.Context, address, storage string, block int64) ([]byte, error) {...}
```


# How to use

### Install ignite cli
To install the latest version of `ignite cli`, execute the following command on your machine:

```
curl https://get.ignite.com/username/ethclient@latest! | sudo bash
```

The latest version at the time of writing is:
```
➜  config ignite version
Ignite CLI version:		     v0.27.1
Ignite CLI build date:		2023-06-13T13:42:09Z
Ignite CLI source hash:		4acd1f185afb6d8d1a837e54f04c091121cfae01
Ignite CLI config version:	v1
Cosmos SDK version:		v0.47.3
```

### Configuration

The local development setup uses additional config flags in `app.toml`. The `ethereum-rpc-url` flag should be set in your `config.yml` or in your `app.toml` if you are not running the chain using `ignire chain serve`.

The blockchain will not start if an ethereum RPC node URL is not available.

```yaml
version: 1
accounts:
- name: alice
  coins:
  - 20000token
  - 200000000stake
- name: bob
  coins:
  - 10000token
  - 100000000stake
client:
  openapi:
    path: docs/static/openapi.yml
faucet:
  name: bob
  coins:
  - 5token
  - 100000stake
validators:
- name: alice
  bonded: 100000000stake
  config:
    moniker: "eth-enabled-node"
  app:
    ethereum-rpc-url: "<eth_rpc_node_api_url>"
build:
  main: cmd/ethclientd
```

## Starting the chain

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.


# Future improvements

At the moment the implementation simply fetches data from ethereum and stores it to state so it can be retrieved. 

However, this implementation is incomplete at best and incorrect at worst. Blockchain data should be verified, and for that the data should be verified before inclusion in the blockchain (otherwise we are fully trusting the RPC node returning the results).

Some hints on verification:
* [minimal imlementation of a Merkle-Patricia trie](https://github.com/zhangchiqing/merkle-patricia-trie)
* a short overview of verifying proofs by building a trie and checking against it:
    * https://flow.com/engineering-blogs/verify-ethereum-smart-contract-state-with-proof
    * https://flow.com/engineering-blogs/ethereum-merkle-patricia-trie-explained

TODOs:
1. Actually verify proofs obtained from `eth_getProof`
2. improve testing
3. improve CRUD

## Add More Tx

Add a Tx allowing users to pay a fee to re-validate proofs.

* `ValidateData`
    * this method checks stored state against data on ethereum
    * validation is performed by fetching proofs for storage and re-validating the storage value
    * if the "old" proof and the "new" are different, the data is `invalid`
    * if the data is not valid, the state object will be marked `invalid` and will not be returned in queries any more

# Sources
* https://blog.infura.io/post/how-to-use-ethereum-proofs-2
* https://ethereum.org/pl/developers/tutorials/merkle-proofs-for-offline-data-integrity/
* https://ethereum.org/fil/developers/docs/data-availability/
* https://www.starknet.io/en/posts/developers/what-are-storage-proofs-and-how-can-they-improve-oracles
* https://geth.ethereum.org/docs/developers/dapp-developer/native
* https://docs.infura.io/networks/ethereum/json-rpc-methods/eth_getstorageat
* https://docs.infura.io/networks/ethereum/json-rpc-methods/eth_getproof
* https://github.com/zhangchiqing/merkle-patricia-trie
* https://flow.com/engineering-blogs/ethereum-merkle-patricia-trie-explained
* https://flow.com/engineering-blogs/verify-ethereum-smart-contract-state-with-proof
