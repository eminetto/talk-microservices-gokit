package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type validateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type validateUserResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeValidateUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateUserRequest)
		token, err := svc.ValidateUser(ctx, req.Email, req.Password)
		if err != nil {
			return validateUserResponse{"", err.Error()}, err
		}
		return validateUserResponse{token, ""}, err
	}
}

type validateTokenRequest struct {
	Token string `json:"token"`
}

type validateTokenResponse struct {
	Email string `json:"email,omitempty"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeValidateTokenEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateTokenRequest)
		email, err := svc.ValidateToken(ctx, req.Token)
		if err != nil {
			return validateTokenResponse{"", err.Error()}, err
		}
		return validateTokenResponse{email, ""}, err
	}
}
