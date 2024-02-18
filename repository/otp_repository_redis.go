package repository

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	otpCacheKeyPrefix = "phone_number"
)

var (
	fiveMinutes = 5 * time.Minute
)

type otpRepositoryRedis struct {
	mu          *sync.Mutex
	RedisClient *redis.Client
}

func getCahceKey(phoneNumber string) string {
	return fmt.Sprintf("%s:%s", otpCacheKeyPrefix, phoneNumber)
}

func NewRedisOtpRepository(redisClient *redis.Client) OTPRepository {
	return &otpRepositoryRedis{
		mu:          new(sync.Mutex),
		RedisClient: redisClient,
	}
}

func (r *otpRepositoryRedis) GenOTP(ctx context.Context, phoneNumber string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	cacheKey := getCahceKey(phoneNumber)

	_, err := r.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		fourDigitOTP := fmt.Sprint(rand.Intn(9999))

		err := r.RedisClient.
			Set(ctx, cacheKey, fourDigitOTP, fiveMinutes).
			Err()
		if err != nil {
			return "", err
		}

		return fourDigitOTP, nil
	}

	if err != nil {
		return "", err
	}

	return "", ErrorOTPAlreadyExists
}

func (r *otpRepositoryRedis) VerifyOTP(ctx context.Context, phoneNumber, otp string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	cacheKey := getCahceKey(phoneNumber)

	storedOTP, err := r.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return ErrorOTPNotFound
	}

	if err != nil {
		return err
	}

	if storedOTP != otp {
		return ErrorOTPMismatch
	}

	return nil
}
