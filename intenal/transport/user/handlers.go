package user

import (
	"GGSAPI/pkg/auth"
	"GGSAPI/pkg/tooling"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AccountService interface {
	GetProfile(ctx context.Context, UserID int) (map[string]interface{}, error)
	AddInfo(ctx context.Context, input *UserCreateInput) (map[string]interface{}, error)
}

type AccountHandlers struct {
	AccountService AccountService
	Ware           auth.MiddleWare
}

func NewAccountHandlers(accountService AccountService, ware auth.MiddleWare) *AccountHandlers {
	return &AccountHandlers{AccountService: accountService, Ware: ware}
}

func (h *AccountHandlers) Register(router *httprouter.Router) {
	router.GET("/profile", h.Ware.IsAuth(h.GetProfile))
	router.POST("/sign/up/complete", h.Ware.IsAuth(h.UpdateProfile))
}

func (h *AccountHandlers) GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserID := request.Context().Value("user_id").(int)
	Profile, err := h.AccountService.GetProfile(request.Context(), UserID)
	if err != nil {
		return
	}
	ProfileBytes, err := json.Marshal(Profile)
	if err != nil {
		return
	}
	writer.Write(ProfileBytes)
}

func (h *AccountHandlers) UpdateProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var Input UserCreateInput
	UserID := request.Context().Value("user_id").(int64)
	err := tooling.UnmarshallAll(request.Body, Input)
	if err != nil {
		return
	}

	Input.ID = UserID
	Response, err := h.AccountService.AddInfo(request.Context(), &Input)
	if err != nil {
		return
	}

	RespBytes, err := json.Marshal(Response)
	if err != nil {
		return
	}

	writer.Write(RespBytes)
}
