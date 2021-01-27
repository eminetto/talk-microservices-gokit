package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger log.Logger) {
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

	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
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
