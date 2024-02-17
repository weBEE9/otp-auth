package service

import (
	"context"

	"github.com/weBEE9/opt-auth-backend/repository"
)

type optServiceRedis struct {
	Repository repository.OTPRepository
}

func NewOTPService(repository repository.OTPRepository) OTPService {
	return &optServiceRedis{Repository: repository}
}

func (s *optServiceRedis) GenOTP(ctx context.Context, phoneNumber string) (string, error) {
	return s.Repository.GenOTP(ctx, phoneNumber)
}

func (s *optServiceRedis) VerifyOTP(ctx context.Context, phoneNumber, otp string) error {
	return s.Repository.VerifyOTP(ctx, phoneNumber, otp)
}
