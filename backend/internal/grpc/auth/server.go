package auth

import (
	"context"
	ssov1 "offline_project/gen/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)

	IsAdmin(
		ctx context.Context,
		userID int64,
	) (isAdmin bool, err error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

const (
	emptyValue = 0
)

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	return &ssov1.LoginResponse{
		Token: "token1234",
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("Implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("Implement me")
}
