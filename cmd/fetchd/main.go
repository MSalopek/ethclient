package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

var apiURL = flag.String("api-url", "", "api-url of an ethereum RPC NODE")

// Storage gives read access to contract internal storage
// Storage keys cannot be iterated over in a straightforward way without knowing the contract
// The keys (and types of keys) to check need to be known beforehand and the
//
// After getting storage key values, the merkle proof can be fetched using gethclient.GetProof
func main() {
	flag.Parse()
	if *apiURL == "" {
		log.Fatal("api-url is required")
	}

	client, err := ethclient.Dial(*apiURL)
	if err != nil {
		log.Fatal(err)
	}
	geth := gethclient.New(client.Client())

	fmt.Println("we have a connection")
	ctx := context.Background()
	addr := common.HexToAddress("0xB9951B43802dCF3ef5b14567cb17adF367ed1c0F")
	zeroHash := common.HexToHash("0x0") // data at slot 0
	blockBigInt := big.NewInt(int64(17813192))
	storageVal, err := client.StorageAt(ctx, addr, zeroHash, blockBigInt) // value is hexadecimal number
	if err != nil {
		fmt.Println("HAVE ERR", err)
		return
	}
	fmt.Println("RESULT", storageVal)
	padded := common.LeftPadBytes([]byte{0x0}, 32)
	hex := common.Bytes2Hex(padded)
	proof, err := geth.GetProof(ctx, addr, []string{hex}, blockBigInt)
	if err != nil {
		fmt.Println("HAVE ERR in PROOF", err)
		return
	}
	buf, err := json.MarshalIndent(proof, "", "  ")
	if err != nil {
		fmt.Println("marshalling err", err)
		return
	}

	fmt.Printf("PROOF\n %+v", string(buf))

	// TODO: figure out how to verify using the proofs and building a trie

	// the storage hash and the proof is the data to be verified
	// storageHash := proof.StorageHash
	// storageProof := proof.StorageProof[0]
	// // encode the key-value pair
	// key := common.LeftPadBytes([]byte(storageProof.Key), 32)
	// value, err := rlp.EncodeToBytes(storageProof.Value)
	// if err != nil {
	// 	fmt.Println("proof err", err)
	// 	return
	// }
	// build a trie with the nodes in the proof
	// 	proofTrie := NewProofDB()
	// 	for _, node := range storageProof.Proof {
	// 		proofTrie.Put(crypto.Keccak256(node), node)
	// 	}
	// 	// verify the proof
	// 	verified, err := VerifyProof(
	// 		storageHash.Bytes(), crypto.Keccak256(key), proofTrie)
	// 	require.True(t, bytes.Equal(verified, value), fmt.Sprintf("%x != %x", verified, value))
	// }
}
