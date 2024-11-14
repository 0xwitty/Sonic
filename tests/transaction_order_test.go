package tests

import (
	"context"
	"github.com/Fantom-foundation/go-opera/tests/contracts/counter_event_emitter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

func TestTransactionOrder(t *testing.T) {
	const numTxs = 10
	net, err := StartIntegrationTestNet(t.TempDir())
	if err != nil {
		t.Fatalf("Failed to start the fake network: %v", err)
	}
	defer net.Stop()

	contract, _, err := DeployContract(net, counter_event_emitter.DeployCounterEventEmitter)
	if err != nil {
		t.Fatalf("failed to deploy contract; %v", err)
	}

	wantOrder := make([]common.Hash, 0, numTxs)
	for i := 0; i < numTxs; i++ {
		counter, err := contract.GetCount(nil)
		if err != nil {
			t.Fatalf("failed to get counter value; %v", err)
		}

		if counter.Cmp(new(big.Int).SetInt64(int64(i))) != 0 {
			t.Fatalf("unexpected counter value; expected %d, got %v", i, counter)
		}

		// Save transaction hashes in order at which the transactions are executed.
		_, err = net.Apply(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			tx, err := contract.Increment(opts)
			wantOrder = append(wantOrder, tx.Hash())
			return tx, err
		})
		if err != nil {
			t.Fatalf("failed to increment counter; %v", err)
		}
	}
	client, err := net.GetClient()
	if err != nil {
		t.Fatalf("cannot get client: %v", err)
	}
	lastBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatalf("cannot find last block number: %v", err)
	}
	if lastBlock == 0 {
		t.Fatal("0 blocks produced")
	}

	gotOrder := make(types.Transactions, 0, numTxs+8)

	for i := uint64(0); i <= lastBlock; i++ {
		blk, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(i))
		if err != nil {
			t.Fatalf("cannot get block number %v: %v", i, err)
		}
		blk.ReceiptHash()
		gotOrder = append(gotOrder, blk.Transactions()...)
	}

	for i, txHash := range wantOrder {
		receipt, err := client.TransactionReceipt(context.Background(), txHash) // first query synchronizes the execution
		if err != nil {
			t.Fatalf("failed to get receipt for tx %s; %v", txHash, err)
		}
		count, err := contract.ParseCount(*receipt.Logs[0])
		if err != nil {
			t.Fatalf("cannot parse count: %v", err)
		}
		if got, want := count.Count.Int64(), int64(i+1); got != want {
			t.Fatalf("incorrect transaction order, got count: %d, want count: %d", got, want)
		}
	}
}
