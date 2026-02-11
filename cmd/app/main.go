package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Use 'localhost' for local testing, or the k8s service name when deployed
	redisAddr := os.Getenv("REDIS_HOST")
	if redisAddr == "" {
		redisAddr = "localhost:6379" 
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Basic route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Incr(ctx, "visits").Result()
		if err != nil {
			fmt.Fprintf(w, "Welcome! (Redis not connected: %v)", err)
			return
		}
		fmt.Fprintf(w, "AIOps Platform - Visitor Count: %d", val)
	})

	// Health check for Kubernetes liveness/readiness probes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	log.Println("Master Node App starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
