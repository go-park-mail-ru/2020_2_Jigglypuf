package main

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/mainserver"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
)

// Backend doc
// @title CinemaScope Backend API
// @version 0.5
// @description This is a backend API
// @host https://cinemascope.space
// @BasePath /
func main() {

	profileServiceConn, profileServiceErr := grpc.Dial("profile:8081", grpc.WithInsecure())
	if profileServiceErr != nil {
		log.Fatalln("MAIN SERVICE INIT: no profile service conn")
	}

	authServiceConn, err := grpc.Dial("auth:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("MAIN SERVICE INIT: no authentication service conn")
	}

	profileServiceClient := profileService.NewProfileServiceClient(profileServiceConn)
	AuthServiceClient := authService.NewAuthenticationServiceClient(authServiceConn)

	mainserver.Start(AuthServiceClient, profileServiceClient)
}
