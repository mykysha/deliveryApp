package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/grpcserver"
	"github.com/nndergunov/deliveryApp/app/pkg/server"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/grpc/handler"
	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/rest/handler/accountinghandler"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/service/accountingservice"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/storage/accountingstorage"
)

const configFile = "/config.yaml"

func main() {
	// Construct the application logger.
	l := logger.NewLogger(os.Stdout, "main: ")

	// Perform the startup and shutdown sequence.
	if err := run(l); err != nil {
		log.Fatal("startup", "ERROR", err)
	}
}

func run(log *logger.Logger) error {
	confPath, err := os.Getwd()
	if err != nil {
		return err
	}

	err = configreader.SetConfigFile(confPath + configFile)
	if err != nil {
		return err
	}

	dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
		" port=" + configreader.GetString("database.port") +
		" user=" + configreader.GetString("database.user") +
		" password=" + configreader.GetString("database.password") +
		" dbname=" + configreader.GetString("database.dbName") +
		" sslmode=" + configreader.GetString("database.sslmode"))

	//*** grpc ***
	grpcDatabase, err := db.OpenDB("postgres", dbURL)
	if err != nil {
		return err
	}

	grpcStorage := accountingstorage.NewStorage(accountingstorage.Params{DB: grpcDatabase})

	grpcService := accountingservice.NewService(accountingservice.Params{
		Storage: grpcStorage,
		Logger:  logger.NewLogger(os.Stdout, "grpc service: "),
	})

	grpcHandler := handler.NewHandler(handler.Params{
		Logger:  logger.NewLogger(os.Stdout, "grpc endpoint: "),
		Service: grpcService,
	})

	grpcServer := grpcserver.NewGRPCServer(grpcHandler, logger.NewLogger(os.Stdout, "grpc server: "))

	grpcServerStopChan := make(chan interface{})
	grpcServer.StartListening(configreader.GetString("server.grpc.address"), grpcServerStopChan)

	//*** rest ***
	restDatabase, err := db.OpenDB("postgres", dbURL)
	if err != nil {
		return err
	}

	restStorage := accountingstorage.NewStorage(accountingstorage.Params{DB: restDatabase})

	restService := accountingservice.NewService(accountingservice.Params{
		Storage: restStorage,
		Logger:  logger.NewLogger(os.Stdout, "rest service: "),
	})

	restHandler := accountinghandler.NewHandler(accountinghandler.Params{
		Logger:  logger.NewLogger(os.Stdout, "rest endpoint: "),
		Service: restService,
	})

	restAPI := api.NewAPI(restHandler, logger.NewLogger(os.Stdout, "rest api: "))

	restServerConfig := getServerConfig(v1.EnableCORS(restAPI), nil, logger.NewLogger(os.Stdout, "rest server: "))

	restServer := server.NewServer(restServerConfig)

	restServerStopChan := make(chan interface{})
	restServer.StartListening(restServerStopChan)

	serverWG := new(sync.WaitGroup)
	numberOfServersRunning := 2

	serverWG.Add(numberOfServersRunning)

	go func(wg *sync.WaitGroup) {
		<-grpcServerStopChan

		wg.Done()
	}(serverWG)

	go func(wg *sync.WaitGroup) {
		<-restServerStopChan

		wg.Done()
	}(serverWG)

	serverWG.Wait()
	return nil
}

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) *config.Config {
	var (
		address          = configreader.GetString("server.rest.address")
		readTime         = configreader.GetDuration("server.rest.readTime")
		writeTime        = configreader.GetDuration("server.rest.writeTime")
		idleTime         = configreader.GetDuration("server.rest.idleTime")
		readerHeaderTime = configreader.GetDuration("server.rest.readerHeaderTime")
	)

	return &config.Config{
		Address:           address,
		ReadTimeout:       readTime,
		WriteTimeout:      writeTime,
		IdleTimeout:       idleTime,
		ReadHeaderTimeout: readerHeaderTime,
		ErrorLog:          errorLog,
		ServerLogger:      serverLogger,
		Handler:           handler,
	}
}
