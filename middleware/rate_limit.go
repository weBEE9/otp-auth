package middleware

import (
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	mu         sync.Mutex
	limiter    map[string]*rate.Limiter
	rate       rate.Limit
	bucketSize int
}

func NewRateLimiter(limitRate, bucketSize int) *RateLimiter {
	limiRate := rate.Every(time.Duration(limitRate) * time.Minute)

	return &RateLimiter{
		limiter:    make(map[string]*rate.Limiter),
		rate:       limiRate,
		bucketSize: bucketSize,
	}
}

func (r *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := log.New(os.Stdout, "[RateLimiter]", log.LstdFlags)

		sourceAddr, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			logger.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		logger.Printf("Request from addr: '%s'", sourceAddr)

		if !r.allow(sourceAddr) {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, req)
	})
}

func (r *RateLimiter) allow(value string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exist := r.limiter[value]
	if !exist {
		limiter = rate.NewLimiter(r.rate, r.bucketSize)
		r.limiter[value] = limiter
	}

	return limiter.Allow()
}
