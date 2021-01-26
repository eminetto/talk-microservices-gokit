package main

import (
	"os"

	"feedbacks/feedback"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8082", "caller", log.DefaultCaller)

	svc := feedback.NewLoggingMiddleware(logger, feedback.NewService())
	feedback.NewHttpServer(svc, logger)
}
