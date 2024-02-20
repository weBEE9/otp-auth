package service

import (
	"context"
	"errors"
	"time"

	"github.com/weBEE9/opt-auth-backend/repository"
)

type optServiceRedis struct {
	Repository repository.OTPRepository
}

func NewOTPService(repository repository.OTPRepository) OTPService {
	return &optServiceRedis{Repository: repository}
}

func (s *optServiceRedis) GenOTP(ctx context.Context, phoneNumber string) (string, error) {
	otp, err := s.Repository.GenOTP(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, repository.ErrorOTPAlreadyExists) {
			return "", repository.ErrorOTPAlreadyExists
		}

		return "", err
	}

	return otp, nil
}

func (s *optServiceRedis) VerifyOTP(ctx context.Context, phoneNumber, otp string) error {
	err := s.Repository.VerifyOTP(ctx, phoneNumber, otp)
	if err != nil {
		if errors.Is(err, repository.ErrorOTPNotFound) {
			return repository.ErrorOTPNotFound
		}

		if errors.Is(err, repository.ErrorOTPMismatch) {
			return repository.ErrorOTPMismatch
		}

		return err
	}

	return nil
}

func (s *optServiceRedis) TTL(ctx context.Context, phoneNumber string) (time.Duration, error) {
	ttl, err := s.Repository.TTL(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, repository.ErrorOTPNotFound) {
			return 0, ErrorOTPNotFound
		}

		return 0, err
	}

	return ttl, nil
}
