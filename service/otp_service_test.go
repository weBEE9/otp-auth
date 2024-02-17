package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/weBEE9/opt-auth-backend/mock"
	"github.com/weBEE9/opt-auth-backend/repository"
	"go.uber.org/mock/gomock"
)

func TestOTPServiceRedis_GenOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock.NewMockOTPRepository(ctrl)

	service := NewOTPService(mock)

	ctx := context.Background()
	phoneNumber := "1234567890"

	t.Run("generate OTP", func(t *testing.T) {
		mock.EXPECT().
			GenOTP(gomock.Any(), phoneNumber).
			Return("1234", nil).
			Times(1)

		otp, err := service.GenOTP(ctx, phoneNumber)
		require.NoError(t, err)
		require.Equal(t, "1234", otp)
	})

	t.Run("OTP already exist", func(t *testing.T) {
		mock.EXPECT().
			GenOTP(gomock.Any(), phoneNumber).
			Return("", repository.ErrorOTPAlreadyExists).
			Times(1)

		otp, err := service.GenOTP(ctx, phoneNumber)
		require.Error(t, err)
		require.ErrorIs(t, err, repository.ErrorOTPAlreadyExists)
		require.Empty(t, otp)
	})
}

func TestOTPServiceRedis_VerifyOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock.NewMockOTPRepository(ctrl)

	service := NewOTPService(mock)

	ctx := context.Background()
	phoneNumber := "1234567890"
	otp := "1234"

	t.Run("OTP not found", func(t *testing.T) {
		mock.EXPECT().
			VerifyOTP(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(repository.ErrorOTPNotFound).
			Times(1)

		err := service.VerifyOTP(ctx, phoneNumber, otp)
		require.ErrorIs(t, err, repository.ErrorOTPNotFound)
	})

	t.Run("OTP mismatched", func(t *testing.T) {
		mock.EXPECT().
			VerifyOTP(gomock.Any(), phoneNumber, otp).
			Return(repository.ErrorOTPMismatch).
			Times(1)

		err := service.VerifyOTP(ctx, phoneNumber, otp)
		require.ErrorIs(t, err, repository.ErrorOTPMismatch)
	})

	t.Run("OTP verified", func(t *testing.T) {
		mock.EXPECT().
			VerifyOTP(gomock.Any(), phoneNumber, otp).
			Return(nil).
			Times(1)

		err := service.VerifyOTP(ctx, phoneNumber, otp)
		require.NoError(t, err)
	})
}
