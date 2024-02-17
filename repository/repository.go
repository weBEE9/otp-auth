package repository

//go:generate mockgen -destination=../mock/repository_mock.go -package=mock github.com/weBEE9/opt-auth-backend/repository OTPRepository

import (
	"context"

	"github.com/weBEE9/opt-auth-backend/model"
)

// OTPRepository is an interface for process OTP(One Time Password) related operations
type OTPRepository interface {
	GenOTP(ctx context.Context, phoneNumber string) (string, error)
	VerifyOTP(ctx context.Context, phoneNumber, otp string) error
}

type UserRepository interface {
	GetUser(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
}
