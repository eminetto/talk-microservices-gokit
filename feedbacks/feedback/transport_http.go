package feedback

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/eminetto/talk-microservices-gokit/pkg/middleware"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger log.Logger) *mux.Router {
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
	r.Use(middleware.IsAuthenticatedMiddleware)
	r.Methods("POST").Path("/v1/feedback").Handler(storeHandler)

	return r
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
