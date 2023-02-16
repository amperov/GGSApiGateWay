package main

import (
	"GGSAPI/intenal/service/auth"
	"GGSAPI/intenal/service/user"
	auth2 "GGSAPI/intenal/transport/auth"
	user2 "GGSAPI/intenal/transport/user"
	auth3 "GGSAPI/pkg/auth"
	"GGSAPI/pkg/auth/grpc"
	grpc3 "GGSAPI/pkg/user/grpc"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	grpc2 "google.golang.org/grpc"
	"net/http"
)

func main() {
	logrus.Print("Service starting")
	router := httprouter.New()
	insecure := grpc2.WithInsecure()

	AuthConn, err := grpc2.Dial("localhost:8082", insecure)
	if err != nil {
		logrus.Print(err)
		return
	}

	UserConn, err := grpc2.Dial("localhost:8083", insecure)
	if err != nil {
		logrus.Print(err)
		return
	}

	AuthClient := grpc.NewAuthorizationClient(AuthConn)
	serviceAuth := auth.NewServiceAuth(AuthClient)
	handlerAuth := auth2.NewHandlerAuth(serviceAuth)
	handlerAuth.Register(router)

	middleWare := auth3.NewMiddleWare(*serviceAuth)
	UserClient := grpc3.NewUserServiceClient(UserConn)
	serviceUser := user.NewServiceUser(UserClient)
	handlers := user2.NewAccountHandlers(serviceUser, middleWare)

	handlers.Register(router)

	Server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err = Server.ListenAndServe()
	if err != nil {
		logrus.Print(err)
		return

	}
}
