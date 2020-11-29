package profileserver

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	profileConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	profileDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/manager"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"log"
	"net"
)

func configureProfileRouter(handler *profileDelivery.ProfileHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET(profileConfig.URLPattern, handler.GetProfile)
	router.PUT(profileConfig.URLPattern, handler.UpdateProfile)

	return router
}

func Start() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, configs.User, configs.Password, configs.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatalln("PROFILE SERVICE: Cannot create conn to postgresql")
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
}
