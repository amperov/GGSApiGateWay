package auth

import (
	auth2 "GGSAPI/intenal/service/auth"
	"GGSAPI/intenal/transport/auth"
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type MiddleWare struct {
	ServiceAuth auth2.ServiceAuth
}

func NewMiddleWare(serviceAuth auth2.ServiceAuth) MiddleWare {
	return MiddleWare{serviceAuth}
}

func (w *MiddleWare) IsAuth(handle httprouter.Handle) httprouter.Handle {

	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var id int
		header := request.Header.Get("Authorization")
		headerArray := strings.Split(header, " ")

		id, err := w.ServiceAuth.CheckAuth(context.Background(), &auth.UserCheck{AccessToken: headerArray[1], RefreshToken: headerArray[2]})
		if err != nil {
			logrus.Println(err)
			return
		}
		ctx := context.WithValue(request.Context(), "user_id", id)

		handle(writer, request.WithContext(ctx), params)
	}
}
