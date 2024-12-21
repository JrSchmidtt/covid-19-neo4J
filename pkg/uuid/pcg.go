package uuid

import (
	"sync"
	"github.com/marcelomd/pcg"
)

type LockedPcg128 struct {
	mu  sync.Mutex
	pcg *pcg.Pcg128
}

func (p *LockedPcg128) Seed(a, b uint64) {
	p.mu.Lock()
	p.pcg.Seed(a, b)
	p.mu.Unlock()
}

func (p *LockedPcg128) Uint64() uint64 {
	p.mu.Lock()
	x := p.pcg.Uint64()
	p.mu.Unlock()
	return x
}

func (p *LockedPcg128) Uint128() (uint64, uint64) {
	p.mu.Lock()
	x, y := p.pcg.Uint128()
	p.mu.Unlock()
	return x, y
}

func NewLockedPcg128() LockedPcg128 {
	return LockedPcg128{pcg: pcg.NewDefaultPcg128()}
}
