package mainserver

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/router"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
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

func startDBWork() (*sql.DB, *tarantool.Connection, error) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"main", "123",configs.Host, configs.Port, "interfacedb")

	PostgreSQLConnection, DBErr := sql.Open("postgres", psqlInfo)
	if DBErr != nil {
		return nil, nil, errors.New("no postgresql connection")
	}

	TarantoolConnection, DBConnectionErr := tarantool.Connect(session.Host+session.Port, tarantool.Opts{
		User: session.User,
		Pass: session.Password,
	})
	if DBConnectionErr != nil {
		return nil, nil, errors.New("no tarantool connection")
	}

	return PostgreSQLConnection, TarantoolConnection, nil
}

func Start(authenticationServiceClient authService.AuthenticationServiceClient, profileServiceClient profileService.ProfileServiceClient) {
	postgresConn, tarantoolConn, DBErr := startDBWork()
	if DBErr != nil{
		log.Fatalln("MAIN SERVER INIT: NO DB CONN")
	}
	application, err := router.ConfigureHandlers(tarantoolConn, postgresConn, authenticationServiceClient, profileServiceClient)
	if err != nil{
		log.Fatalln("MAIN SERVER INIT: NO GRPC CONN")
	}

	mainRouter := router.ConfigureRouter(application)
	httpServer := configureServer("8080", mainRouter)
	log.Println("Starting server at port 8080")
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
