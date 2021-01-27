package user

import (
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(svc Service, logger log.Logger) {
	validateUserHandler := httptransport.NewServer(
		makeValidateUserEndpoint(svc),
		decodeValidateUserRequest,
		encodeResponse,
	)

	validateTokenHandler := httptransport.NewServer(
		makeValidateTokenEndpoint(svc),
		decodeValidateTokenRequest,
		encodeResponse,
	)

	http.Handle("/v1/auth", validateUserHandler)
	http.Handle("/v1/validate-token", validateTokenHandler)
	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", nil))
}
