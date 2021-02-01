package feedback

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type storeRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Email string
}

type storeResponse struct {
	Uuid string `json:"uuid,omitempty"`
	Err  string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeStoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(storeRequest)
		f := Feedback{
			ID:    uuid.New(),
			Email: req.Email,
			Title: req.Title,
			Body:  req.Body,
		}
		_, err := svc.Store(ctx, f)
		if err != nil {
			return storeResponse{"", err.Error()}, err
		}

		return storeResponse{f.ID.String(), ""}, err
	}
}
