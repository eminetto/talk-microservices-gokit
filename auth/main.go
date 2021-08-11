package main

import (
	"auth/user"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8081", "caller", log.DefaultCaller)

	r := user.NewHttpServer(user.NewService(), logger)
	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
}
