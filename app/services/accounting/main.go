package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/handlers/accountinghandler"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/service/accountingservice"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/storage/accountingstorage"
)

const configFile = "/config.yaml"

func main() {
	// Construct the application logger.
	log := logger.NewLogger(os.Stdout, "main: ")

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
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

	log.Println("starting serviceInstance", "version", configreader.GetString("buildmode"))
	defer log.Println("shutdown complete")

	dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
		" port=" + configreader.GetString("database.port") +
		" user=" + configreader.GetString("database.user") +
		" password=" + configreader.GetString("database.password") +
		" dbname=" + configreader.GetString("database.dbName") +
		" sslmode=" + configreader.GetString("database.sslmode"))

	database, err := db.OpenDB("postgres", dbURL)
	if err != nil {
		return err
	}

	storageInstance := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

	serviceInstance := accountingservice.NewService(accountingservice.Params{
		Storage: storageInstance,
		Logger:  logger.NewLogger(os.Stdout, "service: "),
	})

	handler := accountinghandler.NewHandler(accountinghandler.Params{
		Logger:         logger.NewLogger(os.Stdout, "endpoint: "),
		AccountService: serviceInstance,
	})

	apiLogger := logger.NewLogger(os.Stdout, "api: ")
	serverAPI := api.NewAPI(handler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server: ")
	serverConfig := getServerConfig(serverAPI, nil, serverLogger)

	serviceServer := server.NewServer(serverConfig)

	serverErrors := make(chan interface{})

	serviceServer.StartListening(serverErrors)

	<-serverErrors

	return nil
}

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) *config.Config {
	var (
		address          = configreader.GetString("server.address")
		readTime         = configreader.GetDuration("server.readTime")
		writeTime        = configreader.GetDuration("server.writeTime")
		idleTime         = configreader.GetDuration("server.idleTime")
		readerHeaderTime = configreader.GetDuration("server.readerHeaderTime")
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
