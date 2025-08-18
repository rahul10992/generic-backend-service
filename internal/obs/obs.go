package obs

import "sync"

type Readiness struct {
	isReady bool
	mu      sync.RWMutex
}

func NewReadiness() *Readiness {
	return &Readiness{isReady: false}
}

func (r *Readiness) SetReady() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.isReady = true
}

func (r *Readiness) SetNotReady() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.isReady = false
}

func (r *Readiness) IsReady() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ready := r.isReady
	return ready
}
