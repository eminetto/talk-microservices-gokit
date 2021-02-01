package main

import (
	"auth/user"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)

	svc := user.NewLoggingMiddleware(logger, user.NewService())
	user.NewHttpServer(svc, logger)
}
