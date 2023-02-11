package auth

import (
	"GGSAPI/pkg/tooling"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ServiceAuth interface {
	RegisterUser(ctx context.Context, user *UserCreate) (string, string, error)
	AuthorizeUser(ctx context.Context, user *UserAuthorize) (string, string, error)
	RecoverRequest(ctx context.Context, recover *UserRecover) (string, string, error)
	AcceptCode(ctx context.Context, accept *UserAccept) (string, string, error)
	CheckAuth(ctx context.Context, check *UserCheck) (int, error)
}

type ServiceUser interface {
	GetProfile(ctx context.Context, UserID int) (map[string]interface{}, error)
}

type HandlerAuth struct {
	Auth ServiceAuth
	User ServiceUser
}

func NewHandlerAuth(auth ServiceAuth, user ServiceUser) *HandlerAuth {
	return &HandlerAuth{Auth: auth, User: user}
}

func (h *HandlerAuth) Register(r *httprouter.Router) {
	logrus.Info("Auth Handler Initializing")

	r.POST("/auth/sign/up", h.SignUp)
	r.POST("/auth/sign/in", h.SignIn)
	r.POST("/auth/recover/request", h.Recover)
	r.POST("/auth/recover/code", h.AcceptCode)
	r.POST("/auth/me", h.CheckAuth)
}

func (h *HandlerAuth) SignUp(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var CreateUserInput UserCreate

	err := tooling.UnmarshallAll(r.Body, CreateUserInput)
	if err != nil {
		return
	}

	AccessCode, Status, err := h.Auth.RegisterUser(r.Context(), &CreateUserInput)
	if err != nil {
		return
	}

	_, err = writer.Write(tooling.SignInResponse(AccessCode, Status))
	if err != nil {
		return
	}
}

func (h *HandlerAuth) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var AuthorizeUserInput UserAuthorize

	err := tooling.UnmarshallAll(r.Body, AuthorizeUserInput)
	if err != nil {
		return
	}

	AccessCode, Status, err := h.Auth.AuthorizeUser(r.Context(), &AuthorizeUserInput)
	if err != nil {
		return
	}

	_, err = w.Write(tooling.SignInResponse(AccessCode, Status))
	if err != nil {
		return
	}
}

func (h *HandlerAuth) Recover(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var UserRecoverInput UserRecover

	err := tooling.UnmarshallAll(r.Body, UserRecoverInput)
	if err != nil {
		return
	}

	ActionID, Status, err := h.Auth.RecoverRequest(r.Context(), &UserRecoverInput)
	if err != nil {
		return
	}

	w.Write(tooling.RecoverResponse(ActionID, Status))
}

func (h *HandlerAuth) AcceptCode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var AcceptCodeInput UserAccept

	err := tooling.UnmarshallAll(r.Body, &AcceptCodeInput)
	if err != nil {
		return
	}

	AccessCode, Status, err := h.Auth.AcceptCode(r.Context(), &AcceptCodeInput)
	if err != nil {
		return
	}

	_, err = w.Write(tooling.SignInResponse(AccessCode, Status))
	if err != nil {
		return
	}
}

func (h *HandlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var UserCheck UserCheck
	err := tooling.UnmarshallAll(r.Body, UserCheck)
	if err != nil {
		return
	}
	UserID, err := h.Auth.CheckAuth(r.Context(), &UserCheck)
	if err != nil {
		return
	}

	Profile, err := h.User.GetProfile(r.Context(), UserID)
	if err != nil {
		return
	}

	ProfileBytes, err := json.Marshal(Profile)
	if err != nil {
		return
	}

	w.Write(ProfileBytes)
}
