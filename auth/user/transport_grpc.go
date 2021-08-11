package user

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/eminetto/talk-microservices-gokit/auth/pb"
	kitlog "github.com/go-kit/kit/log"
)

type grpcServer struct {
	validateToken    grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(svc Service, logger kitlog.Logger) pb.Auth {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return &grpcServer{
		auth: grpctransport.NewServer(
			makeValidateTokenEndpoint(svc),
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Sum", logger)))...,
		),
	}
}

func (s *grpcServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	_, rep, err := s.auth.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ValidateTokenResponse), nil
}

// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sum request to a user-domain sum request. Primarily useful in a server.
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ValidateTokenRequest)
	return validateTokenRequest{Token: req.Token}, nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(validateTokenResponse)
	return &pb.ValidateTokenResponse{Email: resp.Emai}, nil
}