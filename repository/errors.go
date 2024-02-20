package repository

import "errors"

var (
	ErrorOTPMismatch      = errors.New("otp mismatch with this phone number")
	ErrorOTPNotFound      = errors.New("otp not found with this phone number")
	ErrorOTPAlreadyExists = errors.New("otp already exists for this phone number")
	ErrorOTPWillNotExpire = errors.New("otp will not expire")
)
