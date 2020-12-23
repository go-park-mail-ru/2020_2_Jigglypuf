package mainserver

import (
	"database/sql"
	"errors"
	"fmt"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/router"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/configurator"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"strconv"
	"time"
)

func configureServer(port string, funcHandler http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      funcHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func startDBWork(config *configurator.Config) (*sql.DB, *tarantool.Connection, error) {
	if config.App.DatabaseName == "" || config.App.DatabaseUser == "" || config.App.DatabasePassword == "" {
		return nil, nil, models.ErrFooIncorrectPath
	}
	fmt.Printf("%v",config.App)
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.App.DatabaseUser, config.App.DatabasePassword, config.DatabaseDomain, config.DatabasePort,
		config.App.DatabaseName)


	PostgreSQLConnection, DBErr := sql.Open("postgres", psqlInfo)
	if DBErr != nil {
		return nil, nil, errors.New("no postgresql connection")
	}
	TarantoolConnection, DBConnectionErr := tarantool.Connect(config.Tarantool.Domain+":" +strconv.Itoa(config.Tarantool.Port), tarantool.Opts{
		User: config.Tarantool.User,
		Pass: config.Tarantool.Password,
	})
	if DBConnectionErr != nil {
		return nil, nil, errors.New("no tarantool connection")
	}

	return PostgreSQLConnection, TarantoolConnection, nil
}

func Start(authenticationServiceClient authService.AuthenticationServiceClient, profileServiceClient profileService.ProfileServiceClient,
	config *configurator.Config) {
	postgresConn, tarantoolConn, DBErr := startDBWork(config)
	if DBErr != nil {
		log.Fatalln("MAIN SERVER INIT: NO DB CONN")
	}
	application, err := router.ConfigureHandlers(tarantoolConn, postgresConn, authenticationServiceClient, profileServiceClient, config)
	if err != nil {
		log.Fatalln("MAIN SERVER INIT: NO GRPC CONN")
	}

	mainRouter := router.ConfigureRouter(application)
	httpServer := configureServer(strconv.Itoa(config.App.Port), mainRouter)
	log.Printf("Starting server at port %d\n", config.App.Port)
	serverErr := httpServer.ListenAndServe()
	if serverErr != nil {
		log.Fatalln(err)
	}

	defer func() {
		if postgresConn != nil {
			_ = postgresConn.Close()
		}
		if tarantoolConn != nil {
			_ = tarantoolConn.Close()
		}
	}()
}
