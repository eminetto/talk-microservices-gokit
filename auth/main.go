package main

import (
	"auth/user"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := user.NewService()

	validateUserHandler := httptransport.NewServer(
		user.MakeValidateUserEndpoint(svc),
		user.DecodeValidateUserRequest,
		user.EncodeResponse,
	)

	http.Handle("/v1/auth", validateUserHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
