package authserver

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	authDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/manager"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"log"
	"net"
)

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router {
	authAPIHandler := httprouter.New()

	authAPIHandler.POST(configs.URLPattern+"register/", authHandler.RegisterHandler)
	authAPIHandler.POST(configs.URLPattern+"login/", authHandler.AuthHandler)
	authAPIHandler.POST(configs.URLPattern+"logout/", authHandler.SignOutHandler)

	return authAPIHandler
}

func Start(profileService profileService.ProfileServiceClient, salt string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, configs.User, configs.Password, configs.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatalln("AUTH SERVICE: Cannot create conn to postgresql")
	}
	serv := grpc.NewServer()
	authService.RegisterAuthenticationServiceServer(serv,manager.NewAuthServiceManager(db, profileService, salt))
	lis, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil{
		log.Fatalln("AUTH SERVICE: Cannot create net params")
	}

	err = serv.Serve(lis)
	if err != nil{
		log.Fatalln("AUTH SERVICE: server serving troubles")
	}
}
