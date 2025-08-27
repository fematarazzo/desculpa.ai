package server

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	clients map[string]int
	mu      sync.Mutex
}

var limiter = &rateLimiter{
	clients: make(map[string]int),
}

const maxRequestsPerHour = 10

func init() {
	go func() {
		for {
			time.Sleep(time.Hour)
			limiter.mu.Lock()
			limiter.clients = make(map[string]int)
			limiter.mu.Unlock()
		}
	}()
}

func allowRequest(ip string) bool {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	count := limiter.clients[ip]
	if count >= maxRequestsPerHour {
		return false
	}
	limiter.clients[ip] = count + 1
	return true
}

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if !allowRequest(ip) {
			http.Error(w, "Você atingiu o limite de 10 desculpas por hora. Volte mais tarde!", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
