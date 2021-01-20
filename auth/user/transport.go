package user

import (
	"auth/security"
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

func MakeValidateUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateUserRequest)
		err := svc.ValidateUser(req.Email, req.Password)
		if err != nil {
			return validateUserResponse{"", err.Error()}, err
		}
		token, err := security.NewToken(req.Email)
		if err != nil {
			return validateUserResponse{"", err.Error()}, err
		}
		return validateUserResponse{token, ""}, err
	}
}

func DecodeValidateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request validateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
