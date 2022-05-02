package main

import (
	"net/http"
	"os"

	"feedbacks/feedback"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)

	r := feedback.NewHttpServer(feedback.NewService(), logger)
	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", r))
}
