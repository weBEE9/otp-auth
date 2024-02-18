package repository

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/require"
)

func TestOTPRepositoryRedis_GenOTP(t *testing.T) {
	db, mock := redismock.NewClientMock()

	ctx := context.Background()

	phoneNumber := "1234567890"
	cacheKey := getCahceKey(phoneNumber)
	t.Run("generate OTP", func(t *testing.T) {
		mock.ExpectGet(cacheKey).RedisNil()
		mock.Regexp().ExpectSet(cacheKey, "[0-9]{4}", fiveMinutes).SetVal("OK")

		repo := NewRedisOtpRepository(db)
		otp, err := repo.GenOTP(ctx, phoneNumber)
		require.NoError(t, err)
		require.NotEmpty(t, otp)
	})

	t.Run("generate OTP with same phone number before key exipre, should get error", func(t *testing.T) {
		mock.Regexp().ExpectGet(cacheKey).SetVal("[0-9]{4}")

		repo := NewRedisOtpRepository(db)
		otp, err := repo.GenOTP(ctx, phoneNumber)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrorOTPAlreadyExists)
		require.Empty(t, otp)
	})
}

func TestOTPRepositoryRedis_VerifyOTP(t *testing.T) {
	db, mock := redismock.NewClientMock()

	ctx := context.Background()

	repo := NewRedisOtpRepository(db)

	t.Run("OTP with not found with phone number", func(t *testing.T) {
		phoneNumber := "1234567890"
		otp := "1234"
		cacheKey := getCahceKey(phoneNumber)

		mock.ExpectGet(cacheKey).RedisNil()

		err := repo.VerifyOTP(ctx, "1234567890", otp)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrorOTPNotFound)
	})

	t.Run("OTP mismatch with phone number", func(t *testing.T) {
		phoneNumber := "1234567890"
		otp := "1234"
		cacheKey := getCahceKey(phoneNumber)
		// set key first
		db.Set(ctx, cacheKey, otp, fiveMinutes)

		mock.ExpectGet(cacheKey).SetVal(otp)

		err := repo.VerifyOTP(ctx, "1234567890", "9876")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrorOTPMismatch)
	})

	t.Run("OTP verifyied with phone number", func(t *testing.T) {
		phoneNumber := "1234567890"
		otp := "1234"
		cacheKey := getCahceKey(phoneNumber)
		// set key first
		db.Set(ctx, cacheKey, otp, fiveMinutes)

		mock.ExpectGet(cacheKey).SetVal(otp)

		err := repo.VerifyOTP(ctx, "1234567890", otp)
		require.NoError(t, err)
	})
}
