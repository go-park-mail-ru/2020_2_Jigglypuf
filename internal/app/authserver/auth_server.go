package authserver

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/manager"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/configurator"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func Start(profileService profileService.ProfileServiceClient, config *configurator.Config, salt string) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Auth.DatabaseUser, config.Auth.DatabasePassword, config.DatabaseDomain,
		config.DatabasePort, config.Auth.DatabaseName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("AUTH SERVICE: Cannot create conn to postgresql")
	}
	serv := grpc.NewServer()
	authService.RegisterAuthenticationServiceServer(serv, manager.NewAuthServiceManager(db, profileService, salt))
	lis, err := net.Listen("tcp", config.Auth.Domain + ":" + strconv.Itoa(config.Auth.Port))
	if err != nil {
		log.Fatalln("AUTH SERVICE: Cannot create net params")
	}

	log.Println("Starting Auth Service")
	err = serv.Serve(lis)
	if err != nil {
		log.Fatalln("AUTH SERVICE: server serving troubles")
	}

	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()
}
