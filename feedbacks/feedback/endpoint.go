package feedback

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
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
		fmt.Println(f)
		_, err := svc.Store(f)
		if err != nil {
			return storeResponse{"", err.Error()}, err
		}

		return storeResponse{f.ID.String(), ""}, err
	}
}

func decodeStoreRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("Unauthorized") //@todo criar tipo de erro para tratar o unautorized
	}

	var request storeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	tokenEndpoint, err := makeValidateTokenEndpoint("http://localhost:8081/v1/validate-token")
	if err != nil {
		return nil, err
	}
	tokenResponse, err := tokenEndpoint(ctx, validateTokenRequest{Token: tokenString})
	if err != nil {
		return "", err
	}

	resp := tokenResponse.(validateTokenResponse)
	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}
	request.Email = resp.Email
	return request, nil
}

type validateTokenRequest struct {
	Token string `json:"token"`
}

type validateTokenResponse struct {
	Email string `json:"email,omitempty"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeValidateTokenEndpoint(serviceUrl string) (endpoint.Endpoint, error) {
	u, err := url.Parse(serviceUrl)
	if err != nil {
		return nil, err
	}
	return httptransport.NewClient(
		"POST",
		u,
		encodeValidateTokenRequest,
		decodeValidateTokenResponse,
	).Endpoint(), nil
}

func encodeValidateTokenRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeValidateTokenResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response validateTokenResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
