package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}
	validateUserHandler := httptransport.NewServer(
		makeValidateUserEndpoint(svc),
		decodeValidateUserRequest,
		encodeResponse,
		options...,
	)

	validateTokenHandler := httptransport.NewServer(
		makeValidateTokenEndpoint(svc),
		decodeValidateTokenRequest,
		encodeResponse,
		options...,
	)
	r := mux.NewRouter()
	r.Methods("POST").Path("/v1/auth").Handler(validateUserHandler)
	r.Methods("POST").Path("/v1/validate-token").Handler(validateTokenHandler)
	return r
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrInvalidUser:
		return http.StatusNotFound
	case ErrInvalidToken:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func decodeValidateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request validateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
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
