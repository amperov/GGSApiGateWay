package user

import (
	"GGSAPI/intenal/transport/user"
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
	m["date_birth"] = Profile.DateBirth
	m["location"] = Profile.Location
	m["photos"] = Profile.Photos

	return m, nil
}
func (s *ServiceUser) AddInfo(ctx context.Context, input *user.UserCreateInput) (map[string]interface{}, error) {
	Response, err := s.UserService.AddInfo(ctx, &grpc.AddInfoRequest{
		UserID:   input.ID,
		Username: input.Username,
		Location: &grpc.Location{
			Country: input.Location.Country,
			Region:  input.Location.Region,
			City:    input.Location.City,
		},
		DateBirth: input.DateBirth,
	})
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m["status"] = Response.GetStatus()

	return m, nil
}
