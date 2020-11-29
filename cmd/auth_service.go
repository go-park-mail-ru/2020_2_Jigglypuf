package main

import(
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/authserver"
	profile "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"google.golang.org/grpc"
	"log"
)


func main(){
	profileService, err := grpc.Dial("127.0.0.1:8081")
	if err != nil{
		log.Fatalln("AUTH SERVICE INIT: no connection with profile service")
	}
	profileClient := profile.NewProfileServiceClient(profileService)

	authserver.Start(profileClient, "some_salt")
}
