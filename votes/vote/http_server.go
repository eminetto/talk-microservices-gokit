package vote

import (
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(svc Service, logger log.Logger) {
	storeHandler := httptransport.NewServer(
		makeStoreEndpoint(svc),
		decodeStoreRequest,
		encodeResponse,
	)

	http.Handle("/v1/vote", storeHandler)
	logger.Log("msg", "HTTP", "addr", "8083")
	logger.Log("err", http.ListenAndServe(":8083", nil))
}
