package main

import (
	"auth/user"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)

	svc := user.NewLoggingMiddleware(logger, user.NewService())
	validateUserHandler := httptransport.NewServer(
		user.MakeValidateUserEndpoint(svc),
		user.DecodeValidateUserRequest,
		user.EncodeResponse,
	)

	http.Handle("/v1/auth", validateUserHandler)
	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", nil))
}
