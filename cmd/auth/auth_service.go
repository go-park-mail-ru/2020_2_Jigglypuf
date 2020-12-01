package main

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/authserver"
	profile "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
)

func main() {
	profileService, err := grpc.Dial("profile:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("AUTH SERVICE INIT: no connection with profile service")
	}
	profileClient := profile.NewProfileServiceClient(profileService)

	authserver.Start(profileClient, "some_salt")
}
