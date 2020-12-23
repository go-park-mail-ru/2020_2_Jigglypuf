package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/profileserver"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/configurator"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	configPath := utils.ParseConfigPath()
	config, configErr := configurator.Run(configPath)
	if configErr != nil{
		log.Fatalln("No configuration")
	}
	fmt.Printf("%v\n", config)
	auth, err := grpc.Dial(config.Auth.Domain + ":" + strconv.Itoa(config.Auth.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalln("PROFILE SERVICE INIT: no connection with auth service")
	}
	profileClient := authService.NewAuthenticationServiceClient(auth)
	profileserver.Start(profileClient, config)
}
