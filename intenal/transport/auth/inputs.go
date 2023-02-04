package auth

import (
	"GGSAPI/pkg/auth/grpc"
	"GGSAPI/pkg/tooling"
)

type UserCreate struct {
	ID        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	IP        string `json:"ip,omitempty"`
	DateBirth string `json:"date_birth,omitempty"`
}

func (a *UserCreate) ToGRPCForAuth() *grpc.SignUpRequest {
	var SignUpReq grpc.SignUpRequest
	SignUpReq.Email = a.Email
	SignUpReq.Password = tooling.Hash(a.Password)
	return &SignUpReq
}

type UserAuthorize struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (a *UserAuthorize) ToGRPCForAuth() *grpc.SignInRequest {
	var SignUpReq grpc.SignInRequest
	SignUpReq.Email = a.Email
	SignUpReq.Password = tooling.Hash(a.Password)
	return &SignUpReq
}

type UserRecover struct {
	Email    string `json:"email,omitempty"`
	Location string `json:"location,omitempty"`
}

func (a *UserRecover) ToGRPCForAuth() *grpc.RecoverPasswordRequest {
	var Recover grpc.RecoverPasswordRequest

	Recover.Email = a.Email

	Country, Region, City, err := tooling.GetLocation(a.Location)
	if err != nil {
		return nil
	}
	Recover.Location.Country = Country
	Recover.Location.Region = Region
	Recover.Location.City = City
	return &Recover
}

type UserAccept struct {
	ActionUID string `json:"action_uid,omitempty"`
	Code      int    `json:"code,omitempty"`
}

func (u *UserAccept) ToGRPC() *grpc.AcceptActionRequest {
	var AcceptAction grpc.AcceptActionRequest
	AcceptAction.ActionUID = u.ActionUID
	AcceptAction.ConfirmCode = int32(u.Code)
	return &AcceptAction
}
