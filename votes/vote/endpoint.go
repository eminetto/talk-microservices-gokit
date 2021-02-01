package vote

import (
	"context"

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
