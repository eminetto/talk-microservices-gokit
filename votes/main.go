package main

import (
	"os"

	"votes/vote"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8083", "caller", log.DefaultCaller)

	svc := vote.NewLoggingMiddleware(logger, vote.NewService())
	vote.NewHttpServer(svc, logger)
}
