package user

import (
	"context"
	"encoding/json"

	"net/http"

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
		token, err := svc.ValidateUser(req.Email, req.Password)
		if err != nil {
			return validateUserResponse{"", err.Error()}, err
		}
		return validateUserResponse{token, ""}, err
	}
}

func decodeValidateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request validateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
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
		email, err := svc.ValidateToken(req.Token)
		if err != nil {
			return validateTokenResponse{"", err.Error()}, err
		}
		return validateTokenResponse{email, ""}, err
	}
}

func decodeValidateTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request validateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}