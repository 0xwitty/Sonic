package tests

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestGasFee_CanSendLowPricedTransactions(t *testing.T) {
	dataDir := t.TempDir()
	net, err := StartIntegrationTestNet(dataDir)
	if err != nil {
		t.Fatalf("Failed to start the fake network: %v", err)
	}
	defer net.Stop()

	// Get a test account for this test.
	account := NewAccount()
	accountAddress := account.Address()
	if err := net.EndowAccount(accountAddress, 1e18); err != nil {
		t.Fatalf("Failed to endow the account with tokens: %v", err)
	}

	client, err := net.GetClient()
	if err != nil {
		t.Fatalf("Failed to connect to the integration test network: %v", err)
	}
	defer client.Close()

	chainId := big.NewInt(int64(opera.FakeNetworkID))

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("Failed to get gas price: %v", err)
	}
	t.Logf("Suggested gas price: %v\n", price)

	// sample different gas prices
	for i := range 2 {
		//107_625_105_000

		cap := price
		//cap := big.NewInt(100_000_000_000 + int64(i)*1_000_000_000) // 100 Gwei + i Gwei
		price = new(big.Int).Sub(price, big.NewInt(10_000_000_000)) // decrease by 1 Gwei

		// Type 2 -- Dynamic Fee Transactions (London)
		transaction, err := types.SignTx(types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainId,
			Gas:       21000,
			GasFeeCap: cap,
			GasTipCap: cap, // big.NewInt(200_000), // minimal tip
			To:        &common.Address{},
			Value:     big.NewInt(1000),
			Nonce:     uint64(i),
		}), types.NewLondonSigner(chainId), account.PrivateKey)
		if err != nil {
			t.Fatalf("Failed to sign transaction: %v", err)
		}

		receipt, err := net.Run(transaction)
		if err != nil {
			t.Errorf("Failed to send transaction with cap %d: %v", cap, err)
		} else {
			t.Logf("Cap: %d, Effective price: %d\n", cap, receipt.EffectiveGasPrice)
		}

	}

	t.Fail()
}

func TestUpdate(t *testing.T) {
	encoded := "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000267b2245636f6e6f6d79223a7b224d696e4761735072696365223a353832353738303935347d7d0000000000000000000000000000000000000000000000000000"
	encoded = "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000267b2245636f6e6f6d79223a7b224d696e4761735072696365223a363037393432323431377d7d0000000000000000000000000000000000000000000000000000"
	data := make([]byte, hex.DecodedLen(len(encoded)))

	len, err := hex.Decode(data, []byte(encoded))
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}
	t.Logf("Decoded %d bytes\n", len)

	fmt.Printf(string(data))

	t.Fail()
}
