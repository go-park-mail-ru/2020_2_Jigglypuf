package profileserver

import (
	"database/sql"
	"fmt"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/manager"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/configurator"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func Start(auth authService.AuthenticationServiceClient, config *configurator.Config) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Profile.DatabaseUser, config.Profile.DatabasePassword,
		config.DatabaseDomain, config.DatabasePort, config.Profile.DatabaseName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("PROFILE SERVICE: Cannot create conn to postgresql", psqlInfo)
	}
	serv := grpc.NewServer()
	profileService.RegisterProfileServiceServer(serv, manager.NewProfileServiceManager(db, auth))
	lis, err := net.Listen("tcp", config.Profile.Domain+":"+strconv.Itoa(config.Profile.Port))
	if err != nil {
		log.Fatalln("PROFILE SERVICE: Cannot create net params")
	}

	log.Println("Starting profile service")
	err = serv.Serve(lis)
	if err != nil {
		log.Fatalln("PROFILE SERVICE: server serving troubles")
	}

	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()
}
