package main

import (
	"flag"
	"net/http"

	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/handler"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"

	booking_date_service "github.com/bborbe/booking/date/service"
	booking_date_storage "github.com/bborbe/booking/date/storage"

	booking_model_service "github.com/bborbe/booking/model/service"
	booking_model_storage "github.com/bborbe/booking/model/storage"

	booking_shooting_service "github.com/bborbe/booking/shooting/service"
	booking_shooting_storage "github.com/bborbe/booking/shooting/storage"
	booking_tokengenerator "github.com/bborbe/booking/tokengenerator"
)

var (
	logger          = log.DefaultLogger
	addressPtr      = flag.String("a0", ":48568", "Zero address to bind to.")
	documentRootPtr = flag.String("root", "", "Document root directory")
	logLevelPtr     = flag.String("loglevel", log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
)

func main() {
	defer logger.Close()
	flag.Parse()
	gracehttp.Serve(createServer(*addressPtr, *documentRootPtr))
}

func createServer(address string, documentRoot string) *http.Server {
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)
	db := database.New("/tmp/booking.db", true)
	tokengenerator := booking_tokengenerator.New()
	dateService := booking_date_service.New(booking_date_storage.New(db))
	modelService := booking_model_service.New(booking_model_storage.New(db),tokengenerator)
	shootingService := booking_shooting_service.New(booking_shooting_storage.New(db))
	return &http.Server{Addr: address, Handler: handler.NewHandler(documentRoot, dateService, modelService,shootingService)}
}
