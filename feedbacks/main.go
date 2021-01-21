package main

import (
	"net/http"
	"os"

	"feedbacks/feedback"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8082", "caller", log.DefaultCaller)

	svc := feedback.NewLoggingMiddleware(logger, feedback.NewService())
	storeHandler := httptransport.NewServer(
		feedback.MakeStoreEndpoint(svc),
		feedback.DecodeStoreRequest,
		feedback.EncodeResponse,
	)

	http.Handle("/v1/feedback", storeHandler)
	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", nil))
}
