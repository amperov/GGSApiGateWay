package auth

import (
	"GGSAPI/intenal/transport/auth"
	"GGSAPI/pkg/auth/grpc"
	"context"
)

type ServiceAuth struct {
	AuthClient grpc.AuthorizationClient
}

func NewServiceAuth(authClient grpc.AuthorizationClient) *ServiceAuth {
	return &ServiceAuth{AuthClient: authClient}
}

func (s *ServiceAuth) CheckAuth(ctx context.Context, check *auth.UserCheck) (int, error) {
	Response, err := s.AuthClient.Identity(ctx, check.ToGRPCForAuth())
	if err != nil {
		return 0, err
	}
	return int(Response.GetUserID()), nil
}

func (s *ServiceAuth) RegisterUser(ctx context.Context, user *auth.UserCreate) (string, string, error) {
	Response, err := s.AuthClient.SignUp(ctx, user.ToGRPCForAuth())
	if err != nil {
		return "", err.Error(), err
	}
	return Response.GetAccessCode(), Response.GetStatus(), nil
}

func (s *ServiceAuth) AuthorizeUser(ctx context.Context, user *auth.UserAuthorize) (string, string, error) {
	Response, err := s.AuthClient.SignIn(ctx, user.ToGRPCForAuth())
	if err != nil {
		return "", err.Error(), err
	}
	return Response.GetAccessCode(), Response.GetStatus(), nil
}

func (s *ServiceAuth) RecoverRequest(ctx context.Context, recover *auth.UserRecover) (string, string, error) {
	Response, err := s.AuthClient.RecoverPassword(ctx, recover.ToGRPCForAuth())
	if err != nil {
		return "", err.Error(), err
	}

	return Response.GetActionUID(), Response.GetStatus(), nil
}

func (s *ServiceAuth) AcceptCode(ctx context.Context, accept *auth.UserAccept) (string, string, error) {

	Response, err := s.AuthClient.AcceptAction(ctx, accept.ToGRPC())
	if err != nil {
		return "", err.Error(), err
	}

	return Response.GetAccessCode(), Response.GetStatus(), nil
}
