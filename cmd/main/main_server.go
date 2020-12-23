package main

import (
	_ "github.com/go-park-mail-ru/2020_2_Jigglypuf/docs"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/mainserver"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/configurator"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

// Backend doc
// @title CinemaScope Backend API
// @version 0.5
// @description This is a backend API
// @host https://cinemascope.space
// @BasePath /
func main() {
	configPath := utils.ParseConfigPath()
	config, configErr := configurator.Run(configPath)
	if configErr != nil {
		log.Fatalln("Incorrect config path")
	}

	profileServiceConn, profileServiceErr := grpc.Dial(config.Profile.Domain+":"+strconv.Itoa(config.Profile.Port),
		grpc.WithInsecure())
	if profileServiceErr != nil {
		log.Fatalln("MAIN SERVICE INIT: no profile service conn")
	}

	authServiceConn, err := grpc.Dial(config.Auth.Domain+":"+strconv.Itoa(config.Auth.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalln("MAIN SERVICE INIT: no authentication service conn")
	}

	profileServiceClient := profileService.NewProfileServiceClient(profileServiceConn)
	AuthServiceClient := authService.NewAuthenticationServiceClient(authServiceConn)

	mainserver.Start(AuthServiceClient, profileServiceClient, config)
}
