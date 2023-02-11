package main

import (
	"GGSAPI/intenal/service/auth"
	"GGSAPI/intenal/service/user"
	auth2 "GGSAPI/intenal/transport/auth"
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
	conn, err := grpc2.Dial("localhost:8082", insecure)
	if err != nil {
		logrus.Print(err)
		return
	}
	AuthClient := grpc.NewAuthorizationClient(conn)

	UserClient := grpc3.NewUserServiceClient(conn)

	serviceUser := user.NewServiceUser(UserClient)

	serviceAuth := auth.NewServiceAuth(AuthClient)

	handlerAuth := auth2.NewHandlerAuth(serviceAuth, serviceUser)
	handlerAuth.Register(router)

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
