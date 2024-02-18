package service

//go:generate mockgen -destination=../mock/service_mock.go -package=mock github.com/weBEE9/opt-auth-backend/service OTPService

import "context"

type OTPService interface {
	GenOTP(ctx context.Context, phoneNumber string) (string, error)
	VerifyOTP(ctx context.Context, phoneNumber, otp string) error
}
