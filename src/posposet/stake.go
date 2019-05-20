package posposet

import (
	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/state"
)

// stakeCounter is for PoS balances accumulator.
type stakeCounter struct {
	balances *state.DB
	amount   uint64
	goal     uint64
}

func (s *stakeCounter) Count(node hash.Peer) {
	if s.IsGoalAchieved() {
		return // no sense to count further
	}
	s.amount += s.balances.VoteBalance(node)
}

func (s *stakeCounter) IsGoalAchieved() bool {
	return s.amount > s.goal
}

/*
 * Poset's methods:
 */

// GetBalanceOf returns stake balance of peer.
func (p *Poset) GetBalanceOf(addr hash.Peer) uint64 {
	f := p.frame(p.state.LastFinishedFrameN, false)
	db := p.store.StateDB(f.Balances)
	return db.VoteBalance(addr)
}

// GetStakeOf returns stake of peer as fraction from one.
func (p *Poset) GetStakeOf(addr hash.Peer) float64 {
	balance := p.GetBalanceOf(addr)

	return float64(balance) / float64(p.state.TotalCap)
}

func (p *Poset) newStakeCounter(frame *Frame, goal uint64) *stakeCounter {
	db := p.store.StateDB(frame.Balances)

	return &stakeCounter{
		balances: db,
		amount:   0,
		goal:     goal,
	}
}

func (p *Poset) hasMajority(frame *Frame, roots EventsByPeer) bool {
	stake := p.newStakeCounter(frame,
		p.state.TotalCap*2/3)
	for node := range roots {
		stake.Count(node)
	}
	return stake.IsGoalAchieved()
}

func (p *Poset) hasTrust(frame *Frame, roots EventsByPeer) bool {
	stake := p.newStakeCounter(frame,
		p.state.TotalCap*1/3)
	for node := range roots {
		stake.Count(node)
	}
	return stake.IsGoalAchieved()
}