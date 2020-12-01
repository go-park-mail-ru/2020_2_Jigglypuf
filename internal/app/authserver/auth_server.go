package authserver

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/manager"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"google.golang.org/grpc"
	"log"
	"net"
)


func Start(profileService profileService.ProfileServiceClient, salt string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, "auth", "123", "authdb")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatalln("AUTH SERVICE: Cannot create conn to postgresql")
	}
	serv := grpc.NewServer()
	authService.RegisterAuthenticationServiceServer(serv,manager.NewAuthServiceManager(db, profileService, salt))
	lis, err := net.Listen("tcp", "auth:8082")
	if err != nil{
		log.Fatalln("AUTH SERVICE: Cannot create net params")
	}

	err = serv.Serve(lis)
	if err != nil{
		log.Fatalln("AUTH SERVICE: server serving troubles")
	}

	defer func(){
		if db != nil{
			_ = db.Close()
		}
	}()
}
