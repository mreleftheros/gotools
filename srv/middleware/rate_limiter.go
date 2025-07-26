package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/mreleftheros/gotools/srv/json"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func RateLimit(rps float64, burst int, isEnabled bool, next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, c := range clients {
				if time.Since(c.lastSeen) > time.Minute*3 {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isEnabled {
			ip := realip.FromRequest(r)
			// ip, _, err := net.SplitHostPort(r.RemoteAddr)
			// if err != nil {
			// 	app.serverErrorResponse(w, r, err)
			// 	return
			// }

			mu.Lock()

			if _, ok := clients[ip]; !ok {
				clients[ip] = &client{limiter: rate.NewLimiter(rate.Limit(rps), burst)}
			}

			clients[ip].lastSeen = time.Now()

			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				json.Write(w, r, http.StatusTooManyRequests, json.NewErrorResponse("rate limit exceeded", nil))
				return
			}

			mu.Unlock()
		}
		next.ServeHTTP(w, r)
	})
}
