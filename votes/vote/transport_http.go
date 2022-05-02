package vote

import (
	"context"
	"encoding/json"
	"github.com/eminetto/talk-microservices-gokit/pkg/middleware"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger kitlog.Logger) *mux.Router {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerFinalizer(newServerFinalizer(logger)),
	}
	storeHandler := kithttp.NewServer(
		makeStoreEndpoint(svc),
		decodeStoreRequest,
		encodeResponse,
		options...,
	)
	r := mux.NewRouter()
	r.Use(middleware.IsAuthenticatedMiddleware)
	r.Methods("POST").Path("/v1/vote").Handler(storeHandler)
	return r
}

func newServerFinalizer(logger kitlog.Logger) kithttp.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		logger.Log("status",code, "path", r.RequestURI, "method", r.Method)
	}
}

func decodeStoreRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request storeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	request.Email = r.Header.Get("email")
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
