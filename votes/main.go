package main

import (
	"net/http"
	"os"

	"votes/vote"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8083", "caller", log.DefaultCaller)

	svc := vote.NewLoggingMiddleware(logger, vote.NewService())
	r := vote.NewHttpServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", "8083")
	logger.Log("err", http.ListenAndServe(":8083", r))
}
