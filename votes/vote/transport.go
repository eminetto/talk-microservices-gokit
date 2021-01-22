package vote

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type storeRequest struct {
	TalkName string `json:"talk_name"`
	Score    int    `json:"score,string"`
}

type storeResponse struct {
	Uuid string `json:"uuid,omitempty"`
	Err  string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func MakeStoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(storeRequest)
		f := Vote{
			ID:       uuid.New(),
			Email:    "xxx@xx.com", //@todo: buscar do outro serviço
			TalkName: req.TalkName,
			Score:    req.Score,
		}
		_, err := svc.Store(f)
		if err != nil {
			return storeResponse{"", err.Error()}, err
		}

		return storeResponse{f.ID.String(), ""}, err
	}
}

func DecodeStoreRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request storeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
