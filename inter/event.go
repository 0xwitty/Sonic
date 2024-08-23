package inter

import (
	"crypto/sha256"

	"github.com/Fantom-foundation/lachesis-base/common/bigendian"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type EventI interface {
	dag.Event
	Version() uint8
	NetForkID() uint16
	CreationTime() Timestamp
	MedianTime() Timestamp
	PrevEpochHash() *hash.Hash
	Extra() []byte
	GasPowerLeft() GasPowerLeft
	GasPowerUsed() uint64

	HashToSign() hash.Hash
	Locator() EventLocator

	// Payload-related fields

	AnyTxs() bool
	AnyBridgeVotes() bool
	PayloadHash() hash.Hash
}

type EventLocator struct {
	BaseHash    hash.Hash
	NetForkID   uint16
	Epoch       idx.Epoch
	Seq         idx.Event
	Lamport     idx.Lamport
	Creator     idx.ValidatorID
	PayloadHash hash.Hash
}

type SignedEventLocator struct {
	Locator EventLocator
	Sig     Signature
}

type EventPayloadI interface {
	EventI
	Sig() Signature

	Txs() types.Transactions
	BridgeVotes() []BridgeVote
}

var emptyPayloadHash2 = CalcPayloadHash(&MutableEventPayload{extEventData: extEventData{version: 2}})

func EmptyPayloadHash(version uint8) hash.Hash {
	return emptyPayloadHash2
}

type baseEvent struct {
	dag.BaseEvent
}

type mutableBaseEvent struct {
	dag.MutableBaseEvent
}

type extEventData struct {
	version       uint8
	netForkID     uint16
	creationTime  Timestamp
	medianTime    Timestamp
	prevEpochHash *hash.Hash
	gasPowerLeft  GasPowerLeft
	gasPowerUsed  uint64
	extra         []byte

	anyTxs                bool
	anyBridgeVotes        bool
	payloadHash           hash.Hash
}

type sigData struct {
	sig Signature
}

type payloadData struct {
	txs                types.Transactions
	bridgeVotes        []BridgeVote
}

type Event struct {
	baseEvent
	extEventData

	// cache
	_baseHash    *hash.Hash
	_locatorHash *hash.Hash
}

type SignedEvent struct {
	Event
	sigData
}

type EventPayload struct {
	SignedEvent
	payloadData

	// cache
	_size int
}

type MutableEventPayload struct {
	mutableBaseEvent
	extEventData
	sigData
	payloadData
}

func (e *Event) HashToSign() hash.Hash {
	return *e._locatorHash
}

func asLocator(basehash hash.Hash, e EventI) EventLocator {
	return EventLocator{
		BaseHash:    basehash,
		NetForkID:   e.NetForkID(),
		Epoch:       e.Epoch(),
		Seq:         e.Seq(),
		Lamport:     e.Lamport(),
		Creator:     e.Creator(),
		PayloadHash: e.PayloadHash(),
	}
}

func (e *Event) Locator() EventLocator {
	return asLocator(*e._baseHash, e)
}

func (e *EventPayload) Size() int {
	return e._size
}

func (e *extEventData) Version() uint8 { return e.version }

func (e *extEventData) NetForkID() uint16 { return e.netForkID }

func (e *extEventData) CreationTime() Timestamp { return e.creationTime }

func (e *extEventData) MedianTime() Timestamp { return e.medianTime }

func (e *extEventData) PrevEpochHash() *hash.Hash { return e.prevEpochHash }

func (e *extEventData) Extra() []byte { return e.extra }

func (e *extEventData) PayloadHash() hash.Hash { return e.payloadHash }

func (e *extEventData) AnyTxs() bool { return e.anyTxs }

func (e *extEventData) AnyBridgeVotes() bool { return e.anyBridgeVotes }

func (e *extEventData) GasPowerLeft() GasPowerLeft { return e.gasPowerLeft }

func (e *extEventData) GasPowerUsed() uint64 { return e.gasPowerUsed }

func (e *sigData) Sig() Signature { return e.sig }

func (e *payloadData) Txs() types.Transactions { return e.txs }

func (e *payloadData) BridgeVotes() []BridgeVote { return e.bridgeVotes }

func CalcTxHash(txs types.Transactions) hash.Hash {
	return hash.Hash(types.DeriveSha(txs, trie.NewStackTrie(nil)))
}

func CalcReceiptsHash(receipts []*types.ReceiptForStorage) hash.Hash {
	hasher := sha256.New()
	_ = rlp.Encode(hasher, receipts)
	return hash.BytesToHash(hasher.Sum(nil))
}

func CalcBridgeVotesHash(bvs []BridgeVote) hash.Hash {
	hasher := sha256.New()
	_ = rlp.Encode(hasher, bvs)
	return hash.BytesToHash(hasher.Sum(nil))
}

func CalcPayloadHash(e EventPayloadI) hash.Hash {
	return hash.Of(
		CalcTxHash(e.Txs()).Bytes(),
		CalcBridgeVotesHash(e.BridgeVotes()).Bytes(),
	)
}

func (e *MutableEventPayload) SetVersion(v uint8) { e.version = v }

func (e *MutableEventPayload) SetNetForkID(v uint16) { e.netForkID = v }

func (e *MutableEventPayload) SetCreationTime(v Timestamp) { e.creationTime = v }

func (e *MutableEventPayload) SetMedianTime(v Timestamp) { e.medianTime = v }

func (e *MutableEventPayload) SetPrevEpochHash(v *hash.Hash) { e.prevEpochHash = v }

func (e *MutableEventPayload) SetExtra(v []byte) { e.extra = v }

func (e *MutableEventPayload) SetPayloadHash(v hash.Hash) { e.payloadHash = v }

func (e *MutableEventPayload) SetGasPowerLeft(v GasPowerLeft) { e.gasPowerLeft = v }

func (e *MutableEventPayload) SetGasPowerUsed(v uint64) { e.gasPowerUsed = v }

func (e *MutableEventPayload) SetSig(v Signature) { e.sig = v }

func (e *MutableEventPayload) SetTxs(v types.Transactions) {
	e.txs = v
	e.anyTxs = len(v) != 0
}

func (e *MutableEventPayload) SetBridgeVotes(v []BridgeVote) {
	e.bridgeVotes = v
	e.anyBridgeVotes = len(v) != 0
}

func calcEventID(h hash.Hash) (id [24]byte) {
	copy(id[:], h[:24])
	return id
}

func calcEventHashes(ser []byte, e EventI) (locator hash.Hash, base hash.Hash) {
	base = hash.Of(ser)
	if e.Version() < 1 {
		return base, base
	}
	return asLocator(base, e).HashToSign(), base
}

func (e *MutableEventPayload) calcHashes() (locator hash.Hash, base hash.Hash) {
	b, _ := e.immutable().Event.MarshalBinary()
	return calcEventHashes(b, e)
}

func (e *MutableEventPayload) size() int {
	b, err := e.immutable().MarshalBinary()
	if err != nil {
		panic("can't encode: " + err.Error())
	}
	return len(b)
}

func (e *MutableEventPayload) HashToSign() hash.Hash {
	h, _ := e.calcHashes()
	return h
}

func (e *MutableEventPayload) Locator() EventLocator {
	_, baseHash := e.calcHashes()
	return asLocator(baseHash, e)
}

func (e *MutableEventPayload) Size() int {
	return e.size()
}

func (e *MutableEventPayload) build(locatorHash hash.Hash, baseHash hash.Hash, size int) *EventPayload {
	return &EventPayload{
		SignedEvent: SignedEvent{
			Event: Event{
				baseEvent:    baseEvent{*e.MutableBaseEvent.Build(calcEventID(locatorHash))},
				extEventData: e.extEventData,
				_baseHash:    &baseHash,
				_locatorHash: &locatorHash,
			},
			sigData: e.sigData,
		},
		payloadData: e.payloadData,
		_size:       size,
	}
}

func (e *MutableEventPayload) immutable() *EventPayload {
	return e.build(hash.Hash{}, hash.Hash{}, 0)
}

func (e *MutableEventPayload) Build() *EventPayload {
	locatorHash, baseHash := e.calcHashes()
	payloadSer, _ := e.immutable().MarshalBinary()
	return e.build(locatorHash, baseHash, len(payloadSer))
}

func (l EventLocator) HashToSign() hash.Hash {
	return hash.Of(l.BaseHash.Bytes(), bigendian.Uint16ToBytes(l.NetForkID), l.Epoch.Bytes(), l.Seq.Bytes(), l.Lamport.Bytes(), l.Creator.Bytes(), l.PayloadHash.Bytes())
}

func (l EventLocator) ID() hash.Event {
	h := l.HashToSign()
	copy(h[0:4], l.Epoch.Bytes())
	copy(h[4:8], l.Lamport.Bytes())
	return hash.Event(h)
}
