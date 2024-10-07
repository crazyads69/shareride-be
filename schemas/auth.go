package schemas

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

// Define Payload struct

type Payload struct {
	ID          uuid.UUID `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiredAt   time.Time `json:"expired_at"`
}

// Define GenerateOTPRequest struct
type GenerateOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=9,max=11"`
}

// Define GenerateOTPResponse struct
type GenerateOTPResponse struct {
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
}

// Define struct to first register a user
type RegisterUserRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	FullName    string `json:"full_name" binding:"required,min=3,max=256"`
}

// Define struct to first register a user response
type RegisterUserResponse struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	FullName    string    `json:"full_name" binding:"required,min=3,max=256"`
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
}

// Define struct to verify the OTP
type VerifyRegisterOTPRequest struct {
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	OTP         string    `json:"otp" binding:"required,numeric,min=6,max=6"`
}

// Define struct to verify the OTP response
type VerifyRegisterOTPResponse struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	FullName    string    `json:"full_name" binding:"required,min=3,max=256"`
	IsActivated bool      `json:"is_activated" binding:"required"`
}

// Define struct for VerifyCCCDRequest
// Two images of the CCCD (front and back) as form-data
type VerifyCCCDRequest struct {
	FrontImage  *multipart.FileHeader `form:"front_image" binding:"required"`
	BackImage   *multipart.FileHeader `form:"back_image" binding:"required"`
	UserID      string                `form:"user_id" binding:"required,uuid"`
	PhoneNumber string                `form:"phone_number" binding:"required,numeric,min=9,max=11"`
}

// Define struct for VerifyCCCDResponse
type VerifyCCCDResponse struct {
	User         UserResponse `json:"user" binding:"required"`
	AccessToken  string       `json:"access_token" binding:"required"`
	RefreshToken string       `json:"refresh_token" binding:"required"`
}

// Define struct for RegisterOAuthRequest
type RegisterOAuthRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	FullName    string `json:"full_name" binding:"required,min=3,max=256"`
	Email       string `json:"email" binding:"required,email, min=3,max=256"`
}

// Define struct for RegisterOAuthResponse
type RegisterOAuthResponse struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	FullName    string    `json:"full_name" binding:"required,min=3,max=256"`
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	Email       string    `json:"email" binding:"required,email, min=3,max=256"`
}

// Define struct for VerifyLoginOTPRequest
type VerifyLoginOTPRequest struct {
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
	OTP         string    `json:"otp" binding:"required,numeric,min=6,max=6"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
}

// Define struct for VerifyLoginOTPResponse
type VerifyLoginOTPResponse struct {
	User         UserResponse `json:"user" binding:"required"`
	AccessToken  string       `json:"access_token" binding:"required"`
	RefreshToken string       `json:"refresh_token" binding:"required"`
}

// Define struct for LoginWithOAuthRequest
type LoginWithOAuthRequest struct {
	Email string `json:"email" binding:"required,email,min=3,max=256"`
}

// Define struct for LoginWithOAuthResponse
type LoginWithOAuthResponse struct {
	Email       string    `json:"email" binding:"required,email,min=3,max=256"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required,numeric,min=9,max=11"`
}
