package inter

import (
	"fmt"
	"github.com/Fantom-foundation/go-lachesis/src/common"

	"github.com/Fantom-foundation/go-lachesis/src/crypto"
	"github.com/Fantom-foundation/go-lachesis/src/cryptoaddr"
	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/src/inter/wire"
)

type EventHeaderData struct {
	Version uint32

	Epoch idx.SuperFrame
	Seq   idx.Event

	Frame  idx.Frame
	IsRoot bool

	Creator hash.Peer // TODO common.Address

	Parents hash.Events

	GasLeft uint64
	GasUsed uint64

	Lamport     idx.Lamport
	ClaimedTime Timestamp
	MedianTime  Timestamp

	TxHash common.Hash

	Extra []byte
}

type EventHeader struct {
	EventHeaderData

	Sig []byte

	hash hash.Event // cache for .Hash()
}

func (e EventHeader) HashToSign() hash.Hash {
	// TODO
	/*hasher := sha3.New256()
	err := rlp.Encode(hasher, []interface{}{
		"Fantom signed event header", e.EventHeaderData)
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
	// return 32 bytes hash
	return hash.Hash(hasher.Sum(nil))*/
	return hash.Hash{}
}

func (e EventHeader) SelfParent() *hash.Event {
	if e.Seq <= 1 || len(e.Parents) == 0 {
		return nil
	}
	return &e.Parents[0]
}

func (e EventHeader) SelfParentEqualTo(hash hash.Event) bool {
	if e.SelfParent() == nil {
		return false
	}
	return *e.SelfParent() == hash
}

func (e EventHeader) GenesisHash() *hash.Event {
	if e.Seq > 1 || len(e.Parents) == 0 {
		return nil
	}
	return &e.Parents[0]
}

type Event struct {
	EventHeader
	InternalTransactions []*InternalTransaction
	ExternalTransactions ExtTxns
}

// SignBy signs event by private key.
func (e *Event) SignBy(priv *crypto.PrivateKey) error {
	eventHash := e.Hash()

	sig, err := priv.Sign(eventHash.Bytes())
	if err != nil {
		return err
	}

	e.Sig = sig
	return nil
}

// Verify sign event by public key.
func (e *Event) VerifySignature() bool {
	return cryptoaddr.VerifySignature(e.Creator, hash.Hash(e.Hash()), e.Sig)
}

// Hash calcs hash of event.
func (e *Event) Hash() hash.Event {
	if e.hash.IsZero() {
		// TODO
		/*hasher := sha3.New256()
		err := rlp.Encode(hasher, e.EventHeaderData)
		if err != nil {
			panic("can't encode: " + err.Error())
		}
		// return  epoch | lamport | 24 bytes hash
		e.hash = hash.NewEvent(e.Epoch, e.Lamport, hasher.Sum(nil))*/
	}
	return e.hash
}

// FindInternalTxn find transaction in event's internal transactions list.
// TODO: use map
func (e *Event) FindInternalTxn(idx hash.Transaction) *InternalTransaction {
	for _, txn := range e.InternalTransactions {
		if TransactionHashOf(e.Creator, txn.Nonce) == idx {
			return txn
		}
	}
	return nil
}

// String returns string representation.
func (e *Event) String() string {
	return fmt.Sprintf("Event{%s, %s, t=%d}", e.Hash().String(), e.Parents.String(), e.Lamport)
}

// TODO erase
// ToWire converts to proto.Message.
func (e *Event) ToWire() (*wire.Event, *wire.Event_ExtTxnsValue) {
	return nil, nil
}

// TODO erase
// WireToEvent converts from wire.
func WireToEvent(w *wire.Event) *Event {
	return nil
}

/*
 * Utils:
 */

// FakeFuzzingEvents generates random independent events for test purpose.
func FakeFuzzingEvents() (res []*Event) {
	creators := []hash.Peer{
		{},
		hash.FakePeer(),
		hash.FakePeer(),
		hash.FakePeer(),
	}
	parents := []hash.Events{
		hash.FakeEvents(1),
		hash.FakeEvents(2),
		hash.FakeEvents(8),
	}
	extTxns := [][][]byte{
		nil,
		[][]byte{
			[]byte("fake external transaction 1"),
			[]byte("fake external transaction 2"),
		},
	}
	i := 0
	for c := 0; c < len(creators); c++ {
		for p := 0; p < len(parents); p++ {
			e := &Event{
				EventHeader: EventHeader{
					EventHeaderData: EventHeaderData{
						Seq:     idx.Event(p),
						Creator: creators[c],
						Parents: parents[p],
					},
				},
				InternalTransactions: []*InternalTransaction{
					{
						Amount:   999,
						Receiver: creators[c],
					},
				},
				ExternalTransactions: ExtTxns{
					Value: extTxns[i%len(extTxns)],
				},
			}

			res = append(res, e)
			i++
		}
	}
	return
}
