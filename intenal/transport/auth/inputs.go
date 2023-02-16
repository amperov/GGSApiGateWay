package auth

import (
	"GGSAPI/pkg/auth/grpc"
	"GGSAPI/pkg/tooling"
)

type UserCreate struct {
	ID        int64    `json:"id,omitempty"`
	Username  string   `json:"username,omitempty"`
	Email     string   `json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
	Location  Location `json:"location,omitempty"`
	DateBirth string   `json:"date_birth,omitempty"`
}
type Location struct {
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	City    string `json:"city,omitempty"`
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
	Email    string   `json:"email,omitempty"`
	Location Location `json:"location,omitempty"`
}

func (a *UserRecover) ToGRPCForAuth() *grpc.RecoverPasswordRequest {
	var Recover grpc.RecoverPasswordRequest
	Recover.Email = a.Email
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

type UserCheck struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token"`
}

func (c *UserCheck) ToGRPCForAuth() *grpc.IdentityRequest {

	return &grpc.IdentityRequest{AccessToken: c.AccessToken, RefreshToken: c.RefreshToken}

}
