package user

import (
	"GGSAPI/pkg/user/grpc"
	"context"
)

type ServiceUser struct {
	UserService grpc.UserServiceClient
}

func NewServiceUser(userService grpc.UserServiceClient) *ServiceUser {
	return &ServiceUser{UserService: userService}
}

func (s *ServiceUser) GetProfile(ctx context.Context, UserID int) (map[string]interface{}, error) {
	Profile, err := s.UserService.GetProfile(ctx, &grpc.GetProfileRequest{UserID: int64(UserID)})
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})

	m["username"] = Profile.Username
	m["email"] = Profile.Email
	m["date_birth"] = Profile.DateBirth
	m["location"] = Profile.Location
	m["photos"] = Profile.Photos

	return m, nil
}
