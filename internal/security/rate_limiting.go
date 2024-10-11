/*
Package security provides a rate limiting mechanism to prevent abuse of the service.
*/

package security

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter implements request rate limiting functionality.
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.Mutex
	limit    int           // Maximum number of requests allowed in the interval
	interval time.Duration // Time window for request limiting
}

// visitor holds the visitor's request count and last seen timestamp.
type visitor struct {
	lastSeen time.Time
	requests int
}

// NewRateLimiter creates a new RateLimiter instance.
func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*visitor),
		limit:    limit,
		interval: interval,
	}
}

// Allow checks if a request from a client/IP is allowed based on rate limiting rules.
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		rl.visitors[ip] = &visitor{
			lastSeen: time.Now(),
			requests: 1,
		}
		return true
	}

	// Check the time since the last request
	elapsed := time.Since(v.lastSeen)
	if elapsed > rl.interval {
		v.lastSeen = time.Now()
		v.requests = 1
		return true
	}

	// Increment request count
	v.requests++

	// Deny if request limit exceeded
	return v.requests <= rl.limit
}

// Cleanup removes visitors who haven't made requests in the specified interval.
func (rl *RateLimiter) Cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	for ip, v := range rl.visitors {
		if time.Since(v.lastSeen) > rl.interval {
			delete(rl.visitors, ip)
		}
	}
}

// Middleware applies rate limiting to incoming HTTP requests.
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if !rl.Allow(ip) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
