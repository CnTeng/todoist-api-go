package ws

import (
	"math/rand/v2"
	"sync"
	"time"
)

type backoff struct {
	initInterval time.Duration
	maxInterval  time.Duration
	multiplier   float64

	mu       sync.Mutex
	interval time.Duration
}

func newBackoff(init, maxBackoff time.Duration) *backoff {
	return &backoff{
		initInterval: init,
		maxInterval:  maxBackoff,
		multiplier:   2.0,
	}
}

func (b *backoff) Duration() time.Duration {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.interval == 0 {
		b.interval = b.initInterval
	}

	delay := time.Duration(rand.Int64N(int64(b.interval)))

	if float64(b.interval)*b.multiplier >= float64(b.maxInterval) {
		b.interval = b.maxInterval
	} else {
		b.interval = time.Duration(float64(b.interval) * b.multiplier)
	}

	return delay
}

func (b *backoff) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.interval = b.initInterval
}
