package oauth

import (
	"context"
	"fmt"
	"github.com/flaxius/portia/internal/applications/credentials"
	"github.com/flaxius/portia/proto-gen/authentication/rest_credentials"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type serviceTokenServer struct {
	m sync.Mutex
}

//RestAuthCredentials aa
func RestAuthCredentials() rest_credentials.RestAuthCredentialsServer {
	return &serviceTokenServer{}
}

func (s *serviceTokenServer) Create(ctx context.Context, req *rest_credentials.CreateRestRequest) (*rest_credentials.CreateRestResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var opts []grpc.ServerOption
	fmt.Println(opts)

	clientsId, err := credentials.GenerateRandomAppSecret(17)
	if err != nil {
		fmt.Println("Error to generate the application credentials")
	}

	clientSecret, err := credentials.GenerateRandomAppSecret(44)

	if err != nil {
		fmt.Println("Error to generate the application credentials")
	}
	//TODO generate portia validator grpc
	if req.ProjectId == "" || req.RedirectUris == "" {
		st := status.New(codes.InvalidArgument, "INVALID_NAME")
		// we add details here
		badRequest := &errdetails.BadRequest{}
		violations := make([]*errdetails.BadRequest_FieldViolation, 0)
		violation := errdetails.BadRequest_FieldViolation{Field: "N", Description: "unsupported value"}
		(*badRequest).FieldViolations = append(violations, &violation)
		det, err := st.WithDetails(badRequest)
		if err != nil {
			return nil, st.Err()
		}
		return nil, det.Err()
	}

	return &rest_credentials.CreateRestResponse{
		Rest: &rest_credentials.RestAppDefinition{
			ProjectId:    req.ProjectId,
			ClientId:     "P-" + clientsId + "." + req.ProjectId,
			AuthUri:      "https://portia.io/oauth2/gen/token",
			ClientSecret: clientSecret,
			RedirectUris: []string{req.RedirectUris},
		},
	}, nil
}
