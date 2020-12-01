package main

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/profileserver"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
)

func main() {
	auth, err := grpc.Dial("auth:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("PROFILE SERVICE INIT: no connection with auth service")
	}
	profileClient := authService.NewAuthenticationServiceClient(auth)
	profileserver.Start(profileClient)
}
