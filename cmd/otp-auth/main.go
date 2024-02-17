package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/weBEE9/opt-auth-backend/config"
	"github.com/weBEE9/opt-auth-backend/handler"
	"github.com/weBEE9/opt-auth-backend/repository"
	"github.com/weBEE9/opt-auth-backend/service"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg, err := config.Environ()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: cfg.Redis.Addr(),
		},
	)

	otpRepo := repository.NewRedisOtpRepository(redisClient)
	otpService := service.NewOTPService(otpRepo)
	otpHandler := handler.NewOTPHandler(otpService)

	mux := http.NewServeMux()
	mux.Handle("POST /api/v1/otp", otpHandler.GenOTP())
	mux.Handle("POST /api/v1/otp/verify", otpHandler.VerifyOTP())

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Printf("server listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		pid := syscall.Getpid()
		log.Printf("PID %d. Received SIGINT. Shutting down...", pid)

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("PID %d. Error shutting down: %s", pid, err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
