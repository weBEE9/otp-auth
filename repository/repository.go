package repository

import (
	"context"

	"github.com/weBEE9/opt-auth-backend/model"
)

// OTPRepository is an interface for process OTP(One Time Password) related operations
type OTPRepository interface {
	GenOTP(ctx context.Context, phoneNumber string) (string, error)
	VerifyOTP(ctx context.Context, phoneNumber, otp string) error
}
