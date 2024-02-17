package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/weBEE9/opt-auth-backend/mock"
	"github.com/weBEE9/opt-auth-backend/service"
	"go.uber.org/mock/gomock"
)

func TestOTPHandler_GenOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock.NewMockOTPService(ctrl)

	handler := NewOTPHandler(mock)

	t.Run("generate OTP", func(t *testing.T) {
		mock.EXPECT().
			GenOTP(gomock.Any(), "1234567890").
			Return("1234", nil).
			Times(1)

		reqBytes, err := json.Marshal(genOTPRequest{
			PhoneNumber: "1234567890",
		})
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/gen-otp", bytes.NewReader(reqBytes))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		handler.GenOTP().ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)

		var resp genOTPResponse
		json.Unmarshal(recorder.Body.Bytes(), &resp)

		require.Equal(t, http.StatusOK, resp.Code)
		require.NotEmpty(t, resp.OTP)
	})

	t.Run("OTP already exist", func(t *testing.T) {
		mock.EXPECT().
			GenOTP(gomock.Any(), "1234567890").
			Return("", service.ErrorOTPAlreadyExists).
			Times(1)

		reqBytes, err := json.Marshal(genOTPRequest{
			PhoneNumber: "1234567890",
		})
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/gen-otp", bytes.NewReader(reqBytes))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		handler.GenOTP().ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)

		var resp genOTPResponse
		json.Unmarshal(recorder.Body.Bytes(), &resp)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Equal(t, resp.Message, service.ErrorOTPAlreadyExists.Error())
	})
}

func TestOTPHandler_VerifyOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock.NewMockOTPService(ctrl)

	handler := NewOTPHandler(mock)

	t.Run("verify OTP", func(t *testing.T) {
		mock.EXPECT().
			VerifyOTP(gomock.Any(), "1234567890", "1234").
			Return(nil).
			Times(1)

		reqBody := verifyOTPRequest{
			PhoneNumber: "1234567890",
			OTP:         "1234",
		}
		reqBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/verify-otp", bytes.NewReader(reqBytes))
		recorder := httptest.NewRecorder()

		handler.VerifyOTP().ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)

		var resp verifyOTPResponse
		json.Unmarshal(recorder.Body.Bytes(), &resp)

		require.Equal(t, http.StatusOK, resp.Code)
		require.True(t, resp.Verified)
	})

	t.Run("OTP not found", func(t *testing.T) {
		mock.EXPECT().
			VerifyOTP(gomock.Any(), "1234567890", "1234").
			Return(service.ErrorOTPNotFound).
			Times(1)

		reqBody := verifyOTPRequest{
			PhoneNumber: "1234567890",
			OTP:         "1234",
		}
		reqBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/verify-otp", bytes.NewReader(reqBytes))
		recorder := httptest.NewRecorder()

		handler.VerifyOTP().ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)

		var resp verifyOTPResponse
		json.Unmarshal(recorder.Body.Bytes(), &resp)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.False(t, resp.Verified)
		require.Equal(t, resp.Message, service.ErrorOTPNotFound.Error())
	})

	t.Run("OTP mismatch", func(t *testing.T) {
		mock.EXPECT().
			VerifyOTP(gomock.Any(), "1234567890", "1234").
			Return(service.ErrorOTPMismatch).
			Times(1)

		reqBody := verifyOTPRequest{
			PhoneNumber: "1234567890",
			OTP:         "1234",
		}
		reqBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/verify-otp", bytes.NewReader(reqBytes))
		recorder := httptest.NewRecorder()

		handler.VerifyOTP().ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)

		var resp verifyOTPResponse
		json.Unmarshal(recorder.Body.Bytes(), &resp)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.False(t, resp.Verified)
		require.Equal(t, resp.Message, service.ErrorOTPMismatch.Error())
	})
}
