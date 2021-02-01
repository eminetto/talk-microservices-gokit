package feedback

import (
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger log.Logger) {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}
	storeHandler := httptransport.NewServer(
		makeStoreEndpoint(svc),
		decodeStoreRequest,
		encodeResponse,
		options...,
	)
	r := mux.NewRouter()
	r.Methods("POST").Path("/v1/feedback").Handler(storeHandler)
	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", r))
}
