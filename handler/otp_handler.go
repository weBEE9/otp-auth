package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/weBEE9/opt-auth-backend/service"
)

type OTPHandler struct {
	Service service.OTPService
}

func NewOTPHandler(service service.OTPService) *OTPHandler {
	return &OTPHandler{Service: service}
}

type genOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type genOTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (h *OTPHandler) GenOTP() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			req, err := decode[genOTPRequest](r)
			if err != nil {
				encode[genOTPResponse](w, r, http.StatusInternalServerError, genOTPResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to decode request: " + err.Error(),
				})
				return
			}

			otp, err := h.Service.GenOTP(r.Context(), req.PhoneNumber)
			if err != nil {
				code := http.StatusInternalServerError
				if errors.Is(err, service.ErrorOTPAlreadyExists) {
					code = http.StatusBadRequest
				}

				encode[genOTPResponse](w, r, code, genOTPResponse{
					Code:    code,
					Message: err.Error(),
				})

				return
			}

			if err := sendOTPToPhoneNumber(req.PhoneNumber, otp); err != nil {
				encode[genOTPResponse](w, r, http.StatusInternalServerError, genOTPResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to send OTP: " + err.Error(),
				})

				return
			}

			if err := encode[genOTPResponse](
				w,
				r,
				http.StatusOK,
				genOTPResponse{
					Code:    http.StatusOK,
					Message: fmt.Sprintf("OTP has sent to: [%s]", req.PhoneNumber),
				},
			); err != nil {
				encode[genOTPResponse](w, r, http.StatusInternalServerError, genOTPResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to encode response: " + err.Error(),
				})

				return
			}
		},
	)
}

type verifyOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

type verifyOTPResponse struct {
	Code     int    `json:"code"`
	Verified bool   `json:"verified,omitempty"`
	Message  string `json:"message,omitempty"`
}

func (h *OTPHandler) VerifyOTP() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			req, err := decode[verifyOTPRequest](r)
			if err != nil {
				encode[verifyOTPResponse](w, r, http.StatusInternalServerError, verifyOTPResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to decode request: " + err.Error(),
				})

				return
			}

			if err := h.Service.VerifyOTP(r.Context(), req.PhoneNumber, req.OTP); err != nil {
				code := http.StatusInternalServerError
				if errors.Is(err, service.ErrorOTPMismatch) || errors.Is(err, service.ErrorOTPNotFound) {
					code = http.StatusBadRequest
				}

				encode[verifyOTPResponse](w, r, code, verifyOTPResponse{
					Code:    code,
					Message: err.Error(),
				})

				return
			}

			if err := encode[verifyOTPResponse](
				w,
				r,
				http.StatusOK,
				verifyOTPResponse{
					Code:     http.StatusOK,
					Verified: true,
				},
			); err != nil {
				encode[genOTPResponse](w, r, http.StatusInternalServerError, genOTPResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to encode response: " + err.Error(),
				})

				return
			}
		},
	)
}

type getTTLRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type getTTLResponse struct {
	Code    int    `json:"code"`
	TTL     string `json:"ttl,omitempty"`
	Message string `json:"message,omitempty"`
}

func (h *OTPHandler) GetTTL() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			req, err := decode[getTTLRequest](r)
			if err != nil {
				encode[getTTLResponse](w, r, http.StatusInternalServerError, getTTLResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to decode request: " + err.Error(),
				})

				return
			}

			ttl, err := h.Service.TTL(r.Context(), req.PhoneNumber)
			if err != nil {
				code := http.StatusInternalServerError
				if errors.Is(err, service.ErrorOTPMismatch) || errors.Is(err, service.ErrorOTPNotFound) {
					code = http.StatusBadRequest
				}

				encode[getTTLResponse](w, r, code, getTTLResponse{
					Code:    code,
					Message: err.Error(),
				})

				return
			}

			if err := encode[getTTLResponse](
				w,
				r,
				http.StatusOK,
				getTTLResponse{
					Code: http.StatusOK,
					TTL:  ttl.String(),
				},
			); err != nil {
				encode[getTTLResponse](w, r, http.StatusInternalServerError, getTTLResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to encode response: " + err.Error(),
				})

				return
			}
		},
	)
}

func sendOTPToPhoneNumber(phoneNumber, otp string) error {
	// send OTP to phoneNumber
	return nil
}
