package main

import (
	"net/http"
	"os"

	"votes/vote"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8083", "caller", log.DefaultCaller)

	svc := vote.NewLoggingMiddleware(logger, vote.NewService())
	storeHandler := httptransport.NewServer(
		vote.MakeStoreEndpoint(svc),
		vote.DecodeStoreRequest,
		vote.EncodeResponse,
	)

	http.Handle("/v1/vote", storeHandler)
	logger.Log("msg", "HTTP", "addr", "8083")
	logger.Log("err", http.ListenAndServe(":8083", nil))
}
