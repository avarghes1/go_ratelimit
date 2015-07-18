// Rate Limit.
//
// Package can be used to throttle access to resources.
// An example would be to limit requests to an external api
// which is rate limited.
//
// Example:
//          r := ratelimit.New(5, time.Second, 10*time.Second)
//          for i := 0; i < 100; i++ {
//              go func() {
//                  advance := r.Get()
//                  if advance {
//                      time.Sleep(time.Millisecond * 100)
//                      go r.Put()
//                  }
//              }()
//          }
//
// @author: avarghese
package ratelimit

import (
	"sync"
	"time"
)

type (
	Rate struct {
		s  *sem          // Semaphore
		l  sync.Mutex    // Lock before editing
		t  time.Duration // Rate Limit Duration
		to time.Duration // time out
	}
	sem struct {
		b time.Time // Record when semaphore is borrowed.
		n chan bool // Channel which holds semaphore
	}
)

// Get a Rate object.
//
// Usage:
//      r := ratelimit.New(5, time.Second, 10*time.Second)
//
func New(n int, t, to time.Duration) *Rate {
	c := make(chan bool, n)
	for i := 0; i < n; i++ {
		c <- true
	}
	return &Rate{s: &sem{n: c}, t: t, to: to}
}

// Get the go ahead to access a resource
func (r *Rate) Get() bool {
	var out bool
	select {
	case out = <-r.s.n:
		// lock before setting borrowed time.
		r.l.Lock()
		r.s.b = time.Now()
		// Unlock
		r.l.Unlock()
	case <-time.After(r.to):
		// timeout
		out = false
	}
	return out
}

// Done with resource access. Return Semaphore.
func (r *Rate) Put() {
	r.l.Lock()
	// Wait till duration is done before
	// returning semaphore
	for time.Now().Sub(r.s.b) < r.t {
	}
	r.l.Unlock()
	r.s.n <- true
}
