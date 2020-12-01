package profileserver

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/manager"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"google.golang.org/grpc"
	"log"
	"net"
)


func Start() {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"profile", "123",configs.Host, configs.Port, "profiledb")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatalln("PROFILE SERVICE: Cannot create conn to postgresql", psqlInfo)
	}
	serv := grpc.NewServer()
	profileService.RegisterProfileServiceServer(serv, manager.NewProfileServiceManager(db))
	lis, err := net.Listen("tcp","127.0.0.1:8081")
	if err != nil{
		log.Fatalln("PROFILE SERVICE: Cannot create net params")
	}

	err = serv.Serve(lis)
	if err != nil{
		log.Fatalln("PROFILE SERVICE: server serving troubles")
	}

	defer func(){
		if db != nil{
			_ = db.Close()
		}
	}()
}
