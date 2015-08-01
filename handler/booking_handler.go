package handler

import (
	"net/http"

	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/cachingheader"
	"github.com/bborbe/server/handler/contenttype"
	"github.com/bborbe/server/handler/fallback"
	log_handler "github.com/bborbe/server/handler/log"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
	"github.com/bborbe/server/handler_finder/part"
	"github.com/bborbe/server/handler_finder/rest"

	booking_date_handler_create "github.com/bborbe/booking/date/handler/create"
	booking_date_handler_delete "github.com/bborbe/booking/date/handler/delete"
	booking_date_handler_list "github.com/bborbe/booking/date/handler/list"
	booking_date_service "github.com/bborbe/booking/date/service"
)

var logger = log.DefaultLogger

func NewHandler(documentRoot string, dateService booking_date_service.Service) http.Handler {
	logger.Debugf("root: %s", documentRoot)
	fileServer := cachingheader.NewCachingHeaderHandler(contenttype.NewContentTypeHandler(http.FileServer(http.Dir(documentRoot))))
	handlerFinder := part.New("")
	handlerFinder.RegisterHandler("/", fileServer)
	handlerFinder.RegisterHandler("/css", fileServer)
	handlerFinder.RegisterHandler("/js", fileServer)
	handlerFinder.RegisterHandler("/images", fileServer)
	handlerFinder.RegisterHandlerFinder("/date", createDateHandlerFinder("/date", dateService))
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func createDateHandlerFinder(prefix string, dateService booking_date_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_date_handler_list.New(dateService))
	hf.RegisterCreateHandler(booking_date_handler_create.New(dateService))
	hf.RegisterDeleteHandler(booking_date_handler_delete.New(dateService))
	return hf
}
