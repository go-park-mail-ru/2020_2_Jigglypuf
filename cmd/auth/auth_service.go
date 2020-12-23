package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/authserver"
	profile "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
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
		log.Fatalln("No config setup")
	}
	fmt.Printf("%v\n", config)
	profileService, err := grpc.Dial(config.Profile.Domain + ":" + strconv.Itoa(config.Profile.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalln("AUTH SERVICE INIT: no connection with profile service")
	}
	profileClient := profile.NewProfileServiceClient(profileService)

	authserver.Start(profileClient, config,"some_salt")
}
