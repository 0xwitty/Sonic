package tests

import (
	"context"
	"github.com/Fantom-foundation/go-opera/tests/contracts/counter_event_emitter"
	"github.com/Fantom-foundation/go-opera/utils/signers/gsignercache"
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

	client, err := net.GetClient()
	if err != nil {
		t.Fatalf("cannot get client: %v", err)
	}
	defer client.Close()

	opts, err := net.GetTransactOptions(&net.validator)
	if err != nil {
		t.Fatalf("cannot get transact options: %v", err)
	}
	tx, err := contract.Increment(opts)
	if err != nil {
		t.Fatalf("failed to increment counter; %v", err)
	}

	signer := gsignercache.Wrap(types.NewCancunSigner(new(big.Int).SetUint64(4003)))
	wantOrder := make([]common.Hash, 0, numTxs)
	for i := uint64(0); i < numTxs; i++ {
		signedTx, err := types.SignTx(types.NewTx(&types.DynamicFeeTx{
			ChainID:    new(big.Int).Set(tx.ChainId()),
			Nonce:      i + 2,
			GasTipCap:  new(big.Int).Set(tx.GasTipCap()),
			GasFeeCap:  new(big.Int).Set(tx.GasFeeCap()),
			Gas:        tx.Gas(),
			To:         tx.To(),
			Value:      new(big.Int).Set(tx.Value()),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
		}), signer, net.validator.PrivateKey)
		if err != nil {
			t.Fatalf("cannot sign tx: %v", err)
		}

		err = client.SendTransaction(context.Background(), signedTx) // < from this point on, tx is processed asynchronously
		if err != nil {
			t.Fatalf("failed to send transaction; %v", err)
		}

		wantOrder = append(wantOrder, signedTx.Hash())
	}

	receipts := make([]*types.Receipt, 0, numTxs)
	for _, txHash := range wantOrder {
		receipt, err := client.TransactionReceipt(context.Background(), txHash) // first query synchronizes the execution
		if err != nil {
			t.Fatalf("failed to get receipt for tx %s; %v", txHash, err)
		}
		receipts = append(receipts, receipt)
	}
}
