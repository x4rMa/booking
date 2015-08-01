package main

import (
	"flag"
	"net/http"

	"github.com/bborbe/booking/database"
	"github.com/bborbe/booking/date/service"
	"github.com/bborbe/booking/date/storage"
	"github.com/bborbe/booking/handler"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"
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
	dateService := service.New(storage.New(database.New("/tmp/booking.db", true)))
	return &http.Server{Addr: address, Handler: handler.NewHandler(documentRoot, dateService)}
}
