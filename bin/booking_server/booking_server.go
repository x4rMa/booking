package main

import (
	"flag"
	"net/http"

	booking_handler_configuration "github.com/bborbe/booking/handler_configuration"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"

	booking_date_service "github.com/bborbe/booking/date/service"
	booking_date_storage "github.com/bborbe/booking/date/storage"

	booking_model_service "github.com/bborbe/booking/model/service"
	booking_model_storage "github.com/bborbe/booking/model/storage"

	booking_shooting_service "github.com/bborbe/booking/shooting/service"
	booking_shooting_storage "github.com/bborbe/booking/shooting/storage"

	booking_user_service "github.com/bborbe/booking/user/service"
	booking_user_storage "github.com/bborbe/booking/user/storage"

	booking_authentication_converter "github.com/bborbe/booking/authentication/converter"
	booking_authentication_service "github.com/bborbe/booking/authentication/service"

	booking_authorization_service "github.com/bborbe/booking/authorization/service"
	booking_database_postgres "github.com/bborbe/booking/database/postgres"
	booking_tokengenerator "github.com/bborbe/booking/tokengenerator"
	"github.com/bborbe/eventbus"
)

var (
	logger              = log.DefaultLogger
	addressPtr          = flag.String("a0", ":48568", "Zero address to bind to.")
	documentRootPtr     = flag.String("root", "", "Document root directory")
	logLevelPtr         = flag.String("loglevel", log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	databaseNamePtr     = flag.String("dbname", "", "Database Name")
	databaseUserPtr     = flag.String("dbuser", "", "Database User")
	databasePasswordPtr = flag.String("dbpass", "", "Database Password")
	databaseLoggingPtr  = flag.Bool("dblogging", true, "Dasebase Loggin")
)

func main() {
	defer logger.Close()
	flag.Parse()
	gracehttp.Serve(createServer(*addressPtr, *documentRootPtr, *databaseNamePtr, *databaseUserPtr, *databasePasswordPtr, *databaseLoggingPtr))
}

func createServer(address string, documentRoot string, databaseName string, databaseUser string, databasePassword string, databaseLogging bool) *http.Server {
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)
	db := booking_database_postgres.New(databaseName, databaseUser, databasePassword, databaseLogging)
	dateService := booking_date_service.New(booking_date_storage.New(db))
	tokengenerator := booking_tokengenerator.New()
	modelService := booking_model_service.New(booking_model_storage.New(db), tokengenerator)
	eventbus := eventbus.New()
	shootingService := booking_shooting_service.New(booking_shooting_storage.New(db), eventbus)
	userService := booking_user_service.New(booking_user_storage.New(db))
	authenticationService := booking_authentication_service.New(userService, modelService)
	authorizationService := booking_authorization_service.New(authenticationService.VerifyLogin)
	authenticationConverter := booking_authentication_converter.New()
	handlerConfiguration := booking_handler_configuration.New(documentRoot, dateService, modelService, shootingService, userService, authenticationService,authorizationService ,authenticationConverter)
	return &http.Server{Addr: address, Handler: handlerConfiguration.GetHandler()}
}
