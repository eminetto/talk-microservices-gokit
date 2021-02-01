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
	Email    string
}

type storeResponse struct {
	Uuid string `json:"uuid,omitempty"`
	Err  string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeStoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(storeRequest)
		v := Vote{
			ID:       uuid.New(),
			Email:    req.Email,
			TalkName: req.TalkName,
			Score:    req.Score,
		}
		_, err := svc.Store(ctx, v)
		if err != nil {
			return storeResponse{"", err.Error()}, err
		}

		return storeResponse{v.ID.String(), ""}, err
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
